package internal

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)


var blacklist = [1]string{"domain.com"}
const maxSize int = 800 // medido en caracteres
const responseTimeout time.Duration = 2000 // medido en ms


// verifica que el tamaño del input del usuario sea menor a 'maxSize'
// tambien se validaria en el front como una primera medida de seguridad
func ValidateInputSize(input string) bool {

	if len(input) > maxSize {
		return false
	}
	return true
}

// verifica si el formato de la url es valido ('http://' o 'https://')
func ValidateInputFormat(input string) bool {

	if !strings.HasPrefix(input, "http://") {
		return false
	}
	if !strings.HasPrefix(input, "https://") {
		return false
	}
	return true
}

// verifica que la url no este dentro de la backlist
func CheckBlacklist(input string) bool {

	for i := 0; i < len(blacklist); i++ {
		if !strings.Contains(input,blacklist[i]) {
			return false
		}
	}
	return true
}

// Validamos existencia de la url para evitar cargar datos muertos en la db
func ValidateInputExistance(input string) bool {

	client := http.Client{
		Timeout: responseTimeout * time.Millisecond,
	}

	// uso head en lugar de get para evitar que se transfiera el body y perder tiempo.
	response, err := client.Head(input)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// supuesta buena practica cerrar el body aunque este vacio
	defer response.Body.Close()

	// menor 200 son "informational responses" y mayor o igual a 400 son "error responses"
	if response.StatusCode < 200 && response.StatusCode >= 400 {
		return false
	}
	return true
}


