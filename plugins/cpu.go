package plugins

import (
	"errors"
	"fmt"

	"github.com/raffs/go-labs/sistema"
)

type Cpu struct {
	So string
}

func (c *Cpu) Iniciar() {
	fmt.Println("Inicializando o plugin de CPU", c.So)
}

func (c *Cpu) Coletar() (err error) {
	usoCpu := sistema.CpuUso()
	if usoCpu < 0 {
		err = errors.New("Falha ao coletar a cpu")
		return err
	}

	fmt.Printf("CPU: %f (%%)\n", usoCpu)
	return nil
}

func (c *Cpu) Describe() {
	fmt.Println("CPU - Plugin usado para monitorar o uso da CPU")
}
