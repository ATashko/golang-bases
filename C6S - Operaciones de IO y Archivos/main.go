/*
Se necesita un software desarrollado en Go que imprima las estimaciones de lo que generarían,
los productos de cada categoría en un mayorista, en ventas de productos para el hogar.
Para eso se detallarán los pasos para conseguirlo:
Funcionalidad para generar un archivo CSV, llamado categorías.csv.
Cargar el archivo con las categorías.

	Ejemplo:
		001	Electrodomésticos	ListaProductos
		002	Muebles				ListaProductos
		003	Herramientas		ListaProductos
		004	Pinturas			ListaProductos
		005	Aberturas			ListaProductos
		006	Construccion		ListaProductos
		007	Automotor 			ListaProductos
		Etcétera…

Generar para cada una de estas categorías los productos. Estos tendrán como información:
Código
Nombre
PrecioActual
CantidadActual

Insertar al menos cuatro productos.
Generar otro archivo CSV, llamado estimaciones.csv. Este tendrá los resultados
de la suma de todos los productos de cada una de las categorías.

Imprimir todos los estimativos por consola.

Ejemplo:
Categoría				EstimativoPorCategoría
Construcción			60.700
Pinturas 				40.500
Aberturas				55.300
TotalEstimativo 		156.500

¡Éxitos!
*/
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Producto struct {
	codigo          string
	nombre          string
	precioActual    int
	cantidadActual  int
}

func ArchivoCSV(ruta string) ([]Producto, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, fmt.Errorf("no se puede abrir el archivo CSV: %v", err)
	}
	defer archivo.Close()

	reader := csv.NewReader(archivo)

	lineas, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("no se puede leer y parsear el archivo CSV: %v", err)
	}

	var productos []Producto

	for _, linea := range lineas {
    fmt.Println(linea)
    var p Producto
    for _, campo := range linea {
        partes := strings.Split(campo, ":")
        if len(partes) != 2 {
            continue
        }

        nombreCampo := strings.TrimSpace(partes[0])
        valorCampo := strings.TrimSpace(partes[1])

        switch nombreCampo {
        case "Codigo":
            p.codigo = valorCampo
        case "Nombre":
            p.nombre = valorCampo
        case "PrecioActual":
            precio, err := strconv.Atoi(valorCampo)
            if err != nil {
                return nil, fmt.Errorf("error al convertir precio a entero: %v", err)
            }
            p.precioActual = precio
        case "CantidadActual":
            cantidad, err := strconv.Atoi(valorCampo)
            if err != nil {
                return nil, fmt.Errorf("error al convertir cantidad a entero: %v", err)
            }
            p.cantidadActual = cantidad
        }
    }
    productos = append(productos, p)
}

	return productos, nil
}

func archivoEstimaciones() (*os.File, error) {
	return os.Create("estimaciones.csv")
}

func main() {
	file, err := os.Create("categorias.csv")
	if err != nil {
		fmt.Print("Error al crear el archivo", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	categoriasCSV := []string{
		"001;Electrodomesticos;Codigo: E1, Nombre: Estufa, PrecioActual: 800, CantidadActual: 500; Codigo: E2, Nombre: Licuadora, PrecioActual: 80, CantidadActual: 1500; Codigo: E3, Nombre: Aspiradora, PrecioActual: 400, CantidadActual: 800; Codigo: E4, Nombre: Televisor, PrecioActual: 800, CantidadActual: 500",
		"002;Muebles;Codigo: M1, Nombre: Cama, PrecioActual: 1600, CantidadActual: 500; Codigo: M2, Nombre: Mesa, PrecioActual: 80, CantidadActual: 1500; Codigo: M3, Nombre: Sofa, PrecioActual: 1400, CantidadActual: 80; Codigo: M4, Nombre: Escritorio, PrecioActual: 400, CantidadActual: 500",
		"003;Herramientas;Codigo: H1, Nombre: Taladro, PrecioActual: 800, CantidadActual: 1500; Codigo: H2, Nombre: Caladora, PrecioActual: 120, CantidadActual: 1500; Codigo: H3, Nombre: Lijadora, PrecioActual: 900, CantidadActual: 600; Codigo: H4, Nombre: Pulidora, PrecioActual: 1300, CantidadActual: 800",
		"004;Pinturas;Codigo: P1, Nombre: Vinilo interiores, PrecioActual: 100, CantidadActual: 2500; Codigo: P2, Nombre: Vinilo exteriores, PrecioActual: 130, CantidadActual: 2500; Codigo: P3, Nombre: Esmalte mate, PrecioActual: 100, CantidadActual: 1800; Codigo: P4, Nombre: Esmalte brillante, PrecioActual: 100, CantidadActual: 2500",
	}

	for _, fila := range categoriasCSV {
		elementos := strings.Split(fila, ";")
		if err := writer.Write(elementos); err != nil {
			fmt.Println("Error al escribir en el archivo:", err)
			return
		}
	}

	productos, err := ArchivoCSV("./categorias.csv")
	if err != nil {
		log.Fatalf("Error al procesar el archivo CSV: %v", err)
	}

	categorias := []string{"001", "002", "003", "004"}

	estimativosCategoria := make(map[string]int)
	for _, cat := range categorias {
		estimativosCategoria[cat] = 0
	}

	for _, producto := range productos {
		estimativosCategoria[producto.codigo] += producto.precioActual * producto.cantidadActual
	}

	estimacionesFile, err := archivoEstimaciones()
	if err != nil {
		log.Fatalln("error al crear el archivo de estimaciones", err)
	}
	defer estimacionesFile.Close()

	estimacionesWriter := csv.NewWriter(estimacionesFile)
	defer estimacionesWriter.Flush()

	estimacionesWriter.Write([]string{"Categoría", "EstimativoPorCategoría"})
	totalEstimativo := 0
	for categoria, estimativo := range estimativosCategoria {
		estimacionesWriter.Write([]string{categoria, strconv.Itoa(estimativo)})
		totalEstimativo += estimativo
	}
	estimacionesWriter.Write([]string{"TotalEstimativo", strconv.Itoa(totalEstimativo)})

	fmt.Println("Estimativos por Categoría:")
	for categoria, estimativo := range estimativosCategoria {
		fmt.Printf("%s\t%d\n", categoria, estimativo)
	}
	fmt.Printf("Total Estimativo:\t%d\n", totalEstimativo)
}

 