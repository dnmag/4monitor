package daemon

import (
	"fmt"
	"time"
)

func (d *Daemon) initPlugins() {
	for _, plugin := range d.Plugins {
		plugin.Iniciar()
	}
}

func (d *Daemon) Start() {
	fmt.Println("Inicializando os plugins")
	d.initPlugins()

	for {
		for _, plugin := range d.Plugins {
			err := plugin.Coletar()
			if err != nil {
				fmt.Println("Erro ao coltar o plugin:", err)
			}
		}
		time.Sleep(1 * time.Second)
	}
}
