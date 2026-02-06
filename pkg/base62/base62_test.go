package base62

import (
	"fmt"
	"strings"
	"testing"
)

// TestGenerateCode_Structure verifica que el codigo cumpla las bases
func TestGenerateCode_Structure(t *testing.T) {
	code,_ := GenerateCode()

	// verifico longitud
	if len(code) != 6 {
		t.Errorf("Se esperaba largo 6, se obtuvo %d para el código: %s", len(code), code)
	}

	// verifico que los carateres pertenezcan al alfabeto
	for _, char := range code {
		if !strings.ContainsRune(alphabet, char) {
			t.Errorf("Caracter inválido encontrado: %c en el código %s", char, code)
		}
	}
}

// TestGenerateCode_Uniqueness verifica que no genere duplicados en una muestra grande
func TestGenerateCode_Uniqueness(t *testing.T) {
	samples := 15 // Probamos con 10 mil codigos
	seen := make(map[string]bool)

	for i := 0; i < samples; i++ {
		code,_ := GenerateCode()
		fmt.Println(code)
		
		// Si el codigo ya existe en el mapa, fallamos el test
		if seen[code] {
			t.Fatalf("Colision detectada después de %d iteraciones. El codigo %s se repitio.", i, code)
		}
		
		// Guardamos el codigo en el mapa
		seen[code] = true
	}
}

// Benchmark mide que tan rapida es tu funcion
func BenchmarkGenerateCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateCode()
	}
}