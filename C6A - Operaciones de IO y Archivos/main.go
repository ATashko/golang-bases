/*Una empresa que se encarga de vender productos de limpieza necesita: 
Implementar una funcionalidad para guardar un archivo de texto 
con la información de productos comprados, separados por punto y coma (CSV).
Este archivo debe tener el ID del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
¡Manos al teclado!
*/

package main 

import (
	"fmt"
	"io"
	"os"
)

func main () {

	file, err := os.Create("C:\\Users\\USUARIO\\Documents\\CTD\\Especialización\\BackEndIII-GO\\productos.csv")
	if err != nil {
		fmt.Print("Error al crear el archivo", err)
	}
	defer file.Close()

	product := "ID: 1, precio: 1235.21, cantidad: 65; ID: 2, precio: 1285.41, cantidad: 67"

	_, err = io.WriteString(file, product)
	if err != nil {
		fmt.Println("Error al escribir el archivo")
		return
	}
	 fmt.Println("El archivo fue escrito exitosamente")

}