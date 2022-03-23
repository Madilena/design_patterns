package main

import "fmt"

func getCar(carType string) (iCar, error) {
	if carType == "hondaCivic"{
		return newHondaCivic(), nil
	}

	if carType == "volvoSedan"{
		return newVolvoSedan(), nil
	}

	return nil, fmt.Errorf("wrong car passed")
}