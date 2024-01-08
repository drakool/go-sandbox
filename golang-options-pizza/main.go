package main

import "example.com/pizza"

func main() {
    pizza := Pizza.NewPizza(
        Pizza.WithDough("Regular"),
        Pizza.WithSauce("Tomato"),
        Pizza.WithCheese("Mozzarella"),
        Pizza.WithToppings([]string{"Pepperoni", "Olives", "Mushrooms"}),
    )

    println(pizza.Dough)
    println(pizza.Sauce)
    println(pizza.Cheese)
    println(pizza.Toppings)
}
