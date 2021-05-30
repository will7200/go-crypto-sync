![Go](https://github.com/will7200/go-crypto-sync/workflows/Go/badge.svg)

# Crypto Sync

Sync your Cryptocurrency Holdings to a Personal Capital Account

## Supported Cryptocurrency Holdings

1. Coinbase
2. Ethereum and derivatives via etherscan
3. Coinbase Pro
4. Binance Coin and derivatives via bscscan
5. Binance and Binance.us
6. CoinEx
7. Yieldwatch (stacks and LP on binance network)

## Matrix Support

|             | holdings           | pricing            |
|:------------|:-------------------|:-------------------|
| binance     | :heavy_check_mark: | :heavy_check_mark: |
| bscscan     | :heavy_check_mark: | :x:                |
| coinbase    | :heavy_check_mark: | :heavy_check_mark: |
| coinbasepro | :heavy_check_mark: | :x:                |
| coinex      | :heavy_check_mark: | :x:                |
| coingecko   | :x:                | :heavy_check_mark: |
| etherscan   | :heavy_check_mark: | :x:                |
| nomics      | :x:                | :heavy_check_mark: |
| yieldwatch  | :heavy_check_mark: | :x:                |


## Getting Started

1. Download from releases page for your operating system. Current supported os are: Windows, Mac, and Linux

2. Copy the config.example.toml and rename to config.toml. Remove/comment out any providers that you are not currently
   using. Get the api keys you need for you use case.

3. Create a new investment account name `CryptoSync managed automatically` in Personal Capital or anything you would
   like, just make sure to update the config file to reflect the change.

4. run `crypto-sync.exe sync --destination pc all` and check your account The first time connecting you will be asked
   your MFA if setup in Personal Capital

## Example Commands

1. Sync your Coinbase to your Personal Capital

Sample Config

```toml
debug = false
destination = "personalcapital"
# you may also specify multiple sources seperated by a comma
# this will allow get a value of a token from another provider should the other not provide any data
# make sure to not have spaces between them
# example: nomics,coinbase,coingecko
priceDataSource = "coinbase"
destinationCurrencyAs = "USD"
onHoldingNotFound = "zeroQuantity"
[pricing]
[pricing.nomics]
apiKey = "your-key"
debug = false

[pricing.coinbase]
apiKey = "your-key"
apiSecret = "your-secret"
debug = false

[symbols]
[symbols.global]
[symbols.override]
[symbols.override.nomics]
IOTA = "IOT"
[symbols.override.coingecko]
Cake = "pancakeswap-token"

[holdings]
[holdings.coinbase]
apiKey = "your-key"
apiSecret = "your-secret"
debug = false

[holdings.etherscan]
apiKey = "your-key"
debug = false
[[holdings.etherscan.account]]
address = "your-address"
symbolName = "ETH"
fullName = "Ethereum"
decimals = 18
[[holdings.etherscan.account]]
address = "your-address"
contractAddress = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
symbolName = "USDC"
fullName = "USD Coin"
decimals = 6

[holdings.coinbasepro]
debug = true
[[holdings.coinbasepro.portfolios]]
name = "default_portfolio"
secret = "some_secret"
key = "some_key"
passphrase = "some_pass_phrase"
# If base_url is left blank with default to the value below
base_url = "https://api.pro.coinbase.com"

[holdings.bscscan]
apiKey = "some-api-key"
debug = false
# If base_url is defaulted to the below url
# base_url = "https://api.bscscan.com/api?"
[[holdings.bscscan.account]]
address = "0xaddress"
symbolName = "BNB"
fullName = "Binance Coin"
decimals = 18

[holdings.binance-us]
debug = false
apiKey = "some-api-key"
secret = "some-secret"

[holdings.binance]
debug = false
apiKey = "some-api-key"
secret = "some-secret"

[holdings.coinex]
debug = false
accessId = "someAccessId"
secretKey = "someSecretKey"

[holdings.yieldwatch]
debug = true
platforms = ["pancake"]
walletAddresses = ["0xSomeWalletAddress"]
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

Sync from all holding positions in config file

```bash
crypto-sync sync --destination pc all
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