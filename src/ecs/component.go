package ecs

const (
	COMPONENT_NONE         = 0
	COMPONENT_RECT         = 1 << 1
	COMPONENT_CIRCLE       = 1 << 2
	COMPONENT_POSITION     = 1 << 3
	COMPONENT_DISPLACEMENT = 1 << 4
	COMPONENT_DISPLAYNAME  = 1 << 5
)

type Component struct {
	ComponentType int
	Data          interface{}
}

type Rect struct {
	Width, Height float64
}

type DisplayName struct {
	Name string
}

type Position struct {
	X, Y int
}

type Displacement struct {
	X, Y int
}

type Circle struct {
	Radius float64
}

//type Test int

func NewComponent(data interface{}, componentType int) *Component {
	return &Component{componentType, data}
}
