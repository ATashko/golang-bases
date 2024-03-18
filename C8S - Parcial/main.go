package main

import (
	"desafio-go-bases/internal/tickets"
	"fmt"
	//"time"
)



func main() {


	ruta := "./tickets.csv"

	/*chTickets := make(chan []tickets.Ticket)
	chError := make(chan error)

	go func() {
		t, err := tickets.ArchivoCSV(ruta)
		if err != nil {
			chError <- err
			return
		}
		chTickets <- t
	}()

	select {
	case tickets := <-chTickets:
		fmt.Println("Tickets leídos:", tickets)
	case err := <-chError:
		fmt.Println("Error al leer el archivo:", err)
	}*/

	tickets, err := tickets.ArchivoCSV(ruta)

	if err != nil {
		fmt.Println("Error al obtener los tickets desde el archivo CSV:", err)
		return
	}

	

	destino := "Paris"

	chTotal := make(chan int)
	chError := make(chan error)

	go func() {
		total, err := tickets.GetTotalTickets(destino, tickets) //no puede reconocer el metodo
		if err != nil {
			chError <- err
			return
		}
		chTotal <- total
	}()

	select {
	case total := <-chTotal:
		fmt.Printf("Total de tickets para %s: %d\n", destino, total)
	case err := <-chError:
		fmt.Println("Error al contar tickets:", err)
	} 

	

		
}


	


	
	//total, err := tickets.GetTotalTickets("Brazil")

