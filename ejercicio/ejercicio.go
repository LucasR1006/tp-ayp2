package ejercicio

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/untref-ayp2/data-structures/dictionary"
)

type Lista struct {
	lista *dictionary.Dictionary[string, *Ejercicio]
}

// Definición de la estructura de Ejercicio
type Ejercicio struct {
	Nombre      string
	Descripcion string
	Tipo        *dictionary.Dictionary[string, int]
	Dificultad  string
	Tiempo      int // en segundos
	Calorias    int
	Puntos      int
}

//Funciòn para crear un ejercicio nuevo por el usuario

func CrearEjercicio() *Ejercicio {
	fmt.Println("Creando Ejercicio...")
	fmt.Print("Que Nombre deseas ponerle?:")
	nombre := readLine()
	fmt.Print("Algun detalle para recordarlo?:")
	descipcion := readLine()
	fmt.Print("Cuanto tiempo le vas a poner? (En segundos):")
	var tiempo int
	fmt.Scanln(&tiempo)
	fmt.Print("Gasto estimado en calorias?:")
	var calorias int
	fmt.Scanln(&calorias)
	//fmt.Print("Que tipo de ejercicio es?(cardio, fuerza o flexibilidad):")

	//tipo := readLine()
	fmt.Print("Dificultad? (Facil, Medio, Dificil):")
	dificultad := readLine()

	return &Ejercicio{
		Nombre:      nombre,
		Descripcion: descipcion,
		Tiempo:      tiempo,
		Calorias:    calorias,
		//Tipo:        Tipo,
		Dificultad: dificultad,
	}
}

func NuevoEjercicio(nombre, descripcion string, dificultad string, tipo string, tiempo, calorias int, puntos int) *Ejercicio {
	Tipo := dictionary.NewDictionary[string, int]()
	return &Ejercicio{
		Nombre:      nombre,
		Descripcion: descripcion,
		Tiempo:      tiempo,
		Calorias:    calorias,
		Tipo:        Tipo,
		Puntos:      puntos,
		Dificultad:  dificultad,
	}
}

func ListaPredefinida() []*Ejercicio {
	s := NuevoEjercicio("sentadillas", "flexión de rodillas hasta formar un ángulo de 90 grados",
		"medio", "flexibilidad y fuerza", 40, 100, 30)
	f := NuevoEjercicio("flexiones", "levantando el cuerpo únicamente con los brazos y bajando de nuevo al suelo",
		"medio", "fuerza", 25, 250, 50)
	j := NuevoEjercicio("jumping jack", "saltar mientras abres y cierras piernas y brazos",
		"dificil", "cardio", 30, 560, 60)
	z := NuevoEjercicio("zancadas alternas", "dar una zancada atrás y bajar la rodilla hasta que la pierna quede recta",
		"medio", "flexibilidad", 12, 500, 20)
	p := NuevoEjercicio("levantamiento de pesas", "agarrar pesas y subirlas y bajarlas con los brazos",
		"medio", "fuerza", 30, 200, 20)
	rutina := []*Ejercicio{s, f, j, z, p}
	return rutina
}

func EntrenamientoEspartano() []*Ejercicio {
	f := NuevoEjercicio("flexiones", "100 flexiones", "dificil", "fuerza", 1000, 500, 100)
	a := NuevoEjercicio("abdominales", "100 abs", "dificil", "fuerza", 1000, 500, 100)
	s := NuevoEjercicio("sentadillas", "100 sentadillas", "dificil", "flexibilidad y fuerza", 150, 300, 100)
	c := NuevoEjercicio("corre", "Corre por 10 kilometros", "dificil", "cardio", 300, 500, 100)

	rutina := []*Ejercicio{f, a, s, c}
	return rutina
}

func (l *Lista) Añadir(eje ...*Ejercicio) {
	for _, eje := range eje {
		l.lista.Put(eje.Nombre, eje)
	}
}

func (l *Lista) EliminarEjercicio(nombre string) {
	for _, eje := range l.lista.Values() {
		if eje.Nombre == nombre {
			l.lista.Remove(nombre)
			fmt.Println("Ejercio eliminado.")
			return
		}
	}
	fmt.Println("No se encontró ningun ejercicio con el nombre especificado.")
}

func MostrarEjercicio(ejercicio *Ejercicio) {
	fmt.Println("Nombre del ejercicio:", ejercicio.Nombre)
	fmt.Println("Descripcion del ejercicio:", ejercicio.Descripcion)
	fmt.Println("Tipo:", ejercicio.Tipo)
	fmt.Println("Duración Total:", ejercicio.Tiempo, "segundos")
	fmt.Println("Calorías Quemadas:", ejercicio.Calorias)
	fmt.Println("Dificultad:", ejercicio.Dificultad)
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func (l *Lista) Listado() *dictionary.Dictionary[string, *Ejercicio] {
	return l.lista
}
func (l *Lista) EscribirEjCSV() {

	file, _ := os.Create("ejercicios.csv")
	writer := csv.NewWriter(file)
	var toString []string

	for _, eje := range l.lista.Values() {
		toString = []string{}
		toString = append(toString, eje.Nombre, eje.Descripcion, fmt.Sprint(eje.Tipo.Keys()), eje.Dificultad, fmt.Sprint(eje.Tiempo), fmt.Sprint(eje.Calorias), fmt.Sprint(eje.Tipo.Values()))
		writer.Write(toString)
	}
	writer.Flush()
}
