/*Repetir el proceso de la ejercitación realizada en clase, 
pero ahora implementando fmt.Errorf() para que el mensaje de error 
reciba por parámetro el valor de salary indicando que no alcanza el mínimo imponible. 
El mensaje mostrado por consola deberá decir: 
“Error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]” 
(siendo [salary] el valor de tipo int pasado por parámetro).*/

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
		err := fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d ", salary)
		fmt.Println(err)
	}else{
		fmt.Println("Debe pagar impuesto")
	}
}