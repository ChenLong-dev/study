package strategy

import (
	"fmt"
)

func ExecStrategy() {
	Beijing := NewCity("北京", "四季分明")

	Beijing.SetSeason(NewSpring())
	fmt.Println(Beijing)

	Beijing.SetSeason(NewSummer())
	fmt.Println(Beijing)

	Beijing.SetSeason(NewAutumn())
	fmt.Println(Beijing)

	Beijing.SetSeason(NewWinter())
	fmt.Println(Beijing)
}
