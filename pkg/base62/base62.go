// paquete para la generacion de los codigos de base 62
package base62

import (
	"crypto/rand" // usamos crypto por ser la version de random que es criptograficamente seguro. 'math/rand' podria ser predecido por patrones.
	"math/big"
	"strings"
)

// constante con el alfabeto (base 62)
const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// constante sobre el largo del alfabeto. Tiene que ser int64 para poder pasarselo a big.NewInt
const lenAlphabet int64 = 62
// constante con el largo de los codigos
const lenCode = 5


// Cuando es llamada genera un codigo en base 62 de 'lenCode' caracteres
func GenerateCode() (string, error) {

	var code strings.Builder
	code.Grow(lenCode) // reserva espacio para 6 bytes

	for i := 0; i < lenCode; i++ {
		// elige un numero al azar entre 0 y 61
		num, err := rand.Int(rand.Reader, big.NewInt(lenAlphabet)) // hay que usar big.NewInt porque rand.Int recibe un bigint.

		if err != nil {
			return "", err
		}

		randIndex := num.Int64()

		randChar := alphabet[randIndex] // agarra el caracter en la posicion de randIndex

		code.WriteByte(randChar) // se agrega al constructor
	}

	return code.String(), nil
}