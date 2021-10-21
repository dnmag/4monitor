package daemon

import (
	"monitor/api"
)

type Daemon struct {
	// versao do daemon
	versao string

	// Lista de plugins que o daemon ira utilizar
	Plugins []api.Plugin
}

// Retorna uma nova estrutura daemon inicializada com
// com a versao
func New(versao string) (d Daemon) {
	d = Daemon{
		versao:  versao,
		Plugins: []api.Plugin{},
	}
	return d
}
