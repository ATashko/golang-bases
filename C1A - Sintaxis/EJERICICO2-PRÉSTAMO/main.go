package main

import "fmt" 

/*Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
El banco tiene ciertas reglas para saber a qué cliente se le puede otorgar: solo le otorga 
préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de 
un año de antigüedad en su trabajo. Dentro de los préstamos que otorga, no les cobrará interés 
a los que su sueldo sea mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de 
acuerdo a cada caso.
Pista: Tu código tiene que poder imprimir al menos tres mensajes diferentes.*/

/*credito: 
+22años
empleados
+1año antiguedad trabajo
interes:
-$100.000*/



func main (){
	edad := 25
	empleado := true
	antiguedad := 2
	sueldo := 158000.54

	prestamo := false
	interes := false

	if edad >= 22 && empleado == true && antiguedad >= 2 {
		prestamo = true
	} else if prestamo == true && sueldo <= 100000 {
		interes = true	
	}

	fmt.Println("Datos del solicitante:", edad, empleado, antiguedad, sueldo)
	fmt.Println("El estado de su préstamo es:", prestamo)
	fmt.Println("Le serán cobrados intereses:", interes)

	
}
