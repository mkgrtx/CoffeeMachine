package main

import (
	"fmt"
)

var (
	// water, milk, bean
	espresso = []int{250, 0, 16, 4}
	latte    = []int{350, 75, 20, 7}
	cappu    = []int{200, 100, 12, 6}
)

var (
	dWater int = 400
	dMilk  int = 540
	dBean  int = 120
	dCup   int = 9
	dmoney int = 550
)

func main() {
	status(dWater, dMilk, dBean, dCup, dmoney)

	fmt.Println("Write action (buy, fill, take): ")
	var action string
	fmt.Scanln(&action)
	switch action {
	case "buy":
		buy()
	case "fill":
		fill()
	case "take":
		take()
	}

}

func buy() {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino: ")
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		makeEspresso()
	case "2":
		makeLatte()
	case "3":
		makeCappu()
	}
}

func makeCappu() {
	dWater -= cappu[0]
	dMilk -= cappu[1]
	dBean -= cappu[2]
	dmoney += cappu[3]
	dCup--
	status(dWater, dMilk, dBean, dCup, dmoney)
}

func makeLatte() {
	dWater -= latte[0]
	dMilk -= latte[1]
	dBean -= latte[2]
	dmoney += latte[3]
	dCup--
	status(dWater, dMilk, dBean, dCup, dmoney)
}

func makeEspresso() {
	dWater -= espresso[0]
	dMilk -= espresso[1]
	dBean -= espresso[2]
	dmoney += espresso[3]
	dCup--
	status(dWater, dMilk, dBean, dCup, dmoney)
}

func fill() {
	var water, milk, bean, cup int
	fmt.Println("Write how many ml of water you want to add: ")
	fmt.Scanln(&water)
	fmt.Println("Write how many ml of milk you want to add: ")
	fmt.Scanln(&milk)
	fmt.Println("Write how many grams of coffee beans you want to add: ")
	fmt.Scanln(&bean)
	fmt.Println("Write how many disposable cups you want to add: ")
	fmt.Scanln(&cup)

	dWater += water
	dMilk += milk
	dBean += bean
	dCup += cup

	status(dWater, dMilk, dBean, dCup, dmoney)
}

func take() {
	fmt.Printf("I gave you $%d\n", dmoney)
	dmoney = 0
	status(dWater, dMilk, dBean, dCup, dmoney)
}

func status(dWater int, dMilk int, dBean int, dCup int, dmoney int) {

	fmt.Printf(`The coffee machine has:
%d ml of water
%d ml of milk
%d g of coffee beans
%d disposable cups
$%d of money
`, dWater, dMilk, dBean, dCup, dmoney)

}
