// Pacote tempconv realiza conversões de Celsius para Farenheit.
package tempconv

import "fmt"

/* Exercicio 2.1 - Acrescente tipos, constantes e funções em tempconv para processar temperaturas
 * na escala Kelvin [...] */

type Celsius float64
type Farenheit float64

// -------------------------------
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoillingC     Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Farenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

// ----------------------------------
func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}
