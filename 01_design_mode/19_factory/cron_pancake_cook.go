package factory

// cornPancakeCook 制作玉米面煎饼厨师
type cornPancakeCook struct{}

func NewCornPancakeCook() *cornPancakeCook {
	return &cornPancakeCook{}
}

func (cook *cornPancakeCook) MakePancake() Pancake {
	return NewCornPancake()
}

// milletPancakeCook 制作小米面煎饼厨师
type milletPancakeCook struct{}

func NewMilletPancakeCook() *milletPancakeCook {
	return &milletPancakeCook{}
}

func (cook *milletPancakeCook) MakePancake() Pancake {
	return NewMilletPancake()
}
