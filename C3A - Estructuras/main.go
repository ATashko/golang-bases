/*Una universidad necesita registrar a los estudiantes y generar una funcionalidad 
para imprimir el detalle de los datos de cada uno de ellos, de la siguiente manera:
Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]
Los valores que están en corchetes deben ser reemplazados por los datos brindados 
por los alumnos. Para ello es necesario generar una estructura Alumno con las variables 
Nombre, Apellido, DNI, Fecha y que tenga un método detalle.
*/


package main

import "fmt"

type estudiante struct{
	Nombre string
	Apellido string
	DNI string
	Fecha string
}

func (e estudiante) detalle() {
	fmt.Println("Nombre:", e.Nombre)
	fmt.Println("Apellido:", e.Apellido)
	fmt.Println("DNI:", e.DNI)
	fmt.Println("Fecha:", e.Fecha)
	 
}

func main ()  {
	estudiante := estudiante{
		Nombre : "Alejandra",
		Apellido : "Tashko",
		DNI : "12354dfd",
		Fecha : "02/09/1988"}
	estudiante.detalle()
	
}