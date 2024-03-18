/*Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. 
Según el siguiente map, debemos imprimir la edad de Benjamin.
  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado, también es necesario:
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del map.
*/

/*imprimir la edad de benjamin,
cuantos son mayores de 21
agregar un empleado "Federico": 25
eliminar a Pedro*/

package main

import "fmt"



func main () {

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
  //imprimir la edad de benjamin
  fmt.Println("El empleado Benjamín tiene", employees["Benjamin"], "años")
  //son mayores de 21
  for i, employees := range 
  employees {if employees > 21{fmt.Println("Es mayor de 21 años",i, employees )}}
  //agregar un empleado "Federico": 25
  employees["Federico"]=25
  fmt.Println(employees)
  //eliminar a Pedro
  if employees["Pedro"] != 0 {
  delete(employees, "Pedro")}else{
    fmt.Println("El empleado no existe")
  }
  fmt.Println(employees)

}