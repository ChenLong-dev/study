package builder

// PancakeCook 摊煎饼师傅
type PancakeCook struct {
	builder PancakeBuilder
}

func NewPancakeCook(builder PancakeBuilder) *PancakeCook {
	return &PancakeCook{
		builder: builder,
	}
}

// SetPancakeBuilder 重新设置煎饼构造器
func (p *PancakeCook) SetPancakeBuilder(builder PancakeBuilder) {
	p.builder = builder
}

// MakePancake 摊一个一般煎饼
func (p *PancakeCook) MakePancake() *Pancake {
	p.builder.PutPaste(Middle)
	p.builder.PutEgg(1)
	p.builder.PutWafer()
	p.builder.PutFlavour(true, true, true)
	return p.builder.Build()
}

// MakeBigPancake 摊一个巨无霸煎饼
func (p *PancakeCook) MakeBigPancake() *Pancake {
	p.builder.PutPaste(Large)
	p.builder.PutEgg(3)
	p.builder.PutWafer()
	p.builder.PutFlavour(true, true, true)
	return p.builder.Build()
}

// MakePancakeForFlavour 摊一个自选调料霸煎饼
func (p *PancakeCook) MakePancakeForFlavour(hasCoriander, hasShallot, hasSauce bool) *Pancake {
	p.builder.PutPaste(Large)
	p.builder.PutEgg(3)
	p.builder.PutWafer()
	p.builder.PutFlavour(hasCoriander, hasShallot, hasSauce)
	return p.builder.Build()
}
