package factory

import (
	"fmt"
)

func ExecFactoryMethod() {
	pancakeVendor := NewPancakeVendor(NewCornPancakeCook())
	fmt.Printf("Corn pancake value is %v\n", pancakeVendor.SellPancake())

	pancakeVendor = NewPancakeVendor(NewMilletPancakeCook())
	fmt.Printf("Millet pancake value is %v\n", pancakeVendor.SellPancake())
}
