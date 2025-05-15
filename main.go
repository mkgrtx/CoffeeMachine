package main

import (
	"fmt"
)

const (
	needWater = 200
	needMilk  = 50
	needBean  = 15
)

func main() {
	//	fmt.Println(`Starting to make a coffee
	//Grinding coffee beans
	//Boiling water
	//Mixing boiled water with crushed coffee beans
	//Pouring coffee into the cup
	//Pouring some milk into the cup
	//Coffee is ready!`)

	//amountWater := numCups * 200
	//amountMilk := numCups * 50
	//amountBean := numCups * 15

	//fmt.Printf("For %d cups of coffee you will need\n"+
	//	"%d ml of water\n"+
	//	"%d ml of milk\n"+
	//	"%d g of coffee beans\n", numCups, amountWater, amountMilk, amountBean)

	var (
		dWater int
		dMilk  int
		dBean  int
	)

	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scanln(&dWater)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scanln(&dMilk)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scanln(&dBean)

	canMakeWater := dWater / needWater
	canMakeMilk := dMilk / needMilk
	canMakeBean := dBean / needBean

	numCanHave := min(canMakeWater, canMakeMilk, canMakeBean)

	fmt.Println("Write how many cups of coffee you will need:")
	var numCups int
	fmt.Scanln(&numCups)

	difCanNeed := numCanHave - numCups
	switch {
	case difCanNeed == 0:
		fmt.Println("Yes, I can make that amount of coffee")
	case difCanNeed < 0:
		fmt.Printf("No, I can make only %d cups of coffee\n", numCanHave)
	case difCanNeed > 0:
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", difCanNeed)
	}
}
