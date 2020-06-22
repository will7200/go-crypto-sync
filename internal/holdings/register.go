package holdings

import (
	"fmt"
	"sync"
)

var (
	providerMu sync.RWMutex

	providerMap = make(map[string]Provider)
)

type Provider interface {
	Open(params ...interface{}) (Account, error)
}

func Register(name string, provider Provider) {
	providerMu.Lock()

	defer providerMu.Unlock()

	if provider == nil {

		panic("holding: Registered provider is nil")

	}

	if _, dup := providerMap[name]; dup {

		panic("sql: Register called twice for provider " + name)

	}

	providerMap[name] = provider

}

func GetProvider(name string) (Provider, error) {
	providerMu.RLock()
	provider, ok := providerMap[name]
	providerMu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("sql: unknown provider %q (forgotten import?)", name)
	}
	return provider, nil
}
