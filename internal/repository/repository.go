package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/aiviaio/go-binance/v2"
	"github.com/aiviaio/go-binance/v2/futures"
)

type BinanceRepository struct {
	client        *binance.Client
	futuresClient *futures.Client
}

func NewRepository(client *binance.Client, futureClient *futures.Client) *BinanceRepository {
	return &BinanceRepository{
		client:        client,
		futuresClient: futureClient,
	}
}

// -1 - all pairs, else - count of pairs
func (r *BinanceRepository) GetPairs(ctx context.Context, count int) ([]string, error) {
	pairs, err := r.client.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return nil, err
	}

	if pairs == nil || pairs.Symbols == nil {
		return nil, fmt.Errorf("no pairs")
	}

	if len(pairs.Symbols) < count {
		return nil, fmt.Errorf("not enough pairs")
	}

	if len(pairs.Symbols) < count {
		count = len(pairs.Symbols)
	}

	strPairs := make([]string, count)

	i := 0
	for _, pair := range pairs.Symbols {
		if strings.HasSuffix(pair.Symbol, "USDT") && pair.IsMarginTradingAllowed {
			strPairs[i] = pair.Symbol
			i++
		}

		if i >= count {
			break
		}
	}

	return strPairs, nil
}

func (r *BinanceRepository) GetPrice(ctx context.Context, pair string) (string, error) {
	prices, err := r.futuresClient.NewListPricesService().Symbol(pair).Do(ctx)

	if err != nil {
		return "", err
	}

	if len(prices) == 0 {
		return "", fmt.Errorf("no prices for pair %s", pair)
	}

	return prices[0].Price, nil
}
