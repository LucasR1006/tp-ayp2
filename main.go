package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	ej "github.com/tp/ejercicio"
	ru "github.com/tp/rutina"
	"github.com/untref-ayp2/data-structures/dictionary"
)

// Función principal
func main() {

	clear()

	listaActual, eje := ej.ListaPredefinida()
	listaActual.EscribirEjCSV()

	for {
		fmt.Println("=========================================================")
		fmt.Println("         MENÚ DE EJERCICIOS Y RUTINAS                    ")
		fmt.Println("=========================================================")
		fmt.Println("     1. Generar rutina basica")
		fmt.Println("     2. Generar rutina dificil")
		fmt.Println("     3. Generar rutina dinamica")
		fmt.Println("     4. Opciones de rutinas")
		fmt.Println("     5. Opciones de ejercicios")
		fmt.Println("     6. Borrar archivos .csv")
		fmt.Println("     0. Salir")
		fmt.Println("=========================================================")
		fmt.Print("Ingrese su opción: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 0:
			fmt.Println("=====================================")
			fmt.Println("¡Hasta luego!")
			fmt.Println("=====================================")
			return
		case 1:
			fmt.Println("=====================================")
			fmt.Println("Generando rutina basica...")
			rutinaPredefinida := ru.CrearRutinaBasica()
			fmt.Println(" ")
			fmt.Println("Rutina creada:")
			ru.MostrarRutina(rutinaPredefinida)
			ru.EscribirRuCSV(eje, rutinaPredefinida)
			fmt.Println("=====================================")
			break
		case 2:
			fmt.Println("=====================================")
			fmt.Println("Generando rutina Dificil...")
			rutinaPredefinida := ru.CrearRutinaDificil()
			fmt.Println(" ")
			fmt.Println("Rutina creada:")
			ru.MostrarRutina(rutinaPredefinida)
			ru.EscribirRuCSV(eje, rutinaPredefinida)
			fmt.Println("=====================================")
			break
		case 3:
			OpcionesAutoMagicas(listaActual.Listado())
			break
		case 4:
			opcionesRutinas()
			break
		case 5:
			opcionesEjercicios(listaActual)
			break
		case 6:
			os.Create("ejercicios.csv")
			os.Create("rutinas.csv")
			fmt.Println("Archivos eliminados correctamente")
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

func opcionesRutinas() {
	for {
		fmt.Println("=========================================================")
		fmt.Println("         OPCIONES DE RUTINAS        ")
		fmt.Println("=========================================================")
		fmt.Println("     1. Eliminar rutina")
		fmt.Println("     2. Consultar rutina")
		fmt.Println("     3. Mostrar todas las rutinas")
		fmt.Println("     0. Volver atras")
		fmt.Println("=========================================================")
		fmt.Print("Ingrese una opción: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 0:
			return
		case 1:
			fmt.Println("==================================================")
			fmt.Print("Ingrese el nombre de la rutina que desea eliminar: ")
			var nombre string
			fmt.Scanln(&nombre)
			ru.EliminarRutina(nombre)
			fmt.Println("==================================================")
			break
		case 2:
			fmt.Println("==================================================")
			fmt.Print("Ingrese el nombre de la rutina que desea consultar: ")
			var nombre string
			fmt.Scanln(&nombre)
			rutina := ru.ConsultarRutina(nombre)
			if rutina != nil {
				ru.MostrarRutina(rutina)
			}
			fmt.Println("==================================================")
			break
		case 3:
			fmt.Println("==================================================")
			fmt.Println("Listado de todas las rutinas:")
			for _, rutina := range ru.Rutinas {
				fmt.Println("-", rutina.Nombre)
			}
			fmt.Println("==================================================")
			break
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

func opcionesEjercicios(lista *ej.Lista) {
	for {
		fmt.Println("=========================================================")
		fmt.Println("         OPCIONES DE EJERCICIOS        ")
		fmt.Println("=========================================================")
		fmt.Println("     1. Eliminar ejercicio")
		fmt.Println("     2. Consultar ejercicio")
		fmt.Println("     3. Mostrar todos los ejercicios")
		fmt.Println("     4. Añadir ejercicio")
		fmt.Println("     0. Volver atras")
		fmt.Println("=========================================================")
		fmt.Print("Ingrese una opción: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 0:
			return
		case 1:
			fmt.Println("==================================================")
			fmt.Print("Ingrese el nombre del ejercicio que desea eliminar: ")
			var nombre string
			fmt.Scanln(&nombre)
			lista.EliminarEjercicio(nombre)
			os.Create("ejercicios.csv")
			lista.EscribirEjCSV()
			fmt.Println("==================================================")
			break
		case 2:
			fmt.Println("==================================================")
			fmt.Print("Ingrese el nombre del ejercicio que desea consultar: ")
			var nombre string
			fmt.Scanln(&nombre)
			fmt.Println("En el archivo ejercicios.csv puede buscar la informacion de " + nombre)
			fmt.Println("==================================================")
			break
		case 3:
			fmt.Println("==================================================")
			fmt.Println("Listado de todos los ejercicios:")
			fmt.Println("-", lista.Listado())
			fmt.Println("==================================================")
			break
		case 4:
			fmt.Println("==================================================")
			eje := ej.CrearEjercicio()
			lista.Añadir(eje)
			lista.EscribirEjCSV()
			fmt.Println("==================================================")
			break
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

func OpcionesAutoMagicas(lista *dictionary.Dictionary[string, *ej.Ejercicio]) {
	for {
		fmt.Println("=========================================================")
		fmt.Println("         OPCIONES DE RUTINAS AUTOMAGICAS        ")
		fmt.Println("=========================================================")
		fmt.Println("     1. Crear Rutina Maxima Cantidad")
		fmt.Println("     2. crear Rutina Minima Duracion")
		fmt.Println("     3. Crear Rutina Maximo Puntaje")
		fmt.Println("     4. Crear Rutina Manual")
		fmt.Println("     0. Volver atras")
		fmt.Println("=========================================================")
		fmt.Print("Ingrese una opción: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 0:
			return
		case 1:
			fmt.Println("==================================================")
			fmt.Print("Ingrese el nombre de la rutina: ")
			var nombre string
			fmt.Scanln(&nombre)
			fmt.Print("Ingrese la dificultad que se busca en los ejercicios: ")
			var dificultad string
			fmt.Scanln(&dificultad)
			fmt.Print("Ingrese la duracion en minutos de la rutina: ")
			var duracion int
			fmt.Scanln(&duracion)
			fmt.Print("Ingrese el tipo de ejercicios: ")
			var tipo string
			fmt.Scanln(&tipo)
			rutinaMaxCantidad := ru.CrearRutinaMaximaCantidad(nombre, dificultad, duracion, lista, tipo)
			ru.EscribirRuCSV(rutinaMaxCantidad.Ejercicios, rutinaMaxCantidad)
			ru.MostrarRutina(rutinaMaxCantidad)
			fmt.Println("==================================================")
			break
		case 2:
			fmt.Println("==================================================")
			fmt.Print("Ingrese el nombre de la rutina: ")
			var nombre string
			fmt.Scanln(&nombre)
			fmt.Print("Ingrese las calorias a quemar en la rutina: ")
			var calorias int
			fmt.Scanln(&calorias)
			rutinaMinDuracion := ru.CrearRutinaMinimaDuracion(nombre, calorias, lista)
			ru.MostrarRutina(rutinaMinDuracion)
			ru.EscribirRuCSV(rutinaMinDuracion.Ejercicios, rutinaMinDuracion)
			fmt.Println("==================================================")
			break
		case 3:
			fmt.Println("==================================================")
			fmt.Print("Ingrese el nombre de la rutina: ")
			var nombre string
			fmt.Scanln(&nombre)
			fmt.Print("Ingrese la duracion de la rutina: ")
			var duracion int
			fmt.Scanln(&duracion)
			fmt.Print("Ingrese el tipo de ejercicios: ")
			var tipo string
			fmt.Scanln(&tipo)
			rutinaMaxPuntaje := ru.CrearRutinaMaximoPuntaje(nombre, duracion, tipo, lista)
			ru.EscribirRuCSV(rutinaMaxPuntaje.Ejercicios, rutinaMaxPuntaje)
			ru.MostrarRutina(rutinaMaxPuntaje)
			fmt.Println("==================================================")
			break
		case 4:
			fmt.Println("=====================================")
			fmt.Println("Generando rutina manual...")
			rutinaManual := ru.CrearRutinaManual()
			fmt.Println("Rutina creada:")
			ru.EscribirRuCSV(rutinaManual.Ejercicios, rutinaManual)
			ru.MostrarRutina(rutinaManual)
			fmt.Println("=====================================")
			break
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

// ------------------------------------------------------------------------------------------------------------------

// ------------------------------------------------------------------------------------------------------------------
func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}
