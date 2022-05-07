package composite

// Region 行政区，作为组合模式component接口
type Region interface {
   Name() string    // 名称
   Population() int //人口
   GDP() float64    // gdp
}
