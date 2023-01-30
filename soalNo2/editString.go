package soalNo2

/* No 2. Anda diminta untuk membuat sebuah function dimana function tersebut
berfungsi untuk menentukan apakah dari dua data string yang diberikan
membutuhkan sekali proses edit atau lebih.
Jika lebih dari sekali proses edit berarti function tersebut akan
mengembalikan response False,
sedangkan jika hanya sekali proses edit maka function tersebut akan
mengembalikan response True. Proses edit di sini dapat berarti melakukan
insert sebuah character, remove sebuah character, atau replace sebuah character.
Contoh
GIVEN INPUT 1 → telkom
GIVEN INPUT 2 → telecom
RESULT → False

GIVEN INPUT 1 → telkom
GIVEN INPUT 2 → tlkom
RESULT → True */

func EditString(val1, val2 string) bool {
	if len(val1) != len(val2) {
		return false
	}
	var edit int
	for i := range val1 {
		if val1[i] != val2[i] {
			edit++
		}
	}
	if edit > 1 {
		return false
	}
	return true
}
