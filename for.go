package main

import "fmt"

func main() {

	println("while loop - loop 1")
	i := 1
	for i <= 3 {
		println(i)
		i++
	}

	println("for loop - loop 2")
	for i := 0; i < 10; i++ {
		println(i)
	}

	println("range loop - loop 3")
	for i := range 3 {
		println(i)
	}

	println("infinite loop - loop 4")

	for {
		println("in loop")
		break
	}

	println("odd numbers loop - loop 5")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	println("finished")
	println(i)
}
