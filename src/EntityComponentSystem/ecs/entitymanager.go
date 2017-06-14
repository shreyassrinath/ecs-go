package ecs

import (
	"time"
)

const MaxInt = int(^uint(0) >> 1)

type EntityManager struct {
	lowestUnassignedInt int
	allEntities         []int
	componentStores     map[int]map[int]Component //TODO: Need to be concurrent: https://github.com/streamrail/concurrent-map
	//OR need to Implement via channels(sharing by communicating)
}

func NewEntityManager() *EntityManager {
	return &EntityManager{1, make([]int, 0), make(map[int]map[int]Component)}

}

func (e *EntityManager) GetComponent(entity int, componentType int) (bool, Component) {
	//return nil
	val, _ := e.componentStores[componentType]
	//TODO: handle non presence
	value, isPresent := val[entity]
	/*if !comp {
		return nil
	}*/
	//log.Println("Value is ", value)

	return isPresent, value

}

func (e *EntityManager) GetAllComponentsOfType(componentType int) []Component {
	//return nil
	val, _ := e.componentStores[componentType]
	//TODO: handle non presence
	allComponents := make([]Component, 0)
	for _, value := range val {
		allComponents = append(allComponents, value)

	}
	return allComponents

}

func (e *EntityManager) GetAllEntitiesOfComponentType(componentType int) []int {
	//return nil
	val, _ := e.componentStores[componentType]
	//TODO: handle non presence
	allEntities := make([]int, 0)
	for key, _ := range val {
		allEntities = append(allEntities, key)

	}
	return allEntities

}

func (e *EntityManager) AddComponent(entity int, component *Component) {

	_, isPresent := e.componentStores[component.ComponentType]
	if isPresent {
		e.componentStores[component.ComponentType][entity] = *component

	} else {
		e.componentStores[component.ComponentType] = map[int]Component{entity: *component}

	}
}

func (e *EntityManager) CreateEntity() int {
	newEntityId := e.generateEntityId()
	e.allEntities = append(e.allEntities, newEntityId)
	return newEntityId

}

func (e *EntityManager) DeleteEntity(entity int) {

	e.allEntities = append(e.allEntities[:entity], e.allEntities[entity+1:]...)

}

func (e *EntityManager) GetAllEntities() []int {
	return e.allEntities

}

func (e *EntityManager) Delay(delay int) {
	timer := time.NewTimer(time.Second * time.Duration(2))
	<-timer.C

}

func (e *EntityManager) generateEntityId() int {

	if e.lowestUnassignedInt < MaxInt {
		e.lowestUnassignedInt = e.lowestUnassignedInt + 1
		return e.lowestUnassignedInt

	} else { //TODO: Handle assignment
		return 0
	}

}
