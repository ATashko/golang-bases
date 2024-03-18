/*Un estudio contable necesita acceder a los datos de sus empleados para poder
realizar distintas liquidaciones.
Para ello, cuentan con todo el detalle necesario en un archivo TXT.
Desarrollar el código necesario para leer los datos de un archivo llamado “customers.txt”.
Sin embargo, debemos tener en cuenta que la empresa no nos ha pasado el archivo a leer por el programa.
Dado que no contamos con el archivo necesario, se obtendrá un error.
En tal caso, el programa deberá arrojar un panic al intentar leer un archivo que no existe,
mostrando el mensaje “El archivo indicado no fue encontrado o está dañado”.
Más allá de eso, deberá siempre imprimirse por consola “Ejecución finalizada”.
*/

package main

import (
	"os"
	"fmt"
	
)

func abrirArchivo(nombreArchivo string) (*os.File, error) {


	file, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, err 
	}
	return file, nil 

}

func main ()  {

	//deberá siempre imprimirse por consola “Ejecución finalizada”.
	defer fmt.Println("Ejecución finalizada")

	file, err := abrirArchivo("archivo.txt")

	//arrojar un panic al intentar leer un archivo que no existe
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado" + err.Error())
	}
	defer file.Close()
	
}