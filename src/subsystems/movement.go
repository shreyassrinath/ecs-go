package subsystems

import (
	"ecs"
	//"log"
	"math/rand"
)

type Movement struct {
	Em *ecs.EntityManager
}

func (m Movement) update() {
	//log.Println("In movement update")

	for {
		m.Em.Delay(2)

		for _, val := range m.Em.GetAllEntities() {
			//log.Println("In Movement update For")

			isPosPresent, position := m.Em.GetComponent(val, ecs.COMPONENT_POSITION)
			isDispPresent, displacement := m.Em.GetComponent(val, ecs.COMPONENT_DISPLACEMENT)
			if isPosPresent && isDispPresent {

				pos := position.Data.(*ecs.Position)
				dis := displacement.Data.(*ecs.Displacement)

				dis.X = rand.Intn(3)
				dis.Y = rand.Intn(2)
				pos.X += dis.X
				pos.Y += dis.Y

			}

		}
	}

}
