package rutina

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	ej "github.com/tp/ejercicio"
	"github.com/untref-ayp2/data-structures/dictionary"
)

var Rutinas []*Rutina

type Rutina struct {
	Nombre     string
	Ejercicios []*ej.Ejercicio
	Duracion   int // en minutos
	Calorias   int
	Tipo       *dictionary.Dictionary[string, int]
	Dificultad string
}

// Función para eliminar una rutina del slice de rutinas
func EliminarRutina(nombre string) {
	for i, rutina := range Rutinas {
		if rutina.Nombre == nombre {
			Rutinas = append(Rutinas[:i], Rutinas[i+1:]...)
			fmt.Println("Rutina eliminada correctamente.")
			return
		}
	}
	fmt.Println("No se encontró ninguna rutina con ese nombre.")
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
func CrearRutinaBasica() *Rutina {
	var nombre string
	fmt.Println("Nombre de la rutina: ")
	fmt.Scanln(&nombre)

	_, ejerciciosPredefinidos := ej.ListaPredefinida()
	rutina := crearRutina(ejerciciosPredefinidos, "Basica")
	Rutinas = append(Rutinas, rutina)
	rutina.Nombre = nombre
	return rutina
}

func CrearRutinaDificil() *Rutina {
	var nombre string
	fmt.Println("Nombre de la rutina: ")
	fmt.Scanln(&nombre)

	ejerciciosPredefinidos := ej.EntrenamientoDificil()
	rutina := crearRutina(ejerciciosPredefinidos, "esparta")
	Rutinas = append(Rutinas, rutina)
	rutina.Nombre = nombre
	return rutina
}

func CrearRutinaManual() *Rutina {
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

	rutina := crearRutina(ejercicios, "dinamica")
	rutina.Nombre = nombre
	Rutinas = append(Rutinas, rutina)

	return rutina
}

func crearRutina(ejercicios []*ej.Ejercicio, nombre string) *Rutina {
	nuevaRutina := &Rutina{
		Ejercicios: ejercicios, Nombre: nombre,
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
		fmt.Println("Tiempo Estimado:", ejercicio.Tiempo, "minutos")
		fmt.Println("Calorías:", ejercicio.Calorias)
		fmt.Println("Dificultad:", ejercicio.Dificultad)
		fmt.Println("Tipo:", ejercicio.Tipo)
		fmt.Println()
	}
}

func CrearRutinaMaximoPuntaje(nombre string, duracion int, tipo string, lista *dictionary.Dictionary[string, *ej.Ejercicio]) *Rutina {
	var disponibles []*ej.Ejercicio
	for _, eje := range lista.Values() {
		if eje.Tipo.Keys()[0] == tipo {
			disponibles = append(disponibles, eje)
		}
	}

	puntosPorMinuto := *dictionary.NewDictionary[float64, *ej.Ejercicio]()
	var usables []*ej.Ejercicio
	for _, eje := range disponibles {
		puntos, _ := eje.Tipo.Get(tipo)
		puntosPorMinuto.Put(float64(puntos)/float64(eje.Tiempo), eje)
	}
	var aux float64
	pt := puntosPorMinuto.Keys()

	for i := 1; i < puntosPorMinuto.Size(); i++ {
		for j := 0; j < puntosPorMinuto.Size()-1; j++ {
			if pt[j] < pt[j+1] {
				aux = pt[j]
				pt[j] = pt[j+1]
				pt[j+1] = aux
			}
		}
	}

	for i := 0; i < len(pt); i++ {
		eje, _ := puntosPorMinuto.Get(pt[i])
		if duracion >= eje.Tiempo {
			usables = append(usables, eje)
			duracion = duracion - eje.Tiempo
		}
	}
	return crearRutina(usables, nombre)
}

func CrearRutinaMinimaDuracion(nombre string, caloriasAQuemar int, lista *dictionary.Dictionary[string, *ej.Ejercicio]) *Rutina {
	caloriasPorMinuto := *dictionary.NewDictionary[float64, *ej.Ejercicio]()
	var usables []*ej.Ejercicio
	ej := lista.Values()
	for _, eje := range ej {
		caloriasPorMinuto.Put(float64(eje.Calorias)/float64(eje.Tiempo), eje)
	}
	var aux float64
	cal := caloriasPorMinuto.Keys()

	for i := 1; i < caloriasPorMinuto.Size(); i++ {
		for j := 0; j < caloriasPorMinuto.Size()-1; j++ {
			if cal[j] < cal[j+1] {
				aux = cal[j]
				cal[j] = cal[j+1]
				cal[j+1] = aux
			}
		}
	}

	for i := 0; i < len(cal); i++ {
		eje, _ := caloriasPorMinuto.Get(cal[i])
		if caloriasAQuemar >= eje.Calorias {
			usables = append(usables, eje)
			caloriasAQuemar = caloriasAQuemar - eje.Calorias
		}
	}
	return crearRutina(usables, nombre)
}

func CrearRutinaMaximaCantidad(nombre, dificultad string, duracion int, lista *dictionary.Dictionary[string, *ej.Ejercicio], tipo string) *Rutina {
	var disponibles []*ej.Ejercicio
	var necesarios []*ej.Ejercicio
	for _, eje := range lista.Values() {
		if eje.Dificultad == dificultad /*&& eje.Tipo.Contains(tipo)*/ {
			disponibles = append(disponibles, eje)
		}
	}
	//ordenamiento por burbujeo
	var aux *ej.Ejercicio
	for i := 1; i < len(disponibles); i++ {
		for j := 0; j < len(disponibles)-1; j++ {
			if disponibles[j].Tiempo > disponibles[j+1].Tiempo {
				aux = disponibles[j]
				disponibles[j] = disponibles[j+1]
				disponibles[j+1] = aux
			}
		}
	}
	//--------------------------------------
	for _, eje := range disponibles {
		if duracion >= eje.Tiempo {
			necesarios = append(necesarios, eje)
			duracion = duracion - eje.Tiempo
		}
	}
	rutina := crearRutina(necesarios, nombre)
	Rutinas = append(Rutinas, rutina)
	return rutina
}

func EscribirRuCSV(ejercicios []*ej.Ejercicio, rutina *Rutina) {
	// file, _ := os.Create("rutinas.csv")
	var file *os.File
	//Revisa si existe el archivo
	_, err := os.Stat("rutinas.csv")

	// si no existe, lo crea
	if err != nil {
		file, _ = os.Create("rutinas.csv")
		//fmt.Print("creando archivo...")
	} else {
		file, _ = os.OpenFile("rutinas.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		//fmt.Print("abriendo archivo...")
	}

	fmt.Print("\n")
	// inicializa writer
	writer := csv.NewWriter(file)
	var toString []string

	//append a los datos
	//toString = append(toString, "Nombre de rutina: "+rutina.Nombre, "Dificultad: "+rutina.Dificultad, "Calorias quemadas: "+strconv.Itoa(rutina.Calorias), "Duracion total: "+strconv.Itoa(rutina.Duracion))

	toString = append(toString, rutina.Nombre, rutina.Dificultad, strconv.Itoa(rutina.Calorias), strconv.Itoa(rutina.Duracion))

	for _, eje := range ejercicios {
		toString = append(toString, eje.Nombre)
	}

	//escribe
	writer.Write(toString)

	// for num, eje := range ejercicios {
	// 	toString = []string{}
	// 	toString = append(toString, fmt.Sprint(num+1)+"-"+eje.Nombre)
	// 	writer.Write(toString)
	// }

	writer.Flush()
}
