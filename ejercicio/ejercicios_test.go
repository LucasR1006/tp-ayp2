package ejercicio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test para verificar la creación de un nuevo ejercicio
func TestNuevoEjercicio(t *testing.T) {
	ejercicio := NuevoEjercicio("sentadillas", "sube y baja con el abdomen", "facil", "flexibilidad", 40, 100, 30)

	// Verificar que el ejercicio no sea nil
	if ejercicio == nil {
		t.Error("Se esperaba un ejercicio, pero se obtuvo nil")
	}

	// Verificar que los atributos del ejercicio sean los esperados
	if ejercicio.Nombre != "sentadillas" {
		t.Errorf("Nombre del ejercicio incorrecto. Se esperaba 'sentadillas' pero se obtuvo '%s'", ejercicio.Nombre)
	}
}

func TestObtenerEjercicio(t *testing.T) {
	c := Listado()
	assert.Equal(t, "sentadillas", c[0].Nombre)
}

func TestEliminarEjercico(t *testing.T) {
	c := Listado()
	assert.Equal(t, 5, len(c))
	EliminarEjercico("flexiones")
	assert.Equal(t, 4, len(c))
}
