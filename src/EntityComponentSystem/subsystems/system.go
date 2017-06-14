package subsystems

import (
//"log"
)

type System interface {
	update()
}

func Run(s System) {
	//log.Println("In Sys update")

	go s.update()
}
