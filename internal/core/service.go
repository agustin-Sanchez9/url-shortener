package core

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)


var blacklist = [1]string{"domain.com"} // ejemplo de lista de blacklist. propenso a cambiar en un futuro
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

	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		return true
	}
	return false
}

// verifica que la url no este dentro de la backlist. Esta solucion es bastante simplista, verificar una mejor en un futuro
func CheckBlacklist(input string) bool {

	for i := range len(blacklist) {
		if strings.Contains(input,blacklist[i]) {
			return false
		}
	}
	return true
}


// Validamos existencia de la url para evitar cargar datos muertos en la db
func ValidateInputExistance(input string) bool {

	// si no se recibe respuesta en 'responseTimeout' se presume la no existencia o inaccesibilidad del input
	client := http.Client{
		Timeout: responseTimeout * time.Millisecond,
	}

	// usamos metodo head() en lugar de get() para no recibir el body de la url
	response, err := client.Head(input)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// supuesta buena practica cerrar el body aunque este vacio
	defer response.Body.Close()

	// menor a 200 son "informational responses" y mayor o igual a 400 son "error responses"
	if response.StatusCode > 199 && response.StatusCode < 400 {
		return true
	}
	return false
}


