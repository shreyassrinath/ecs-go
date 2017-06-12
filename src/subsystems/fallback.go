package subsystems

import (
	"../ecs"
	"log"
	"math/rand"
	"strconv"
)

type FallBack struct {
	Em *ecs.EntityManager
}

func (f FallBack) update() {

	for {
		f.Em.Delay(2)

		for _, val := range f.Em.GetAllEntities() {

			isPosPresent, position := f.Em.GetComponent(val, ecs.COMPONENT_POSITION)
			isDispPresent, _ := f.Em.GetComponent(val, ecs.COMPONENT_DISPLACEMENT)
			if isPosPresent && isDispPresent {

				pos := position.Data.(*ecs.Position)
				if pos.X > 20 || pos.Y > 20 {
					log.Println("Falling Back! \n")
					pos.X = rand.Intn(10)
					pos.Y = rand.Intn(15)
					log.Println("Adding random tree!! \n")
					e2 := f.Em.CreateEntity()
					f.Em.AddComponent(e2, ecs.NewComponent(&ecs.DisplayName{Name: "Tree" + strconv.Itoa(pos.X) + strconv.Itoa(pos.Y)}, ecs.COMPONENT_DISPLAYNAME))
					f.Em.AddComponent(e2, ecs.NewComponent(&ecs.Position{X: pos.X, Y: pos.Y}, ecs.COMPONENT_POSITION))

				}

			}

		}

	}

}
