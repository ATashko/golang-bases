package tickets

import (
	
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	//"time"
	"strings"

	
)

type Ticket struct {
	id      int
	nombre  string
	email   string
	destino string
	hora    float64
	precio  float64
}

func obtenerHora(str string) (float64, error) {
	
	parts := strings.Split(str, ":")
	if len(parts) > 0 {
		hourStr := parts[0]

		hour, err := strconv.ParseFloat(hourStr, 64)
		if err != nil {
			return 0, fmt.Errorf("error al convertir la hora a float64: %v", err)
		}
		return hour, nil
	}
	return 0, fmt.Errorf("no se encontró el separador ':' en el string")
}

func ArchivoCSV(ruta string) ([]Ticket, error) {

	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, fmt.Errorf("no se puede abrir el archivo CSV: %v", err)
	}
	defer archivo.Close()

	lector := csv.NewReader(archivo)

	lineas, err := lector.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("no se puede leer y parsear el archivo CSV: %v", err)
	}

	

	var tickets []Ticket

	for _, linea := range lineas {
		id, err := strconv.Atoi(linea[0])
		if err != nil {
			return nil, fmt.Errorf("error al convertir id a entero: %v", err)
		}

		
		hora, err := obtenerHora(linea[4])
		if err != nil {
		return nil, fmt.Errorf("error al convertir hora: %v", err)
		}

		
		monto, err := strconv.ParseFloat(linea[5], 64)
		if err != nil {
			return nil, fmt.Errorf("error al convertir monto a flotante: %v", err)
		}

		
		ticket := Ticket{
			id:      id,
			nombre:  linea[1],
			email:   linea[2],
			destino: linea[3],
			hora:    hora,
			precio:  monto,
		}

		
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}


func GetTotalTickets(destino string, tickets []Ticket) (int, error) {
	total := 0
    
	for _, ticket := range tickets {
		if ticket.destino == destino {
			total++
		}
	}
	return total, nil
}


func GetMadrugada(hora float64, tickets []Ticket) int {
	contadorMadrugada := 0
	for _, ticket := range tickets {
		
		if ticket.hora >= 0 && ticket.hora < 6 {
			contadorMadrugada++
			
		}
	}
	return contadorMadrugada 
	
}

func GetManana(hora float64, tickets []Ticket) int {
	contadorManana := 0
	for _, ticket := range tickets {
		
		if ticket.hora >= 7 && ticket.hora < 12 {
			contadorManana++
			
		}
	}
	return contadorManana 
	
}



/*func GetManana(t time.Time, tickets []Ticket) int {
	contadorManana := 0
	for _, ticket := range tickets {
		t := ticket.hora
		if t.Hour() >= 7 && t.Hour() < 12 {
			contadorManana++
			
		}

	}
	return contadorManana
}

func GetTarde(t time.Time, tickets []Ticket) int {
	contadorTarde := 0
	for _, ticket := range tickets {
		t := ticket.hora
		if t.Hour() >= 13 && t.Hour() < 20 {
			contadorTarde++
		
		}
	}
	return contadorTarde
}

func GetNoche(t time.Time, tickets []Ticket) int {
	contadorNoche := 0
	for _, ticket := range tickets {
		t := ticket.hora
		if t.Hour() >= 20 && t.Hour() <= 23 {
			contadorNoche++
			
		} 
	}
	return contadorNoche
}*/
 



func AverageDestination(destino string, tickets []Ticket) (float64, error) {
	totalDestino, err := GetTotalTickets(destino, tickets)
	if err != nil {
		return 0, err
	}
	totalTickets := len(tickets)

	porcentaje := float64(totalDestino) / float64(totalTickets) * 100.0

	return porcentaje, nil
}