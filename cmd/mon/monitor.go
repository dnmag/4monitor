package main

import (
	"fmt"
	"monitor/api"
	"monitor/daemon"
	"monitor/plugins"
	"os"
)

func usage() {
	fmt.Println(`Usage: monitor <command>

Comandos:

	versao		- Mostra a versão do daemon
	start		- Inicializa o daemon
	plugin		- Descreve cada plugin usado no daemon
`)
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	command := os.Args[1]
	daemon := daemon.New("v0.0.1-alpha")

	// Registrando todos os plugins que serão codificados no pacote plugins
	daemon.Plugins = []api.Plugin{
		&plugins.Cpu{So: "linux"},
		&plugins.Memoria{So: "linux"},
	}

	switch command {
	case "versao":
		daemon.Versao()
	case "start":
		daemon.Start()
	case "plugin":
		for _, plugin := range daemon.Plugins {
			plugin.Describe()
		}
	}
}
