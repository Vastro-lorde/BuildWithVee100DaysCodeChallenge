package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type reciept struct {
	user_name string
	items     map[string]float64
	tip       float64
	createdAt string
	created   time.Time
}

func recieptMaker(user_name string) reciept {
	newReceipt := reciept{
		user_name: user_name,
		items:     map[string]float64{},
		tip:       0,
		created:   time.Now(),
		createdAt: time.Now().String(),
	}
	return newReceipt
}

func (r *reciept) totalCost() float64 {
	var total float64 = 0
	for _, v := range r.items {
		total += v
	}
	return total
}

func (r *reciept) printFullList() {
	printer := fmt.Sprintf("%v Invoice on : %-25v \n", r.user_name, r.created.Month())
	for k, v := range r.items {
		printer += fmt.Sprintf("%-15v %v\n", k+":", v)
	}
	fmt.Println(printer)
}

func createReciept() reciept {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Create a new biil,\n enter bill name:")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	name = strings.TrimSpace(name)
	newReceipt := recieptMaker(name)
	return newReceipt
}
