/*Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go
para administrar productos y retornar el valor del precio total.
La empresa tiene tres tipos de productos: Pequeño, Mediano y Grande.
Pero se espera que sean muchos más.
Los costos adicionales para cada uno son:
Pequeño: solo tiene el costo del producto.
Mediano: el precio del producto + un 3% por mantenerlo en la tienda.
Grande: el precio del producto + un 6% por mantenerlo en la tienda
y un adicional de $2500 de costo de envío.
El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.

Se requiere una función factory que reciba el tipo de producto y el precio,
y retorne una interfaz Producto que tenga el método Precio.
Se debe poder ejecutar el método Precio y que el método devuelva el precio total
basándose en el costo del producto y los adicionales, en caso de que los tenga.
*/

package main

import "fmt" 

type producto interface {
	precio() float64 
}

type articulo struct{
	nombre string
	tipo string
	costo  float64
}

func factory (nombre, tipo string , costo float64) producto{

	return &articulo{
		nombre : nombre, 
		tipo : tipo, 
		costo: costo}
}

func (a articulo) precio() float64 {
	var pTotal float64

	if a.tipo == "pequeño" {
		pTotal = a.costo
	}else if a.tipo == "mediano"{
		pTotal = a.costo + a.costo * 0.3 
	}else if a.tipo == "grande"{
		pTotal = a.costo + a.costo * 0.6 + 2500
	}
	return pTotal
}

func main ()  {
	a1 := factory("lampara", "pequeño", 12.3)
	a2 := factory("mesa", "mediano", 42.3) 
	a3 := factory("cama", "grande", 102.3)  
	
	fmt.Println("el precio del articulo", a1.(*articulo).nombre, "es", a1.precio())
	fmt.Println("el precio del articulo", a2.(*articulo).nombre, "es", a2.precio())
	fmt.Println("el precio del articulo", a3.(*articulo).nombre, "es", a3.precio())

}