package num2word

import "testing"

var samples = []struct {
	amount   float64
	params   []RuMoneyOption
	expected string
}{
	{1, []RuMoneyOption{WithUpperFirst(true)}, "Один рубль 00 копеек"},
	{100.21, nil, "сто рублей 21 копейка"},
	{184, []RuMoneyOption{WithFraction(false)}, "сто восемьдесят четыре рубля"},
	{208676, []RuMoneyOption{WithFraction(false), WithUpperFirst(true)}, "Двести восемь тысяч шестьсот семьдесят шесть рублей"},
	{4702, []RuMoneyOption{WithCurrency(false)}, "четыре тысячи семьсот два"},
}

func Test_RuMoney(t *testing.T) {
	for _, tt := range samples {
		res := RuMoney(tt.amount, tt.params...)
		if res != tt.expected {
			t.Errorf("RuMoney(%.2f): expected '%s', got '%s'", tt.amount, tt.expected, res)
		}
	}
}
