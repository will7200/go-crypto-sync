package holdings

import (
	"math/big"
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

func (h Holdings) MapReduce() (final Holdings) {
	rHolds := make(map[string]Holdings)
	for i, holding := range h {
		keyHolds, ok := rHolds[holding.SymbolName]
		if !ok {
			keyHolds = make(Holdings, 0, 1)
			rHolds[holding.SymbolName] = keyHolds
		}
		keyHolds = append(keyHolds, h[i])
		rHolds[holding.SymbolName] = keyHolds
	}
	final = make(Holdings, len(rHolds))
	i := 0
	for _, value := range rHolds {
		totalShares := value[0].TotalShares
		if len(value) > 1 {
			shares := new(big.Float)
			for _, hc := range value {
				t, _, err := big.ParseFloat(hc.TotalShares, 10, 18, big.ToNearestEven)
				if err != nil {
					panic(err)
				}
				shares.Add(shares, t)
			}
			totalShares = shares.String()
		}
		final[i] = Holding{
			SymbolName:  value[0].SymbolName,
			FullName:    value[0].FullName,
			TotalShares: totalShares,
		}
		i += 1
	}
	return final
}

type HasInfo struct {
	LPos int
	RPos int
}

func (hi HasInfo) FoundBoth() bool {
	return hi.LPos != -1 && hi.RPos != -1
}

func (hi HasInfo) LeftOnly() bool {
	return hi.LPos != -1 && hi.RPos == -1
}

func (hi HasInfo) RightOnly() bool {
	return hi.LPos == -1 && hi.RPos != -1
}

func (h Holdings) HasCurrencyMap(left func(l IHolding) string, right func(r IHolding) string, ih ...IHolding) map[string]HasInfo {
	mb := make(map[string]HasInfo, len(h))
	for index, v := range h {
		mb[strings.ToLower(left(v))] = HasInfo{
			LPos: index,
			RPos: -1,
		}
	}
	for index, v := range ih {
		key := strings.ToLower(right(v))
		if val, ok := mb[key]; ok {
			if val.RPos != -1 {
				panic("conflicting entries")
			}
			val.RPos = index
			mb[key] = val
		} else {
			if len(key) > 0 {
				mb[key] = HasInfo{
					LPos: -1,
					RPos: index,
				}
			}
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
