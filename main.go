package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	ru "github.com/tp/rutina"
)

// Función principal
func main() {

	clear()
	for {
		fmt.Println("=========================================================")
		fmt.Println("         MENÚ DE GESTIÓN DE EJERCICIOS Y RUTINAS        ")
		fmt.Println("=========================================================")
		fmt.Println("     1. Generar rutina predefinida")
		fmt.Println("     2. Generar rutina dinamica")
		fmt.Println("     3. Opciones de rutinas")
		fmt.Println("     0. Salir")
		fmt.Println("=========================================================")
		fmt.Print("Ingrese una opción: ")

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
			fmt.Println("Generando rutina dinámica...")
			rutinaDinamica := ru.CrearRutinaDinamica()
			fmt.Println("Rutina creada:")
			ru.MostrarRutina(rutinaDinamica)
			fmt.Println("=====================================")
			break
		case 3:
			otrasOpciones()
			break
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

func otrasOpciones() {
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
