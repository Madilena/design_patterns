package main

type iCar interface {
    setName(name string)
    setHorsePower(power int)
    getName() string
    getHorsePower() int
}