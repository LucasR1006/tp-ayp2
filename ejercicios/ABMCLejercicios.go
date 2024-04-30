package ejercicios

func AltaDeEjercicio(ejercicio *Ejercicio) {
	ejercicios[ejercicio.Nombre] = ejercicio
}

func BajaDeEjercicio(nombre string) {
	delete(ejercicios, nombre)
}

func ModificarEjercicio(nombre string, ejercicio *Ejercicio) {
	ejercicios[nombre] = ejercicio
}

func ConsultaDeEjercicio(nombre string) *Ejercicio {
	return ejercicios[nombre]
}

func ListadoDeEjercicios() []*Ejercicio {
	var lista []*Ejercicio
	for _, ejercicio := range ejercicios {
		lista = append(lista, ejercicio)
	}
	return lista
}
