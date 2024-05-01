package rutina

import (
	"fmt"

	ej "github.com/tp/ejercicio"
)

var Rutinas []*Rutina

// Definición de la estructura de Rutina
type Rutina struct {
	Nombre     string
	Ejercicios []*ej.Ejercicio
	Duracion   int // en minutos
	Calorias   int
	Tipo       map[string]int
	Dificultad string
}

// Función para eliminar una rutina del slice de rutinas
func EliminarRutina(nombre string) {
	for i, rutina := range Rutinas {
		if rutina.Nombre == nombre {
			Rutinas = append(Rutinas[:i], Rutinas[i+1:]...)
			fmt.Println("Rutina eliminada exitosamente.")
			return
		}
	}
	fmt.Println("No se encontró ninguna rutina con el nombre especificado.")
}

// Función para consultar una rutina del slice de rutinas
func ConsultarRutina(nombre string) *Rutina {
	for _, rutina := range Rutinas {
		if rutina.Nombre == nombre {
			return rutina
		}
	}
	fmt.Println("No se encontró ninguna rutina con el nombre especificado.")
	return nil
}

// Función para crear una rutina predefinida y agregarla al slice de rutinas
func CrearRutinaPredefinida() *Rutina {
	var nombre string
	fmt.Println("Nombre de la rutina: ")
	fmt.Scanln(&nombre)

	ejerciciosPredefinidos := ej.ListaPredefinida()
	rutina := crearRutina(ejerciciosPredefinidos)
	Rutinas = append(Rutinas, rutina)
	rutina.Nombre = nombre
	return rutina
}

func CrearRutinaEspartana() *Rutina {
	var nombre string
	fmt.Println("Nombre de la rutina predefinida: ")
	fmt.Scanln(&nombre)

	ejerciciosPredefinidos := ej.EntrenamientoEspartano()
	rutina := crearRutina(ejerciciosPredefinidos)
	Rutinas = append(Rutinas, rutina)
	rutina.Nombre = nombre
	return rutina
}

func CrearRutinaDinamica() *Rutina {
	fmt.Println("Ingrese el nombre de la rutina Dinamica: ")
	var nombre string
	fmt.Scanln(&nombre)

	var cantidadEjercicios int
	fmt.Print("Ingrese la cantidad de ejercicios: ")
	fmt.Scanln(&cantidadEjercicios)

	var ejercicios []*ej.Ejercicio
	for i := 0; i < cantidadEjercicios; i++ {
		ejercicio := ej.CrearEjercicio()
		ejercicios = append(ejercicios, ejercicio)
	}

	rutina := crearRutina(ejercicios)
	rutina.Nombre = nombre
	Rutinas = append(Rutinas, rutina)

	return rutina
}

func crearRutina(ejercicios []*ej.Ejercicio) *Rutina {
	nuevaRutina := &Rutina{
		Ejercicios: ejercicios,
	}

	for _, ejercicio := range ejercicios {
		nuevaRutina.Duracion += ejercicio.Tiempo
		nuevaRutina.Calorias += ejercicio.Calorias
	}

	nuevaRutina.calcularDificultad()

	return nuevaRutina
}

func (r *Rutina) calcularDificultad() {
	var dificultadFacil, dificultadMedia, dificultadDificil int

	for _, ejercicio := range r.Ejercicios {
		switch ejercicio.Dificultad {
		case "facil":
			dificultadFacil++
		case "medio":
			dificultadMedia++
		case "dificil":
			dificultadDificil++
		}
	}

	if dificultadDificil > dificultadMedia && dificultadDificil > dificultadFacil {
		r.Dificultad = "dificil"
	} else if dificultadMedia > dificultadDificil && dificultadMedia > dificultadFacil {
		r.Dificultad = "medio"
	} else {
		r.Dificultad = "facil"
	}
}

func MostrarRutina(rutina *Rutina) {
	fmt.Println("Nombre de la Rutina:", rutina.Nombre)
	fmt.Println("Duración Total:", rutina.Duracion, "minutos")
	fmt.Println("Calorías Quemadas:", rutina.Calorias)
	fmt.Println("Dificultad:", rutina.Dificultad)
	fmt.Println("Ejercicios:")
	for _, ejercicio := range rutina.Ejercicios {
		fmt.Println("Nombre:", ejercicio.Nombre)
		fmt.Println("Descripción:", ejercicio.Descripcion)
		fmt.Println("Tiempo Estimado:", ejercicio.Tiempo, "segundos")
		fmt.Println("Calorías:", ejercicio.Calorias)
		fmt.Println("Dificultad:", ejercicio.Dificultad)
		fmt.Println()
	}
}
