// Pacote MassUnits realiza deversas conversões de massa, como Libras para Quilos.
package massunits

import "fmt"

type Onça float64
type Libra float64
type Grama float64
type Quilograma float64

func (ot Onça) String() string {
	return fmt.Sprintf("%g oz", ot)
}

func (lb Libra) String() string {
	return fmt.Sprintf("%g lb", lb)
}

func (g Grama) String() string {
	return fmt.Sprintf("%g g", g)
}

func (kg Quilograma) String() string {
	return fmt.Sprintf("%g kg", kg)
}
