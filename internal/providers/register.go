package providers

import (
	"fmt"
	"sync"

	"github.com/will7200/go-crypto-sync/internal/common"
	"go.uber.org/zap"
)

var (
	providerMu sync.RWMutex

	providerMap = make(map[string]Provider)
)

// Base Provider interface
type IProvider interface {
	// Returns Friendly Name
	Name() string
}

type Config struct {
	Logger *zap.Logger
	Config *common.Config
}

type Provider interface {
	Open(config Config, params ...interface{}) (IProvider, error)
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
		return nil, fmt.Errorf("provider: unknown provider %q (forgotten import?)", name)
	}
	return provider, nil
}

type ProviderIterator struct {
	currentPosition int
	keys            []string
}

func (it *ProviderIterator) Next() Provider {
	if len(it.keys) > it.currentPosition {
		defer func() {
			it.currentPosition += 1
		}()
		if v, exists := providerMap[it.keys[it.currentPosition]]; exists {
			return v
		}
	}
	return nil
}

func (it *ProviderIterator) HasNext() bool {
	if len(it.keys) > it.currentPosition {
		return true
	}
	return false
}

func GetProviderIterator() *ProviderIterator {
	providerMu.RLock()
	pi := ProviderIterator{0, make([]string, len(providerMap))}
	i := 0
	for k, _ := range providerMap {
		pi.keys[i] = k
		i++
	}
	providerMu.RUnlock()

	return &pi
}
