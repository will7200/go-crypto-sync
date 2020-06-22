package pc

import "github.com/will7200/go-crypto-sync/pkg/personalcapital"

func HoldingTypeToHoldingRequest(t personalcapital.HoldingsType) personalcapital.HoldingsUpdateRequest {
	return personalcapital.HoldingsUpdateRequest{
		Ticker:                       t.Ticker,
		Quantity:                     t.Quantity,
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
		Price:                        t.Price,
		HoldingPercentage:            t.HoldingPercentage,
		UserAccountId:                t.UserAccountID,
		PriceSource:                  t.PriceSource,
		ValueSortIndex:               t.ValueSortIndex,
		CostBasis:                    t.CostBasis,
		Value:                        t.Value,
		OneDayPercentChange:          t.OneDayPercentChange,
	}
}
