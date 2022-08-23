package massunits

func GToKg(g Grama) Quilograma {
	return Quilograma(g / 1000)
}

func GToOz(g Grama) Onça {
	return Onça(g / 28.35)
}

func GToLb(g Grama) Libra {
	return Libra(g / 453.6)
}

func KgToG(kg Quilograma) Grama {
	return Grama(kg * 1000)
}

func KgToOz(kg Quilograma) Onça {
	return Onça(kg * 35.274)
}

func KgToLb(kg Quilograma) Libra {
	return Libra(kg * 2.205)
}

func OzToG(oz Onça) Grama {
	return Grama(oz * 28.35)
}

func OzToKg(oz Onça) Quilograma {
	return Quilograma(oz / 35.274)
}

func OzToLb(oz Onça) Libra {
	return Libra(oz / 16)
}

func LbToG(lb Libra) Grama {
	return Grama(lb * 453.6)
}

func LbTo(lb Libra) Quilograma {
	return Quilograma(lb / 2.205)
}

func LbToOz(lb Libra) Onça {
	return Onça(lb * 16)
}
