package factory

// Pancake 煎饼
type Pancake interface {
	// ShowFlour 煎饼使用的面粉
	ShowFlour() string
	// Value 煎饼价格
	Value() float32
}

// PancakeCook 煎饼厨师
type PancakeCook interface {
	// MakePancake 摊煎饼
	MakePancake() Pancake
}

// PancakeVendor 煎饼小贩
type PancakeVendor struct {
	PancakeCook
}

// NewPancakeVendor ...
func NewPancakeVendor(cook PancakeCook) *PancakeVendor {
	return &PancakeVendor{
		PancakeCook: cook,
	}
}

// SellPancake 卖煎饼，先摊煎饼，再卖
func (vendor *PancakeVendor) SellPancake() (money float32) {
	return vendor.MakePancake().Value()
}
