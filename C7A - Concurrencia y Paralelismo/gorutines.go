package main

import (
	"fmt"
	//"time"
)

/*Se requieren tres estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.
Y se requieren tres funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio por cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total 
(precio por media hora trabajada. Si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.
Las tres se deben ejecutar concurrentemente y, al final, se debe mostrar por pantalla el monto final (sumando el total de las tres).
*/

type Producto struct{
	nombre string
	precio float64
	cantidad float64
}

type Servicio struct {
	nombre string
	precio float64
	mT float64
}

type Mantenimiento struct{
	nombre string
	precio float64
}



func main(){

		
	productos := []Producto{
			{
			nombre: "lavadora",
		 	precio: 125.25,
		 	cantidad:2		},
			{
			nombre: "secadora",
		 	precio: 165.75,
		 	cantidad:3	}	}

	servicios := []Servicio{
			{
			nombre: "instalación de lavadora",
			precio: 100.00,
			mT: 25},
			{
			nombre: "instalación de secadora",
			precio: 150.00,
			mT: 45}	}

	mantenimientos := []Mantenimiento{
			{
			nombre: "mantenimiento de lavadora",
			precio: 150.00	},
			{
			nombre: "mantenimiento de secadora",
			precio: 250.00	} }

	serv := sumarProductos(productos)
	prod := sumarServicio(servicios)
	mant := sumarMantenimiento(mantenimientos)

	ch := make (chan float64)

	go sumarTotales (serv, prod, mant, ch)

	fmt.Println(<- ch)


}

func sumarProductos (slice[]Producto) float64  {
	var total float64
	for _, p := range slice {
		total += p.precio * p.cantidad
	}
	return total
}

func sumarServicio (slice[]Servicio) float64 {
	var total float64
	for _, s := range slice {
		if s.mT < 30 {
			s.mT = 30
		}

		total += s.precio * s.mT
	}

	return total
}

func sumarMantenimiento (slice[]Mantenimiento) float64{
	var total float64
	for _, m := range slice {
		total += m.precio
	}
	return total
}

func sumarTotales (serv float64, prod float64, mant float64, ch chan float64 )  {
	
		
	ch <- serv + prod + mant
}



 