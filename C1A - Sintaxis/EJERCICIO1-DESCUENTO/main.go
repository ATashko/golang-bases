package main 

import

	"fmt"


 /*Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos.
 Para ello necesitan una aplicación que les permita calcular el descuento basándose 
 en dos variables: su precio y el descuento en porcentaje. La tienda espera obtener 
 como resultado el valor con el descuento aplicado y luego imprimirlo en consola.*/

func porcentaje (precio float64, descuento float64) float64 {

	per := precio * descuento / 100 

	total := precio - per

	return total

}

func main () {

	fmt.Println(porcentaje(15.45 , 9.5))
	 
}

