package main

import "fmt"

func main() {

	var a string = "saurabh"
	var b string = "martande"
	var i int = 100
	//var f float64 = 99.9999

	fmt.Print("my name is ", a, " and surname is ", b)
	//1st char of Print method is in upper case
	//because it is exported identifier(a identifiers which is allowed to access it from another
	//package)
	fmt.Print("\n1st line\n") //"\n" for new line
	fmt.Print("2nd line\n")
	//---------------------------------------------------------
	fmt.Println("i=", i)

	//----------------------------------------------------------
	fmt.Printf("%s got %d marts in unit test", a, 10)

}
