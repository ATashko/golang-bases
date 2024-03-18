package main

import (
	"encoding/json" 

	"github.com/gin-gonic/gin"

	"log" 
)

/*Lista de productos
Vamos a crear una aplicación web con el framework Gin que tenga un endpoint /productos que responda con una lista de productos.
Crear una estructura Productos con los valores:
Id
Nombre
Precio
Stock
Codigo
Publicado
FechaDeCreacion
Crear un archivo productos.json con al menos seis productos (deben seguir la misma estructura ya mencionada).
Leer el archivo productos.json.
Imprimir la lista de productos por consola.
Imprimir la lista de productos a través del endpoint en formato JSON. El endpoint deberá ser de método GET.
*/

type productos struct {
	Id  			int 	`json:"Id"`
	Nombre 			string 	`json:"Nombre"`
	Precio 			float64 `json:"Precio"`
	Stock 			int 	`json:"Stock"`
	Codigo 			int 	`json:"Codigo"`
	Publicado 		bool 	`json:"Publicado"`
	FechaDeCreacion string 	`json:"FechaDeCreacion"`
}
func main(){

	router := gin.Default()

	jsonProducto := `[
    {"Id": 1, "Nombre": "Vaso", "Precio": 21.2, "Stock": 21, "Codigo": 12, "Publicado": true, "FechaDeCreacion": "02-02-2024"},
    {"Id": 2, "Nombre": "Plato", "Precio": 21.2, "Stock": 21, "Codigo": 13, "Publicado": true, "FechaDeCreacion": "05-02-2024"},
    {"Id": 3, "Nombre": "Pocillo", "Precio": 21.2, "Stock": 21, "Codigo": 14, "Publicado": true, "FechaDeCreacion": "04-02-2024"},
    {"Id": 4, "Nombre": "Taza", "Precio": 21.2, "Stock": 21, "Codigo": 15, "Publicado": true, "FechaDeCreacion": "12-02-2024"},
    {"Id": 5, "Nombre": "Copa", "Precio": 21.2, "Stock": 21, "Codigo": 16, "Publicado": true, "FechaDeCreacion": "09-02-2024"},
    {"Id": 6, "Nombre": "Jarra", "Precio": 21.2, "Stock": 21, "Codigo": 17, "Publicado": true, "FechaDeCreacion": "07-02-2024"}
	]`

	var productosList []productos


	if err := json.Unmarshal([]byte(jsonProducto), &productosList); err != nil {
		log.Fatal(err)
	}

    router.GET("/productos", func(c *gin.Context){
		c.JSON(200, gin.H{
			"producto" : productosList,
		})
	})

	

	router.Run(":8081")
}

