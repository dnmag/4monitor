package daemon

import (
	"fmt"
)

func (d *Daemon) Versao() {
	fmt.Printf("Monitor version %s \n", d.versao)
}
