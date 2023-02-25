module github.com/will7200/go-crypto-sync

go 1.16

require (
	github.com/HereMobilityDevelopers/mediary v1.0.0
	github.com/alecthomas/kong v0.2.16
	github.com/davecgh/go-spew v1.1.1
	github.com/google/go-querystring v1.0.0
	github.com/mitchellh/mapstructure v1.4.1
	github.com/nanmu42/etherscan-api v1.2.0
	github.com/pelletier/go-toml v1.8.0
	github.com/preichenberger/go-coinbasepro/v2 v2.0.5
	github.com/shopspring/decimal v1.2.0
	github.com/stretchr/testify v1.5.1
	go.uber.org/zap v1.16.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.1.0
)

replace github.com/nanmu42/etherscan-api v1.2.0 => github.com/will7200/etherscan-api v1.2.1-0.20210516225556-39daf32d6557
