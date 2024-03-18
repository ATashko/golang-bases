/*Ejercicio 1 - Listado de nombres
Una profesora de la universidad quiere tener un listado con todos sus estudiantes. 
Es necesario crear una aplicación que contenga dicha lista. 
Sus estudiantes son: Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, 
Federico, Hernán, Leandro, Eduardo, Duvraschka.
Luego de dos clases, se sumó un estudiante nuevo. 
Es necesario agregarlo al listado, sin modificar el código que escribiste inicialmente. 
El nombre de la nueva estudiante es Gabriela.
*/

package main

import "fmt"

var estudiantes[]string

func agregar(estudiantes []string, value string) []string {
   return append(estudiantes, value)
} 

func main ()  {
  
	estudiantes = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	
	fmt.Println("Estudiantes:",estudiantes[1: ])

	fmt.Println(append(estudiantes, "Gabriela"))

}