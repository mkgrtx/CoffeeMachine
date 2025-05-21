package main

import (
	"fmt"
)

var (
	// water, milk, bean, money
	expresso = [4]int{250, 0, 16, 4}
	latte    = [4]int{350, 75, 20, 7}
	cappu    = [4]int{200, 100, 12, 6}
)

var (
	dWater int = 400
	dMilk  int = 540
	dBean  int = 120
	dCup   int = 9
	dmoney int = 550
)

func main() {
loop:
	for {
		fmt.Println("Write action (buy, fill, take, remaining, sing, exit): ")
		var action string
		fmt.Scanln(&action)
		switch action {
		case "buy":
			buy()
		case "fill":
			fill()
		case "take":
			take()
		case "sing":
			sing()
		case "remaining":
			remain(dWater, dMilk, dBean, dCup, dmoney)
		case "exit":
			//os.Exit(0)
			break loop
		}

	}
}

func sing() {
	fmt.Println("La La La ....\nLa La La.....\nLa La La....")
}

func buy() {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino: , back")
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		makeEspresso()
	case "2":
		makeLatte()
	case "3":
		makeCappu()
	case "back":
		main()
	}
}

func checkIfEnough(water, milk, bean int) bool {
	if water > dWater {
		println("Sorry, not enough water!")
		return false
	}
	if milk > dMilk {
		println("Sorry, not enough milk!")
		return false
	}
	if bean > dBean {
		println("Sorry, not enough bean!")
		return false
	}
	if dCup < 1 {
		println("Sorry, not enough cup!")
		return false
	}
	println("I have enough resources, making you a coffee!")
	return true

}

func makeCappu() {
	if checkIfEnough(cappu[0], cappu[1], cappu[2]) {
		dWater -= cappu[0]
		dMilk -= cappu[1]
		dBean -= cappu[2]
		dmoney += cappu[3]
		dCup--
	}
}

func makeLatte() {
	if checkIfEnough(latte[0], latte[1], latte[2]) {
		dWater -= latte[0]
		dMilk -= latte[1]
		dBean -= latte[2]
		dmoney += latte[3]
		dCup--
	}

}

func makeEspresso() {
	if checkIfEnough(expresso[0], expresso[1], expresso[2]) {
		dWater -= expresso[0]
		dMilk -= expresso[1]
		dBean -= expresso[2]
		dmoney += expresso[3]
		dCup--
	}
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

}

func take() {
	fmt.Printf("I gave you $%d\n", dmoney)
	dmoney = 0
}

func remain(dWater int, dMilk int, dBean int, dCup int, dmoney int) {

	fmt.Printf(`The coffee machine has:
%d ml of water
%d ml of milk
%d g of coffee beans
%d disposable cups
$%d of money

`, dWater, dMilk, dBean, dCup, dmoney)

}
