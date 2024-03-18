package main

/*Ejercicio: Hola mundo
Vamos a crear una aplicación web con el framework Gin que tenga un endpoint 
/hola-mundo que responda con un mensaje.
Tener en cuenta que:
El endpoint deberá ser de método GET.
La respuesta deberá ser recibida en formato JSON. Ejemplo:
{
    "mensaje": "¡Hola Mundo!"
}
*/

import "github.com/gin-gonic/gin"

func main () {

	//creacion del router
	router := gin.Default()

	//sintaxis de función handler es router.GET("endpoint, Handler")
	//endpoint es la ruta relativa en este caso es /hola-mundo
	//y handler es la función que toma *gin.Context como argumento
	
	router.GET("/hola-mundo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "¡Hola Mundo!",
		})
	})

	router.Run()
}

