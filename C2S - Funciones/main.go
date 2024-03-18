/*Una empresa marinera necesita calcular el salario de sus empleados basándose
en la cantidad de horas trabajadas por mes y la categoría.
Categoría C: su salario es de $1.000 por hora.
Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos
trabajados por mes y la categoría, y que devuelva su salario.*/

package main 

import "fmt"

var m float64
const c = 1000.0
const b = 1500.0
const a = 3000.0

func salario (m float64, ctg string) float64 {
	var sTotal float64

	if ctg == "C" {
		sTotal = c * m / 60
	}else if ctg == "B" {
		s := b * m / 60
		sTotal = s + s * 0.20
	}else if ctg == "A" {
		s := a * m / 60
		sTotal = s + s * 0.50
	}
	return sTotal

}

func main()  {
	m = 9000
	ctg := "A"
	fmt.Println(salario(m, ctg))	
}