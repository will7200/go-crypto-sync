![Go](https://github.com/will7200/go-crypto-sync/workflows/Go/badge.svg)

# Crypto Sync

Sync your Cryptocurrency to a Personal Capital Account

## Supported Cryptocurrency Holdings

1. Coinbase
2. Ethereum via etherscan

## Example Commands

1. Sync your Coinbase to your Personal Capital

Sample Config

```toml
# config.example.toml
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
[holdings.ether]
apiKey = "your-key"
debug = false
[[holdings.ether.account]]
address = "your-address"
symbolName = "ETH"
fullName = "Ethereum"
decimals = 18
[[holdings.ether.account]]
address = "your-address"
contractAddress = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
symbolName = "USDC"
fullName = "USD Coin"
decimals = 6
[destinations]
[destinations.personalcapital]
email = "example@email.com"
password = "password"
accountName = "CryptoSync managed automatically"
``` 

Commands to Execute

Sync only from coinbase

```bash
crypto-sync sync --destination pc coinbase 
```

Sync from coinbase and Ethereum

```bash
crypto-sync sync --destination pc coinbase ether
```

### Sync Command

```bash
Usage: crypto-sync.exe sync <holding-accounts> ...

Sync holdings to another account

Arguments:
  <holding-accounts> ...    Holdings to fetch from

Flags:
  --help                       Show context-sensitive help.
  --debug                      Enable debug mode.
  --file-name="config.toml"    File to read conf from

  --destination=STRING         Sync to Destination

```

## Tools Required to

1. OpenAPI Generator  
   The one used for this project is from [here](https://openapi-generator.tech/docs/installation)

```bash
npm install @openapitools/openapi-generator-cli -g
```