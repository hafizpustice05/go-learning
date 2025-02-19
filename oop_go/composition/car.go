package composition

type Engine struct {
	hp   int
	name string
}

func (e Engine) Hp() int {
	return e.hp
}

func (e Engine) Name() string {
	return e.name
}

type Wheel struct {
	wheelDimension int
	name           string
}

func (w Wheel) WheelDimension() int {
	return w.wheelDimension
}

func (w Wheel) Name() string {
	return w.name
}

type Car struct {
	name string
	Engine
	Wheel
}

func NewCar(name string, hp int, wheelDimension int) Car {
	return Car{
		name: name,
		Engine: Engine{
			hp:   hp,
			name: "shimano Engine",
		},
		Wheel: Wheel{
			wheelDimension: wheelDimension,
			name:           "RockRider pro 200",
		},
	}
}

func (c Car) EngineName() string {
	return c.Engine.Name()
}

func (c Car) WheelName() string {
	return c.Wheel.Name()
}
