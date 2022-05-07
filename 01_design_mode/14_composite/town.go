package composite

// town 区镇，组合模式中相当于叶子节点
type town struct {
  name       string
  population int
  gdp        float64
}

// NewTown 创建区镇，根据名称、人口、GDP
func NewTown(name string, population int, gdp float64) *town {
  return &town{
    name:       name,
    population: population,
    gdp:        gdp,
  }
}

func (c *town) Name() string {
  return c.name
}

func (c *town) Population() int {
  return c.population
}

func (c *town) GDP() float64 {
  return c.gdp
}
