package common

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/will7200/go-crypto-sync/internal/providers"
	"github.com/will7200/go-crypto-sync/pkg/personalcapital"
)

func ReadGlob(p interface{}, filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = gob.NewDecoder(bytes.NewBuffer(b)).Decode(p)
	if err != nil {
		panic(err)
	}
}

func TestAlgo(t *testing.T) {
	var pcHoldings personalcapital.GetHoldingsResponse
	var _holdings providers.Holdings
	ReadGlob(&pcHoldings, "pc_holdings.gob")
	ReadGlob(&_holdings, "coinbase_holdings.gob")
	//for _, pcHolding := range pcHoldings.Holdings {
	//	spew.Dump(_holdings.HasCurrencyName(pcHolding.Ticker))
	//}
	spew.Dump(pcHoldings.Holdings)
	spew.Dump(_holdings)
	spew.Dump(_holdings.HasCurrencyMap(func(l providers.IHolding) string {
		return l.CurrencyName()
	}, func(r providers.IHolding) string {
		return r.CurrencyName()
	}, PCHoldingsToIHoldings(pcHoldings.Holdings)...))

}

func PCHoldingsToIHoldings(h []personalcapital.HoldingsType) []providers.IHolding {
	a := make([]providers.IHolding, len(h))
	for i := 0; i < len(h); i++ {
		a[i] = h[i]
	}
	return a
}
