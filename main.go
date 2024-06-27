package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	ej "github.com/tp/ejercicio"
	ru "github.com/tp/rutina"
)

// Función principal
func main() {

	clear()
	for {
		fmt.Println("=========================================================")
		fmt.Println("         MENÚ DE EJERCICIOS Y RUTINAS                    ")
		fmt.Println("=========================================================")
		fmt.Println("     1. Generar rutina basica")
		fmt.Println("     2. Generar rutina espartana")
		fmt.Println("     3. Generar rutina dinamica")
		fmt.Println("     4. Opciones de rutinas")
		fmt.Println("     5. Opciones de ejercicios")
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
			fmt.Println("Generando rutina predefinida...")
			rutinaPredefinida := ru.CrearRutinaPredefinida()
			fmt.Println(" ")
			fmt.Println("Rutina creada:")
			ru.MostrarRutina(rutinaPredefinida)
			fmt.Println("=====================================")
			break
		case 2:
			fmt.Println("=====================================")
			fmt.Println("Generando rutina predefinida...")
			rutinaPredefinida := ru.CrearRutinaEspartana()
			fmt.Println(" ")
			fmt.Println("Rutina creada:")
			ru.MostrarRutina(rutinaPredefinida)
			fmt.Println("=====================================")
			break
		case 3:
			fmt.Println("=====================================")
			fmt.Println("Generando rutina dinámica...")
			rutinaDinamica := ru.CrearRutinaDinamica()
			fmt.Println("Rutina creada:")
			ru.MostrarRutina(rutinaDinamica)
			fmt.Println("=====================================")
			break
		case 4:
			opcionesRutinas()
			break
		case 5:
			opcionesEjercicios()
			break
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

func opcionesEjercicios() {
	for {
		fmt.Println("=========================================================")
		fmt.Println("         OPCIONES DE EJERCICIOS        ")
		fmt.Println("=========================================================")
		fmt.Println("     1. Eliminar ejercicio")
		fmt.Println("     2. Consultar ejercicio")
		fmt.Println("     3. Mostrar todas los ejercicio")
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
			ej.EliminarEjercicio(nombre)
			fmt.Println("==================================================")
			break
		case 2:
			fmt.Println("==================================================")
			fmt.Print("Ingrese el nombre del ejercicio que desea consultar: ")
			var nombre string
			fmt.Scanln(&nombre)
			ejercicio := ej.ConsultarEjercicio(nombre)
			if ejercicio != nil {
				ej.MostrarEjercicio(ejercicio)
			}
			fmt.Println("==================================================")
			break
		case 3:
			fmt.Println("==================================================")
			fmt.Println("Listado de todos los ejercicios:")
			for _, ejercicios := range lista {
				fmt.Println("-", ejercicios.Nombre)
			}
			fmt.Println("==================================================")
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
