package mataUang

func PecahanMataUang(harga int) map[string]int {
	pecahan := map[string]int{
		"100.000": 100000,
		"50.000":  50000,
		"20.000":  20000,
		"10.000":  10000,
		"5.000":   5000,
		"2.000":   2000,
		"1.000":   1000,
		"500":     500,
		"200":     200,
		"100":     100,
	}

	result := make(map[string]int)

	for key, value := range pecahan {
		if harga >= value {
			result[key] = harga / value
			harga = harga % value
		}
	}

	return result
}
