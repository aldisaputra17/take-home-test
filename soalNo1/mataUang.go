package soalNo1

/* No 1. Di Indonesia, ada pecahan mata uang rupiah, yaitu :
* 100.000,-
* 50.000,-
* 20.000,-
* 10.000,-
* 5.000,-
* 2.000,-
* 1.000,-
* 500,-
* 200,-
* 100,-
Buatlah sebuah fungsi untuk menghitung berapa lembar pecahan yang harus dikeluarkan dari input harga (dengan pembulatan ke atas jita punya harga pecahan antara 1 sampai 99)
Input : 145.000
Output:
            {
                 “Rp. 100.000”:1,
                 “Rp. 20.000”:2,
                 “Rp. 5.000”:1,
            }
            Input: 2050
            Ouput:
            {
                 “Rp. 2.000”:1,
                 “Rp. 100”:1,
            }
*/

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
