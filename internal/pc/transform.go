package pc

import (
	"github.com/shopspring/decimal"

	"github.com/will7200/go-crypto-sync/pkg/personalcapital"
)

func HoldingTypeToHoldingRequest(t personalcapital.HoldingsType) personalcapital.HoldingsUpdateRequest {
	quantity := decimal.NewFromFloat(t.Quantity)
	price := decimal.NewFromFloat(t.Price)
	costBasis := decimal.NewFromFloat(t.CostBasis)
	value := decimal.NewFromFloat(t.Value)
	return personalcapital.HoldingsUpdateRequest{
		Ticker:                       t.Ticker,
		Quantity:                     quantity.StringFixed(2),
		ManualClassification:         t.ManualClassification,
		IsMarketMover:                t.IsMarketMover,
		AccountName:                  t.AccountName,
		OneDayPercentChangeSortIndex: t.OneDayPercentChangeSortIndex,
		OneDayValueChange:            t.OneDayValueChange,
		Change:                       t.Change,
		Description:                  t.Description,
		Source:                       t.Source,
		ChangeSortIndex:              t.ChangeSortIndex,
		OneDayValueChangeSortIndex:   t.OneDayValueChangeSortIndex,
		MarketType:                   t.MarketType,
		SourceAssetId:                t.SourceAssetID,
		External:                     t.External,
		HoldingType:                  t.HoldingType,
		Price:                        price.StringFixed(18),
		HoldingPercentage:            t.HoldingPercentage,
		UserAccountId:                t.UserAccountID,
		PriceSource:                  t.PriceSource,
		ValueSortIndex:               t.ValueSortIndex,
		CostBasis:                    costBasis.StringFixed(2),
		Value:                        value.StringFixed(2),
		OneDayPercentChange:          t.OneDayPercentChange,
	}
}
