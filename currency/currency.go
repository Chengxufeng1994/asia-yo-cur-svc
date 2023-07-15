package currency

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Currencies struct {
	Currencies map[string]ExchangeRate `json:"currencies"`
}

type ExchangeRate map[string]float64

func LoadCurrencies(path string) (*Currencies, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	currencies := Currencies{}
	if err := json.Unmarshal(data, &currencies); err != nil {
		return nil, err
	}

	return &currencies, nil
}

func (curr *Currencies) CheckCurrencyInUse(currency string) error {
	for key := range curr.Currencies {
		if key == currency {
			return nil
		}
	}

	return fmt.Errorf("currency '%s' cannot use", currency)
}

func (curr *Currencies) GetRates(currency string) (ExchangeRate, error) {
	if val, ok := curr.Currencies[currency]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("currency '%s' exchange rate not found", currency)
}

func (curr *Currencies) Convertor(source, target, amount string) (string, error) {
	rates, err := curr.GetRates(source)
	if err != nil {
		return "", err
	}

	var newAmount float64
	if rate, ok := rates[target]; ok {
		oldAmount, _ := strconv.ParseFloat(amount, 64)
		// 四捨五入到小數點第二位
		newAmount = oldAmount * rate
		newAmount = math.Trunc(newAmount*1e2+0.5) * 1e-2
	}

	return fmt.Sprintf("%0.2f", newAmount), nil
}
