package main

import (
	"fmt"
)

func main() {
	hondaCivic, _ := getCar("hondaCivic")
	printDetails(hondaCivic)

	volvoSedan, _ := getCar("volvoSedan")
	printDetails(volvoSedan)
}

func printDetails(c iCar){
	fmt.Println("car:", c.getName(),"has horsepower", c.getHorsePower())
}