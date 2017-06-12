package subsystems

import (
	"ecs"
	"log"
)

type Render struct {
	Em *ecs.EntityManager
}

func (r Render) update() {

	for {
		r.Em.Delay(2)

		for _, val := range r.Em.GetAllEntities() {
			isPosPresent, position := r.Em.GetComponent(val, ecs.COMPONENT_POSITION)
			isDispPresent, displayName := r.Em.GetComponent(val, ecs.COMPONENT_DISPLAYNAME)
			if isPosPresent && isDispPresent {

				pos := position.Data.(*ecs.Position)
				dis := displayName.Data.(*ecs.DisplayName)

				log.Println(dis.Name, " is at X:", pos.X, " Y:", pos.Y)

			}

		}
		log.Println("\n")
	}

}
