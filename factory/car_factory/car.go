package main

type car struct {
    name  string
    horsePower int
}

func (c *car) setName(name string) {
    c.name = name
}

func (c *car) getName() string {
    return c.name
}

func (c *car) setHorsePower(horsePower int) {
    c.horsePower = horsePower
}

func (c *car) getHorsePower() int {
    return c.horsePower
}