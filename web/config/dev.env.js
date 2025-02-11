'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  BASE_API: '"https://calibration-api.filswan.com/"',
  BASE_PAYMENT_GATEWAY_API: '"https://calibration-mcp-api.filswan.com/"',
  BASE_MINERADDRESS: '"https://calibration.filscan.io/address/miner?address="'
})
