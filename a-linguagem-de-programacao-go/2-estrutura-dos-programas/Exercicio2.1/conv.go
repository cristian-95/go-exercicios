package tempconv

/* Exercicio 2.1 - Acrescente tipos, constantes e funções em tempconv para processar temperaturas
 * na escala Kelvin [...] */

// CToF converte uma temperatura em Celsius para Farenheit.
func CToF(c Celsius) Farenheit {
	return Farenheit(c*9/5 + 32)
}

// FToC converte uma temperatura em Farenheit para Celsius.
func FToC(f Farenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// ---------------------------
// KToC converte uma temperatura em Kelvin para Celcius.
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// CToK converte uma temperatura em Celcius para Kelvin.
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// KToF converte uma temperatura em Kelvin para Farenheit.
func KToF(k Kelvin) Farenheit {
	c := KToC(k)
	return CToF(c)
}

// FToK converte uma temperatura em Farenheit para Kelvin.
func FToK(f Farenheit) Kelvin {
	c := FToC(f)
	return CToK(c)
}
