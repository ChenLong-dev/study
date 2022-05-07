package flyweight

import (
  "fmt"
)

// Taxi 出租车，享元对象，保存不变的内在属性信息
type Taxi struct {
  licensePlate string // 车牌
  color        string // 颜色
  brand        string // 汽车品牌
  company      string // 所属公司
}

// LocateFor 获取定位信息
func (t *Taxi) LocateFor(monitorMap string, x, y int) string {
  return fmt.Sprintf("%s,对于车牌号%s,%s,%s品牌,所属%s公司,定位(%d,%d)", monitorMap,
    t.licensePlate, t.color, t.brand, t.company, x, y)
}

// taxiFactoryInstance 出租车工厂单例
var taxiFactoryInstance = &TaxiFactory{
  taxis: make(map[string]*Taxi),
}

// GetTaxiFactory 获取出租车工厂单例
func GetTaxiFactory() *TaxiFactory {
  return taxiFactoryInstance
}

// TaxiFactory 出租车工厂类
type TaxiFactory struct {
  taxis map[string]*Taxi // key为车牌号
}

// getTaxi 获取出租车
func (f *TaxiFactory) getTaxi(licensePlate, color, brand, company string) *Taxi {
  if _, ok := f.taxis[licensePlate]; !ok {
    f.taxis[licensePlate] = &Taxi{
      licensePlate: licensePlate,
      color:        color,
      brand:        brand,
      company:      company,
    }
  }
  return f.taxis[licensePlate]
}
