package subsystems

import (
	"EntityComponentSystem/ecs"
	"log"
)

type Cleanup struct {
	Em *ecs.EntityManager
}

func (c Cleanup) update() {

	for {
		c.Em.Delay(2)

		lenOfEntities := len(c.Em.GetAllEntities())

		if lenOfEntities > 7 {
			log.Println("We need to clean up growing entities!! \n")
			c.Em.DeleteEntity(lenOfEntities - 1)

		}

	}

}
