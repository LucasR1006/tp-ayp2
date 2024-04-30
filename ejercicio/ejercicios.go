package ejercicio

import "fmt"

var lista []*Ejercicio

// Definición de la estructura de Ejercicio
type Ejercicio struct {
	Nombre         string
	Descripcion    string
	Tipo           []string
	Dificultad     string
	TiempoEstimado int // en segundos
	Calorias       int
	Puntos         int
}

//Funciòn para crear un ejercicio nuevo por el usuario

func CrearEjercicio() *Ejercicio {
	fmt.Println("Creando Ejercicio")

	return nil
}
