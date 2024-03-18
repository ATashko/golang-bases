/*Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros
y devuelva el promedio. No se pueden introducir notas negativas.*/

package main

import "fmt"

var notas[] int
func promedio (notas ...int) float64 {
	var sumatoria float64
	for _, notas := range
	notas {
		sumatoria += float64(notas)
	}

	return sumatoria / float64(len(notas))

}

func main ()  {

	notas := []int{2, 5, 8, 9, 10, 15}
	 fmt.Printf("El promedio de las notas del estudiante es: %.2f ", promedio(notas...))
	
}

