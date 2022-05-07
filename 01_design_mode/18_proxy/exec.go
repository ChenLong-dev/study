package proxy

import (
  "fmt"
)

func ExecProxy() {
  proxy := NewHouseProxy(&houseOwner{})
  fmt.Println(proxy.SellHouse("北京市海淀区中关村大街，2号院1号楼4单元502室", "李四"))
}
