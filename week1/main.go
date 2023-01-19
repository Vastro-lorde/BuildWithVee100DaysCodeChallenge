package main

import (
	"fmt"
	"strings"
)

func main() {
	var sex string = "male"
	name := "Spirit Seun" //can't declare a variable like this outside a function
	// age := 16;
	// var bags uint8 = 129;
	// var money float32 = 5.5;
	// ages := [3]int{3,4,5}
	// rates := []int{1,2,3,4,5} // a slice is an array without the fixed range

	fmt.Println(strings.Contains(name, "Seun"))
	fmt.Println(strings.Repeat(sex+" ", 5))
	fmt.Println(strings.Index(sex, "le"))

	for i := 0; i < 5; i++ {
		fmt.Println(name, i)
	}

	// fmt.Println(name, sex, age, bags, money);

	// fmt.Println(age);
	// fmt.Println(ages);

	// fmt.Println(rates);

	// rates = append(rates, 9);
	// fmt.Println(rates);

	// firstRates := rates[0:3];
	// fmt.Println(firstRates);

	// secondRates := rates[:4];
	// fmt.Println(secondRates);
}
