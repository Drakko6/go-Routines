package procesos

import(
	"fmt"
	"time"
)

type Procesos struct {
	Procesos []*Proceso
}

func (p *Procesos) Mostrar(imprimiendo *int) {

	for _, p := range p.Procesos{
		
		go p.Mostrar(imprimiendo)
		if *imprimiendo == 0 {
			break
		}
	}
}

func (p* Procesos) Buscar (ID uint64) int{ 

	var indice int
	for i:=0; i<len(p.Procesos); i++{
		if p.Procesos[i].ID == ID {
			indice = i
			return indice
		}
	}
	return -1

}

func (p *Procesos)Borrar( i int) []*Proceso {
    return append(p.Procesos[:i], p.Procesos[i+1:]...)
}


type Proceso struct {
	ID uint64 
	i uint64 

}


func (p *Proceso) Start() {
	 
		for {
			p.i = p.i + 1
			time.Sleep(time.Millisecond * 500)
		}
	
}


func (p * Proceso) Mostrar(imprimiendo *int) {

		for{
			fmt.Printf("id %d: %d", p.ID, p.i)
			fmt.Println("")
			time.Sleep(time.Millisecond * 500)
			if *imprimiendo == 0 {
				break
			}
		}
	
}



