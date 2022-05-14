package main

import "fmt"

func main() {
	var t int
	var a int
	divi := []int{}

	fmt.Scanln(&t)
	r := []string{}
	for i := 0; i < t; i++ {
		fmt.Scanln(&a)
		divi = append(divi, a)
	}
	for i := 0; i < len(divi); i++ {
		if divi[i] <= 1399 {
			r = append(r, "Division 4")
		}
		if 1400 <= divi[i] && divi[i] <= 1599 {
			r = append(r, "Division 3")
		}
		if 1600 <= divi[i] && divi[i] <= 1899 {
			r = append(r, "Division 2")
		}
		if 1900 <= divi[i] {
			r = append(r, "Division 1")
		}

	}
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i])
	}
}
