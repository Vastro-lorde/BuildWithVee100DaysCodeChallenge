package main

import "fmt"

// func countOddNumbers(endAt int ) {
// 	for i := 0; i < endAt; i++ {
// 		if i % 2 != 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

func main() {
	seun_reciept := recieptMaker("Seun")
	seun_reciept.items["Biro"] = 45.67
	seun_reciept.items["Laptop"] = 790000.23
	seun_reciept.items["Mifi"] = 50000.7

	fmt.Println(seun_reciept);

	seun_reciept.printFullList()

	// ages := []int{23, 34, 18, 16, 25, 17, 50, 49, 51, 67, 46, 88, 12, 0}
	// fmt.Println(checkModeMean(ages))

	// fmt.Print(age)

	// checkList(ages, checkEven)

	// fmt.Print(ages)

	//variables declaration
	// names := []string{"Seun", "Blessing", "Love"}
	// var sex string = "male"
	// name := "Spirit Seun" //can't declare a variable like this outside a function
	// age := 16;
	// var bags uint8 = 129;
	// var money float32 = 5.5;
	// ages := [3]int{3,4,5}
	// rates := []int{1,2,3,4,5} // a slice is an array without the fixed range

	// fmt.Println(strings.Contains(name, "Seun"))
	// fmt.Println(strings.Repeat(sex+" ", 5))
	// fmt.Println(strings.Index(sex, "le"))

	// for index, value := range names {
	// 	fmt.Println(index, value)
	// }

	// for _, value := range names {
	// 	fmt.Println(value == "Seun")
	// }

	// for _, value := range ages {
	// 	if value < 18 { fmt.Println("Minor")}
	// 	if value > 18 && value < 50 {
	// 		fmt.Println("Adult")
	// 	}
	// 	if value >= 50 { fmt.Println("Close to the grave")}
	// }

	// countOddNumbers(20)

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
