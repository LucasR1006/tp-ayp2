package ejercicio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lista []*Ejercicio

// Definición de la estructura de Ejercicio
type Ejercicio struct {
	Nombre      string
	Descripcion string
	Tipo        []string
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
	fmt.Print("Dificultad? (Facil, Medio, Dificil):")
	dificultad := readLine()

	return &Ejercicio{
		Nombre:      nombre,
		Descripcion: descipcion,
		Tiempo:      tiempo,
		Calorias:    calorias,
		Dificultad:  dificultad,
	}
}

func NuevoEjercicio(nombre, descripcion string, dificultad string, tipo string, tiempo, calorias int, puntos int) *Ejercicio {

	return &Ejercicio{
		Nombre:      nombre,
		Descripcion: descripcion,
		Tiempo:      tiempo,
		Calorias:    calorias,
		Tipo:        []string{tipo},
		Puntos:      puntos,
		Dificultad:  dificultad,
	}
}

func ListaPredefinida() []*Ejercicio {
	s := NuevoEjercicio("sentadillas", "sube y baja con el abdomen", "facil", "flexibilidad", 40, 100, 30)
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
	s := NuevoEjercicio("sentadillas", "100 sentadillas", "dificil", "flexibilidad", 150, 300, 100)
	c := NuevoEjercicio("corre", "Corre por 10 kilometros", "dificil", "cardio", 300, 500, 100)

	rutina := []*Ejercicio{f, a, s, c}
	return rutina
}

func Añadir(eje *Ejercicio) {
	lista = append(lista, eje)
}

func EliminarEjercico(nombre string) {
	for i, eje := range Listado() {
		if eje.Nombre == nombre {
			lista = append(lista[:i], lista[i+1:]...)
			fmt.Println("ejercio eliminad exitosamente.")
			return
		}
	}
	fmt.Println("No se encontró ninguna rutina con el nombre especificado.")
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func Listado() []*Ejercicio {
	lista = ListaPredefinida()
	return lista
}
