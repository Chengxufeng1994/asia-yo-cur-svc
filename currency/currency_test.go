package currency

import (
	"os"
	"path"
	"testing"
)

var currencies *Currencies

func TestMain(m *testing.M) {
	currenciesPath := path.Join("..", "currencies.json")
	currencies, _ = LoadCurrencies(currenciesPath)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestConvertor(t *testing.T) {
	tests := []struct {
		name   string
		source string
		target string
		amount string
		answer string
	}{
		{
			name:   "source: 'USD', target: 'JPY', amount:'1525'",
			source: "USD",
			target: "JPY",
			amount: "1525",
			answer: "170496.53",
		},
		{
			name:   "source: 'USD', target: 'TWD', amount:'1525'",
			source: "USD",
			target: "TWD",
			amount: "1525",
			answer: "46427.10",
		},
		{
			name:   "source: 'TWD', target: 'USD', amount:'1525'",
			source: "TWD",
			target: "USD",
			amount: "1525",
			answer: "50.04",
		},
		{
			name:   "source: 'TWD', target: 'JPY', amount:'1525'",
			source: "TWD",
			target: "JPY",
			amount: "1525",
			answer: "5595.23",
		},
		{
			name:   "source: 'JPY', target: 'USD', amount:'1525'",
			source: "JPY",
			target: "USD",
			amount: "1525",
			answer: "13.50",
		},
		{
			name:   "source: 'JPY', target: 'TWD', amount:'1525'",
			source: "JPY",
			target: "TWD",
			amount: "1525",
			answer: "411.08",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := currencies.Convertor(test.source, test.target, test.amount)
			if err != nil {
				t.Errorf("error: %#v", err.Error())
			}
			if result != test.answer {
				t.Errorf("result: %v, expect: %v", result, test.answer)
			}
		})

	}
}
