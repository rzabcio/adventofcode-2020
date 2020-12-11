package main


func Day1_1(filename string) int {
	for no1 := range inputChInt(filename) {
		for no2 := range inputChInt(filename) {
			if no1+no2 == 2020 {
				return no1 * no2
			}
		}
	}
	return 0
}

func Day1_2(filename string) int {
	for no1 := range inputChInt(filename) {
		for no2 := range inputChInt(filename) {
			for no3 := range inputChInt(filename) {
				if no1+no2+no3 == 2020 {
					return no1 * no2 * no3
				}
			}
		}
	}
	return 0
}
