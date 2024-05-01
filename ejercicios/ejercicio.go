package ejercicios

type Ejercicio struct {
	Nombre         string
	Descripcion    string
	Tipo           []string
	Dificultad     string
	TiempoEstimado int // en segundos
	Calorias       int
	Puntos         int
}

var ejercicios = make(map[string]*Ejercicio)

func NuevoEjercicio(nombre, descripcion string, dificultad string, tipo string, tiempoEstimado, calorias int, puntos int) *Ejercicio {

	return &Ejercicio{
		Nombre:         nombre,
		Descripcion:    descripcion,
		TiempoEstimado: tiempoEstimado,
		Calorias:       calorias,
		Tipo:           []string{tipo},
		Puntos:         puntos,
		Dificultad:     dificultad,
	}
}
