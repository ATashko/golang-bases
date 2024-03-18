/*Una empresa de redes sociales requiere implementar una estructura
usuarios con funciones que vayan agregando información a la misma.
Para optimizar y ahorrar memoria requieren que la estructura usuarios
ocupe el mismo lugar en memoria para el main del programa y para las funciones.
La estructura debe tener los campos: nombre, apellido, edad, correo y contraseña.
Y deben implementarse las funciones:
cambiarNombre: permite cambiar el nombre y apellido.
cambiarEdad: permite cambiar la edad.
cambiarCorreo: permite cambiar el correo.
cambiarContraseña: permite cambiar la contraseña.
*/

package main

import "fmt"

type usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func (u *usuario) detalles(){
	fmt.Println(u.nombre, u.apellido)
	fmt.Println(u.edad)
	fmt.Println(u.correo)
	fmt.Println(u.contraseña)
}

func (u *usuario) cambiarNombre (newNombre string, newApellido string){
	u.nombre = newNombre
	u.apellido = newApellido
}

func (u *usuario) cambiarEdad (newEdad int){
	u.edad = newEdad
}

func (u *usuario) cambiarCorreo (newCorreo string){
	u.correo = newCorreo
}

func (u *usuario) cambiarContraseña (newContraseña string){
	u.contraseña = newContraseña
}

func main ()  {
	
	u := usuario {
		nombre     : "Alejandra",
		apellido   : "Tashko",
		edad       : 36,
		correo     : "alejandra@correo.com",
		contraseña : "8521"}

	u.detalles()

	u.cambiarNombre("Majo","Tashko")
	fmt.Println("El nuevo nombre es:", u.nombre, u.apellido )

	u.cambiarEdad(7)
	fmt.Println("La nueva edad es:", u.edad)

	u.cambiarCorreo("majo@correo.com")
	fmt.Println("El nuevo correo es:", u.correo)

	u.cambiarContraseña("1234")
	fmt.Println("La contraseña ha sido cambiada exitosamente")

	u.detalles()
}