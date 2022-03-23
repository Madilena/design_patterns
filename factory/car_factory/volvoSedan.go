package main

type volvoSedan struct {
	car
}

func newVolvoSedan() iCar {
	return &volvoSedan {
		car: car{
			name: "volvoSedan",
			horsePower: 455,
		},
	}

}