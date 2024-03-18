/*Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo.
Para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario, 
teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y 
si gana más de $150.000 se le descontará, además, un 10 % (27 % en total).
*/
/*func impuesto return impuesto de salario
si sueldo > 50000 -17%
si sielgo > 150000 -27% */

package main

import
"fmt"

var sueldo float64
func impuesto(sueldo float64) float64 {
	var i float64
	if sueldo > 50000 && sueldo < 150000 {
		i = sueldo*0.17
		fmt.Printf("El impuesto a cobrar es: %.2f, y corresponde al 17%%", i)
	} else if sueldo > 150000 {
		i = sueldo*0.27
		fmt.Printf("El impuesto a cobrar es: %.2f, y corresponde al 27%%", i)
	} 
	return i
}
func main (){
	sueldo := 289650.54
	impuesto(sueldo)
} 


