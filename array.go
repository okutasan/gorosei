package main

import "fmt"

func main() {
	var names [3]string
	names[0]= "Okta"
	names[1]= "Arif"
	names[2]= "Cahyawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [4]int{
		90,
		30,
		40,
		70,
	}
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))
}