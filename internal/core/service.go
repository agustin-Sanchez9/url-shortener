package internal

import "strings"




var blacklist = [1]string{"domain.com"}

const maxSize int = 800


// verifica que el tamaño del input del usuario sea menor a 'maxSize'
func ValidateInputSize(input string) bool {

	if len(input) > maxSize {
		return false
	}
	return true
}

// verifica si el formato de la url es valido ('http://' o 'https://')
func ValidateInputFormat(input string) bool {

	if !strings.Contains(input, "http://") {
		return false
	}
	if !strings.Contains(input, "https://") {
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


