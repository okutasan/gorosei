package main

import "fmt"

func main()  {
	type NoKTP string
	type Married bool

	var ktpOkta NoKTP = "0555"
	var marriedStatus Married = true
	fmt.Println(ktpOkta)
	fmt.Println(marriedStatus)
}