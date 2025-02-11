package scheduler

import (
	"fmt"
	"payment-bridge/common/constants"
	"payment-bridge/models"
	"payment-bridge/on-chain/client"
	"payment-bridge/on-chain/goBind"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
)

func Refund() error {
	ethClient, _, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	swanPaymentTransactor, err := client.GetSwanPaymentTransactor(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	//refund(int64(903), swanPaymentTransactor, tansactOpts)

	dealFiles, err := models.GetDealFilesByStatus(constants.PROCESS_STATUS_DEAL_SENT)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, dealFile := range dealFiles {
		err = refund(ethClient, dealFile.ID, swanPaymentTransactor)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}

func refund(ethClient *ethclient.Client, dealFileId int64, swanPaymentTransactor *goBind.SwanPaymentTransactor) error {
	offlineDealsNotUnlocked, err := models.GetOfflineDealsNotUnlockedByDealFileId(dealFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	if len(offlineDealsNotUnlocked) > 0 {
		msg := fmt.Sprintf("%d deals not unlocked or unlock failed, cannot refund for the deal file", len(offlineDealsNotUnlocked))
		logs.GetLogger().Info(msg)
		return nil
	}

	srcFiles, err := models.GetSourceFilesByDealFileId(dealFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	var srcFilePayloadCids []string
	for _, srcFile := range srcFiles {
		lockedPayment, err := client.GetLockedPaymentInfo(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			return err
		}

		err = models.UpdateSourceFileRefundAmount(srcFile.ID, lockedPayment.LockedFee)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			return err
		}

		srcFilePayloadCids = append(srcFilePayloadCids, srcFile.PayloadCid)
	}

	privateKey, publicKeyAddress, err := client.GetPrivateKeyPublicKey(constants.PRIVATE_KEY_ON_POLYGON)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	tansactOpts, err := client.GetTransactOpts(ethClient, privateKey, *publicKeyAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	refundStatus := constants.PROCESS_STATUS_UNLOCK_REFUNDED
	tx, err := swanPaymentTransactor.Refund(tansactOpts, srcFilePayloadCids)
	if err != nil {
		refundStatus = constants.PROCESS_STATUS_UNLOCK_REFUNDFAILED
		logs.GetLogger().Error(err.Error())
	}

	for _, srcFile := range srcFiles {
		txHash := ""
		if tx != nil {
			txHash = tx.Hash().Hex()
		}

		logs.GetLogger().Info("refund stats:", refundStatus, " tx hash:", txHash)

		err = models.UpdateSourceFileRefundStatus(srcFile.ID, refundStatus, txHash)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			continue
		}
	}

	err = models.UpdateDealFileStatus(dealFileId, refundStatus)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	return nil
}
