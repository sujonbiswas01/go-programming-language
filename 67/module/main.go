// main.go
package main

import (
	"restaurant-app/delivery"
	"restaurant-app/food"
	"restaurant-app/kitchen"
	"restaurant-app/payment"

	// third-party-package
	"github.com/fatih/color"
)

func main() {
	sujon := "sujon"
	color.Red("Prints text in cyan.", sujon)

	kitchen.MakeFood("Burger")

	payment.Pay(250)
	delivery.Deliver("Sylhet")
	food.MakeFood("pizza")

}
