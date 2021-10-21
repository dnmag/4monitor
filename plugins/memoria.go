package plugins

import (
	"errors"
	"fmt"

	"github.com/raffs/go-labs/sistema"
)

type Memoria struct {
	So string
}

func (c *Memoria) Iniciar() {
	fmt.Println("Inicializando o plugin de Memoria", c.So)
}

func (c *Memoria) Coletar() (err error) {
	usoMem := sistema.MemUso()
	if usoMem < 0 {
		err = errors.New("Falha ao coletar a memoria")
		return err
	}

	fmt.Printf("MEM: %f (%%)\n", usoMem)
	return nil
}

func (c *Memoria) Describe() {
	fmt.Println("Memoria - Plugin usado para monitorar o uso da Memoria")
}
