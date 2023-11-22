package main

import (
	"context"
	"fmt"
	"sync"
	"test_task/internal/repository"
	"test_task/internal/service"
	"time"

	"github.com/aiviaio/go-binance/v2"
)

const (
	pairsCount = 5
)

func main() {
	ch := make(chan map[string]string, pairsCount)

	wg := &sync.WaitGroup{}
	startWorkers(ch, wg)

	for i := 0; i < pairsCount; i++ {
		for k, v := range <-ch {
			fmt.Printf("%s %s\n", k, v)
		}
	}

	wg.Wait()
}

func startWorkers(ch chan map[string]string, wg *sync.WaitGroup) {
	repository := repository.NewRepository(binance.NewClient("", ""), binance.NewFuturesClient("", ""))
	service := service.NewService(repository)

	pairs, err := service.GetPairs(context.Background(), pairsCount)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	for _, e := range pairs {
		time.Sleep(time.Second)
		wg.Add(1)
		go func(pair string) {
			price, err := service.GetPrice(context.Background(), pair)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}

			pairAndPrice := map[string]string{pair: price}

			defer wg.Done()

			ch <- pairAndPrice
		}(e)
	}

}
