package holdings

import (
	"regexp"
	"strings"
)

type (
	IHolding interface {
		CurrencySymbolName() string
		CurrencyName() string
		TotalSharesString() string
	}
	Holding struct {
		SymbolName  string
		FullName    string
		TotalShares string
	}

	Holdings []Holding
)

func (h Holding) CurrencySymbolName() string {
	return h.SymbolName
}

func (h Holding) CurrencyName() string {
	return h.FullName
}

func (h Holding) TotalSharesString() string {
	return h.TotalShares
}

func (h Holdings) HasCurrencySymbolName(symbol string) bool {
	for _, v := range h {
		if strings.ToLower(symbol) == strings.ToLower(v.SymbolName) {
			return true
		}
	}
	return false
}

func (h Holdings) HasCurrencyName(name string) bool {
	for _, v := range h {
		if strings.ToLower(name) == strings.ToLower(v.FullName) {
			return true
		}
	}
	return false
}

type HasInfo struct {
	Exists bool
	LPos   int
	RPos   int
}

func (h Holdings) HasCurrencyMap(left func(l IHolding) string, right func(r IHolding) string, ih ...IHolding) map[string]HasInfo {
	mb := make(map[string]HasInfo, len(h))
	for index, v := range h {
		mb[strings.ToLower(left(v))] = HasInfo{
			Exists: false,
			LPos:   index,
			RPos:   -1,
		}
	}
	for index, v := range ih {
		key := strings.ToLower(right(v))
		if val, ok := mb[key]; ok {
			if val.Exists {
				panic("conflicting entries")
			}
			val.Exists = true
			val.RPos = index
			mb[key] = val
		}
	}
	return mb
}

func (h Holdings) SearchByPattern(pattern regexp.Regexp) []int {
	panic("implement me")
}

type Account interface {
	GetHoldings() (Holdings, error)
}

type Price interface {
	GetExchange(currency1, currency2 string) (string, error)
}
