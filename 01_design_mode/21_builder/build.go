package builder

type normalPancakeBuilder struct {
	pasteQuantity Quantity // 面糊量
	eggNum        int      // 鸡蛋数量
	friedWafer    string   // 油炸薄脆
	hasCoriander  bool     // 是否放香菜
	hasShallot    bool     // 是否放葱
	hasHotSauce   bool     // 是否放辣味酱
}

func NewNormalPancakeBuilder() *normalPancakeBuilder {
	return &normalPancakeBuilder{}
}

func (n *normalPancakeBuilder) PutPaste(quantity Quantity) {
	n.pasteQuantity = quantity
}

func (n *normalPancakeBuilder) PutEgg(num int) {
	n.eggNum = num
}

func (n *normalPancakeBuilder) PutWafer() {
	n.friedWafer = "油炸的薄脆"
}

func (n *normalPancakeBuilder) PutFlavour(hasCoriander, hasShallot, hasSauce bool) {
	n.hasCoriander = hasCoriander
	n.hasShallot = hasShallot
	n.hasHotSauce = hasSauce
}

func (n *normalPancakeBuilder) Build() *Pancake {
	return &Pancake{
		pasteQuantity: n.pasteQuantity,
		eggNum:        n.eggNum,
		wafer:         n.friedWafer,
		hasCoriander:  n.hasCoriander,
		hasShallot:    n.hasShallot,
		hasSauce:      n.hasHotSauce,
	}
}
