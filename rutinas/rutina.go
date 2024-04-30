package rutinas

import ej "github.com/tp/ejercicios"

type Rutina struct {
	Nombre     string
	Ejercicios []*ej.Ejercicio
	Duracion   int // en minutos
	Calorias   int
	Tipo       map[string]int
	Dificultad string
}

var rutinas = make(map[string]*Rutina)

func NuevaRutina(nombre string) *Rutina {
	return &Rutina{
		Nombre: nombre,
	}
}

func AgregarRutina(nombre string, rutina *Rutina) {
	rutinas[nombre] = rutina
}

func EliminarRutina(nombre string) {
	delete(rutinas, nombre)
}

func ModificarRutina(nombre string, rutina *Rutina) {
	rutinas[nombre] = rutina
}

func ConsultarRutina(nombre string) *Rutina {
	return rutinas[nombre]
}

func ListarRutinas() []*Rutina {
	var lista []*Rutina
	for _, rutina := range rutinas {
		lista = append(lista, rutina)
	}
	return lista
}
