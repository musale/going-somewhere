package main

func main() {
	Repeat("a", 7)
}

// Repeat a character repeat number of times
func Repeat(str string, repeat int) string {
	var rpt string
	for idx := 0; idx < repeat; idx++ {
		rpt += str
	}
	return rpt
}
