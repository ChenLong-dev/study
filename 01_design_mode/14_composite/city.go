package composite

// cities 市，包括县市或者地市，组合模式中相当于composite
type cities struct {
  name    string
  regions map[string]Region
}

// NewCities 创建一个市
func NewCities(name string) *cities {
  return &cities{
    name:    name,
    regions: make(map[string]Region),
  }
}

func (c *cities) Name() string {
  return c.name
}

func (c *cities) Population() int {
  sum := 0
  for _, r := range c.regions {
    sum += r.Population()
  }
  return sum
}

func (c *cities) GDP() float64 {
  sum := 0.0
  for _, r := range c.regions {
    sum += r.GDP()
  }
  return sum
}

// Add 添加多个行政区
func (c *cities) Add(regions ...Region) {
  for _, r := range regions {
    c.regions[r.Name()] = r
  }
}

// Remove 递归删除行政区
func (c *cities) Remove(name string) {
  for n, r := range c.regions {
    if n == name {
      delete(c.regions, name)
      return
    }
    if city, ok := r.(*cities); ok {
      city.Remove(name)
    }
  }
}

func (c *cities) Regions() map[string]Region {
  return c.regions
}
