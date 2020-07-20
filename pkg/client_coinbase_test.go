package pkg

import (
	"context"
	"os"
	"testing"

	"github.com/will7200/go-crypto-sync/pkg/coinbase"
)

func TestCoinbaseAPI(t *testing.T) {
	config := coinbase.NewConfiguration()
	config.Debug = true
	config.APIKey = coinbase.APIKey{
		Key:    os.Getenv("COINBASE_API_KEY"),
		Secret: os.Getenv("COINBASE_API_SECRET"),
	}
	client := coinbase.NewAPIClient(config)
	ctx := context.Background()
	_, _, _ = client.UsersApi.CurrentUser(ctx).Execute()
}
