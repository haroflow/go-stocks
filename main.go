package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/haroflow/go-stocks/stocks"
	"github.com/hashicorp/go-hclog"
)

var (
	logger hclog.Logger

	// Flags
	fDebug = flag.Bool("debug", false, "show debug messages")
)

func main() {
	flag.Parse()

	ativos := flag.Args()
	if len(ativos) < 1 {
		fmt.Println("stocks [-debug] ativo1 [ativo2 ativo3 ...]")
		os.Exit(0)
	}

	logger = hclog.Default()

	if *fDebug {
		logger.SetLevel(hclog.Debug)
	} else {
		logger.SetLevel(hclog.Error)
	}

	// ativos := []string{
	// 	"petr4",
	// 	"azul4",
	// 	"ciel3",
	// 	"hgtx3",
	// }

	var wg sync.WaitGroup
	for _, ativo := range ativos {
		ativo := ativo
		wg.Add(1)

		go func() {
			logger.Info("Pegando dados do ativo", "ativo", ativo)

			cotacao, err := stocks.GetCotacao(ativo)
			if err != nil {
				logger.Error("Falha ao buscar cotação", "ativo", ativo, "error", err)
				os.Exit(1)
			}
			logger.Info("Sucesso", "cotacao", fmt.Sprintf("%#v", cotacao))

			fmt.Printf("%10s: %8.2f\n", cotacao.Symbol, cotacao.LastPrice)
			wg.Done()
		}()
	}

	wg.Wait()
}
