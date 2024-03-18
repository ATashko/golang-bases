/*En la función main, definir una variable salary y asignarle un valor de tipo “int”.
Luego, crear un error personalizado con un struct que implemente Error()
con el mensaje “Error: el salario ingresado no alcanza el mínimo imponible”
y lanzarlo en caso de que salary sea menor a 150.000. De lo contrario,
imprimir por consola el mensaje “Debe pagar impuesto”.*/

package main

import "fmt"

type pError struct {

	message string
}

func (e *pError) Error() string{

	return e.message
} 
	
func main ()  {
	salary := 89000

	if salary <= 150000 {
		err := &pError{"Error: el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(err)
	}else{
		fmt.Println("Debe pagar impuesto")
	}
}