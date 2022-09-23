package main /*The package “main” tells the Go compiler that the package
  should compile as an executable program instead of a shared library.*/

import "fmt" /*fmt stands for the Format package.This package is all about formatting input and output.
it is a preprocessor command which tells the compiler to include the files lying in the package.*/

func main() { //The main function in the package “main” will be the entry point of our executable program.
	fmt.Println("hello world")
}
