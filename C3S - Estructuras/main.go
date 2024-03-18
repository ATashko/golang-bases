/*Crear un programa que cumpla los siguientes puntos:
-Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
-Tener un slice global de Product llamado Products, instanciado con valores.
-Dos métodos asociados a la estructura Product: Save(), GetAll(). 
-El método Save() deberá tomar el slice de Products y 
 añadir el producto desde el cual se llama al método. 
-El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
-Una función getById() a la cual se le deberá pasar un INT como parámetro y 
 retorna el producto correspondiente al parámetro pasado. 
-Ejecutar al menos una vez cada método y función definidos desde main().
*/

package main

import "fmt"

type product struct {
	ID int 
	Name string
	Price float64
	Description string
	Category string

}
/*
Un "slice global" se refiere a un slice que se declara fuera de cualquier función,
lo que significa que está disponible para todo el paquete en el que se declara. 
Un slice global se puede acceder y modificar desde cualquier función dentro del mismo paquete.*/
var products = []product{
	{ID: 1, Name: "Cookies", Price: 15.25, Description:"Amazing cookies", Category:"Food"},
	{ID: 2, Name: "Cake", Price: 17.25, Description:"Amazing cake", Category:"Food"},
	{ID: 3, Name: "Bread", Price: 15.24, Description:"Amazing bread", Category:"Food"},
	{ID: 4, Name: "Croissant", Price: 4.25, Description:"Amazing croissant", Category:"Food"}}

func (p *product) save(){
	/*miSlice = append(miSlice, 4)*/

	products = append(products, *p)

}

func getAll(){
	
	fmt.Println(products)
}

func getById(id int) product{
	for _, p := range products{
		if id == p.ID{
			return p
		}
	} 
	return product{}
}

func main ()  {
	
	newProduct := product{ID: 5, Name: "Coke", Price: 17.25, Description:"Amazing coke", Category:"Soda"}
	newProduct.save()

    getAll()

	findProduct := getById(5)
	fmt.Println(findProduct)

}