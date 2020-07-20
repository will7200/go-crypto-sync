![Go](https://github.com/will7200/go-crypto-sync/workflows/Go/badge.svg)
# Crypto Sync
Sync your Cryptocurrency to a Personal Capital Account
## Supported Cryptocurrency Holdings
1. Coinbase

## Example Commands
1. Sync your Coinbase to your Personal Capital
  
Sample Config
```toml
debug = false
destination = "personalcapital"
priceDataSource = "coinbase"
destinationCurrencyAs = "USD"
onHoldingNotFound = "zeroQuantity"
[holdings]
    [holdings.coinbase]
    apiKey = "your-key"
    apiSecret = "your-secret"
    debug = false
[destinations]
    [destinations.personalcapital]
    email = "example@email.com"
    password = "password"
    accountName = "CryptoSync managed automatically"
``` 
Commands to Execute
```bash
crypto-sync sync --destination pc coinbase 
```

## Tools Required to
1. OpenAPI Generator  
The one used for this project is from [here](https://openapi-generator.tech/docs/installation)
```bash
npm install @openapitools/openapi-generator-cli -g
```