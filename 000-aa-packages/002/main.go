package main

import (
	"fmt"
	"mymodule/000-aa-packages/002/mypackage"
)

func main() {
	i, _ := 3, 4
	fmt.Println("from 002", i)
	mypackage.PrintHello()
}
