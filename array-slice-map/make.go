package main

import "fmt"

type floatMap map[string]string

func (m floatMap) print() {
	fmt.Println(m)
}

func main() {
	usernames := make([]string, 2, 5)

	fmt.Println(usernames)
	usernames[0] = "alice"
	usernames[1] = "bob"

	usernames = append(usernames, "foo")
	usernames = append(usernames, "bar")

	fmt.Println(usernames)

	courseRatings := make(floatMap, 3) // helps in memory allocation. no other advantage as such

	courseRatings["Go"] = "5"
	courseRatings["Python"] = "4"
	courseRatings["Ruby"] = "3"

	courseRatings.print()

	for index, value := range usernames {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}
