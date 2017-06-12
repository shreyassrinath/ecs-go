package ecsinitialize

import (
	"../ecs"
	"../subsystems"
)

func InitializeECS() *ecs.EntityManager {

	em := ecs.NewEntityManager()

	//TODO: Need to create template

	e1 := em.CreateEntity()
	em.AddComponent(e1, ecs.NewComponent(&ecs.DisplayName{Name: "Player"}, ecs.COMPONENT_DISPLAYNAME))
	em.AddComponent(e1, ecs.NewComponent(&ecs.Position{X: 3, Y: 4}, ecs.COMPONENT_POSITION))
	em.AddComponent(e1, ecs.NewComponent(&ecs.Displacement{X: 0, Y: 0}, ecs.COMPONENT_DISPLACEMENT))

	e2 := em.CreateEntity()
	em.AddComponent(e2, ecs.NewComponent(&ecs.DisplayName{Name: "Tree"}, ecs.COMPONENT_DISPLAYNAME))
	em.AddComponent(e2, ecs.NewComponent(&ecs.Position{X: 2, Y: 5}, ecs.COMPONENT_POSITION))

	e3 := em.CreateEntity()
	em.AddComponent(e3, ecs.NewComponent(&ecs.DisplayName{Name: "AI"}, ecs.COMPONENT_DISPLAYNAME))
	em.AddComponent(e3, ecs.NewComponent(&ecs.Position{X: 9, Y: 6}, ecs.COMPONENT_POSITION))
	em.AddComponent(e3, ecs.NewComponent(&ecs.Displacement{X: 0, Y: 0}, ecs.COMPONENT_DISPLACEMENT))

	m := subsystems.Movement{Em: em}
	r := subsystems.Render{Em: em}
	f := subsystems.FallBack{Em: em}
	c := subsystems.Cleanup{Em: em}
	//log.Println("Running movement!!!!!!!!")
	subsystems.Run(m)
	subsystems.Run(r)
	subsystems.Run(f)
	subsystems.Run(c)

	return em

}
