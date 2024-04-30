package ejercicios

func DarDeAlta(ejercicio *Ejercicio) {
	ejercicios[ejercicio.Nombre] = ejercicio
}

func BajarEjercicio(nombre string) {
	delete(ejercicios, nombre)
}

func ModificarEjercicio(nombre string, ejercicio *Ejercicio) {
	ejercicios[nombre] = ejercicio
}

func ConsultarEjercicio(nombre string) *Ejercicio {
	return ejercicios[nombre]
}

func ListarEjercicios() []*Ejercicio {
	var lista []*Ejercicio
	for _, ejercicio := range ejercicios {
		lista = append(lista, ejercicio)
	}
	return lista
}
