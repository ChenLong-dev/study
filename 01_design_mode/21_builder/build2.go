package builder

type healthyPancakeBuilder struct {
	milletPasteQuantity Quantity // 小米面糊量
	chaiEggNum          int      // 柴鸡蛋数量
	nonFriedWafer       string   // 非油炸薄脆
	hasCoriander        bool     // 是否放香菜
	hasShallot          bool     // 是否放葱
}

func NewHealthyPancakeBuilder() *healthyPancakeBuilder {
	return &healthyPancakeBuilder{}
}

func (n *healthyPancakeBuilder) PutPaste(quantity Quantity) {
	n.milletPasteQuantity = quantity
}

func (n *healthyPancakeBuilder) PutEgg(num int) {
	n.chaiEggNum = num
}

func (n *healthyPancakeBuilder) PutWafer() {
	n.nonFriedWafer = "非油炸的薄脆"
}

func (n *healthyPancakeBuilder) PutFlavour(hasCoriander, hasShallot, _ bool) {
	n.hasCoriander = hasCoriander
	n.hasShallot = hasShallot
}

func (n *healthyPancakeBuilder) Build() *Pancake {
	return &Pancake{
		pasteQuantity: n.milletPasteQuantity,
		eggNum:        n.chaiEggNum,
		wafer:         n.nonFriedWafer,
		hasCoriander:  n.hasCoriander,
		hasShallot:    n.hasShallot,
		hasSauce:      false,
	}
}
