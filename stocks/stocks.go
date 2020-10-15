package stocks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Stock representa a estrutura recebida da API
//
// Exemplo: curl https://mfinance.com.br/api/v1/stocks/petr4
//
// {
// 	"change":-3,
// 	"closingPrice":21.65,
// 	"eps":-0.95,
// 	"high":21.56,
// 	"lastPrice":21,
// 	"lastYearHigh":31.24,
// 	"lastYearLow":10.85,
// 	"low":20.71,
// 	"marketCap":280142730000,
// 	"name":"Petroleo Brasileiro SA Petrobras Preference Shares",
// 	"pe":0,
// 	"priceOpen":21.49,
// 	"shares":5602042788,
// 	"symbol":"PETR4",
// 	"volume":65162400,
// 	"volumeAvg":83971577,
// 	"sector":"Petróleo. Gás e Biocombustíveis",
// 	"subSector":"Petróleo. Gás e Biocombustíveis",
// 	"segment":"Exploração. Refino e Distribuição"
// }
type Stock struct {
	Name         string
	Symbol       string
	High         float64
	Low          float64
	LastPrice    float64
	ClosingPrice float64
	PriceOpen    float64
}

func GetCotacao(symbol string) (Stock, error) {
	url := "https://mfinance.com.br/api/v1/stocks/" + symbol
	response, err := getHttpBody(url)
	if err != nil {
		return Stock{}, errors.Wrapf(err, "Falha ao buscar cotação", "symbol", symbol)
	}

	j := Stock{}
	err = json.Unmarshal([]byte(response), &j)
	if err != nil {
		return Stock{}, errors.Wrapf(err, "Falha ao converter cotação em JSON, stockData:", response)
	}

	return j, nil
}

func getHttpBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrapf(err, "Falha ao fazer requisição para a url %s", url)
	}
	defer resp.Body.Close()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "Falha ao ler resposta da url %s", url)
	}

	response := strings.TrimSpace(string(bts))
	return response, nil
}
