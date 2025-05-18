package main

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func main2() {
	hobbies := [3]string{"gaming", "coding", "sleeping"}

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	featuredHobbies := hobbies[:2]
	fmt.Println(featuredHobbies)

	featuredHobbies = featuredHobbies[1:3]
	fmt.Println(featuredHobbies)

	products := []Product{{1, "A book", 10.99}, {2, "A carpet", 20.00}}
	fmt.Println(products)
	products[1].price = 15.99
	fmt.Println(products)

	products = append(products, Product{3, "A dishwasher", 200.99})
	fmt.Println(products)
}
