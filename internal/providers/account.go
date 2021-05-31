package providers

import (
	"fmt"
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

func (h *Holdings) AddHolding(hh Holding) {
	*h = append(*h, hh)
}

func (h *Holdings) AddHoldings(hh Holdings) {
	*h = append(*h, hh...)
}

// HasCurrencySymbolName check if holdings account has a currency by symbol name
func (h Holdings) HasCurrencySymbolName(symbol string) bool {
	for _, v := range h {
		if strings.ToLower(symbol) == strings.ToLower(v.SymbolName) {
			return true
		}
	}
	return false
}

// HasCurrencyName check if holdings account has a currency by name
func (h Holdings) HasCurrencyName(name string) bool {
	for _, v := range h {
		if strings.ToLower(name) == strings.ToLower(v.FullName) {
			return true
		}
	}
	return false
}

// MapReduce will return Holdings that have been summed across the same SymbolName
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

// FilterTokens by the provided regular expressions
func (h Holdings) FilterTokens(by string, rexes ...*regexp.Regexp) (final Holdings) {
	final = make(Holdings, 0, len(h))
	var accessor func(holding Holding) string
	switch by {
	case "symbolname", "symbol-name", "symbol":
		accessor = func(h Holding) string {
			return h.CurrencySymbolName()
		}
	case "name", "full-name", "fullname":
		accessor = func(h Holding) string {
			return h.CurrencyName()
		}
	case "shares", "total-shares", "totalshares":
		accessor = func(h Holding) string {
			return h.TotalSharesString()
		}
	default:
		accessor = func(h Holding) string {
			return h.CurrencySymbolName()
		}
	}
	for _, holding := range h {
		keep := 1
		for _, rex := range rexes {
			matched := rex.Match([]byte(accessor(holding)))
			res := 0
			if matched {
				res = 1
			}
			keep &= res
			if keep == 0 {
				continue
			}
		}
		if keep == 1 {
			final.AddHolding(holding)
		}
	}
	return
}

type Existence int

const (
	ExistenceUnknown Existence = iota
	ExistsInOriginationOnly
	ExistsInTargetOnly
	ExistsInBoth
)

// CommonElements contains positioning information about holdings in two lists
type CommonElements struct {
	LPos int
	RPos int
}

// Result returns a enum result of what is the account status across the two lists
func (hi CommonElements) Result() Existence {
	if hi.FoundBoth() {
		return ExistsInBoth
	} else if hi.LeftOnly() {
		return ExistsInOriginationOnly
	} else if hi.RightOnly() {
		return ExistsInTargetOnly
	} else {
		return ExistenceUnknown
	}
}

// FoundBoth when both are not -1
func (hi CommonElements) FoundBoth() bool {
	return hi.LPos != -1 && hi.RPos != -1
}

// LeftOnly when holding found in the Left List or the origination list
func (hi CommonElements) LeftOnly() bool {
	return hi.LPos != -1 && hi.RPos == -1
}

// RightOnly when holding found in the Right List or the target list
func (hi CommonElements) RightOnly() bool {
	return hi.LPos == -1 && hi.RPos != -1
}

// HasCurrencyMap compares holdings with another Holdings account to check for existence
func (h Holdings) HasCurrencyMap(left func(l IHolding) string, right func(r IHolding) string, ih ...IHolding) map[string]CommonElements {
	mb := make(map[string]CommonElements, len(h))
	for index, v := range h {
		mb[strings.ToLower(left(v))] = CommonElements{
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
				mb[key] = CommonElements{
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

// Account Interface to providing holding across different third party accounts
type Account interface {
	IProvider
	// GetHoldings returns all holdings that is has
	GetHoldings() (Holdings, error)
}

// GetAccountProvider gets a provider by name that implements the Account interface
func GetAccountProvider(name string) (Provider, error) {
	providerMu.RLock()
	provider, ok := providerMap[name]
	providerMu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("provider: unknown provider %q (forgotten import?)", name)
	}

	_, ok = provider.(Account)
	if !ok {
		return nil, fmt.Errorf("provider %s: does not implement Account interface", name)
	}

	return provider, nil
}
