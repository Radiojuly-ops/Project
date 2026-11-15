package main

import "fmt"

func main() {
	numb := 0
	fmt.Scan(&numb)
	Nubmer(numb)
	Chislo(numb)
}

func Nubmer(a int) {
	if a == 0 {
		fmt.Println("Number is null")
	} else if a > 0 {
		fmt.Println("Number is Positive")
	} else if a < 0 {
		fmt.Println("Number is Negative")
	}
}

func Chislo(a int) {
	for i := 0; i < a; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
