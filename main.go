package main

import (
	"fmt"
	"./procesos"
	"os"
	"os/exec"
	
)


func main () {


	procesosSlice := procesos.Procesos{Procesos: []*procesos.Proceso{}}
	id:= uint64(0)


	var op uint = 1

	imprimiendo := 0
	cont :=0

	for op != 0 {
		cmd := exec.Command("cmd", "/c", "cls") 
        cmd.Stdout = os.Stdout
        cmd.Run()
		fmt.Println("")
		fmt.Println("	ADMINISTRADOR DE PROCESOS")
		fmt.Println("1.- Agregar proceso")
		fmt.Println("2.- Mostrar procesos")
		fmt.Println("3.- Terminar proceso")
		fmt.Println("0.- Salir")

		fmt.Scan(&op)

		switch {
		case op == 1:
			
			p := procesos.Proceso{ID: id}
			procesosSlice.Procesos = append(procesosSlice.Procesos, &p)
			go p.Start(0)
			id++

		case op == 2:

			if(imprimiendo == 0){
				cont = 0
			}
			imprimiendo = 1

			if cont != 0  {
				imprimiendo = 0 
			}

			if(imprimiendo==1){
			procesosSlice.Mostrar(&imprimiendo)
			}

			cont++


		case op == 3:

			var numProceso uint64

			fmt.Println("ID de proceso a terminar: ")
			fmt.Scan(&numProceso)

			
			i:= procesosSlice.Buscar(numProceso)
			procesosSlice.Procesos = procesosSlice.Borrar(i)
			
		
		case op == 0:
			fmt.Println("Saliendo...")

		default:
			fmt.Println("Opción inválida")
		}

	}


}