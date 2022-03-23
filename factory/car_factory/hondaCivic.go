package main

type hondaCivic struct {
	car
}

func newHondaCivic() iCar {
	return &hondaCivic {
		car: car{
			name: "hondaCivic",
			horsePower: 180,
		},
	}

}