package bridge

import "fmt"

// Location 地点
type Location interface {
  Name() string // 地点名称
  PlaySports() string // 参与运动
}

// namedLocation 被命名的地点，统一引用此类型，声明名字字段及获取方法
type namedLocation struct {
  name string
}

// Name 获取地点名称
func (n namedLocation) Name() string {
  return n.name
}

// seaside 海边
type seaside struct {
  namedLocation
}

// NewSeaside 创建指定名字的海边，比如三亚湾
func NewSeaside(name string) *seaside {
  return &seaside{
    namedLocation: namedLocation{
      name: name,
    },
  }
}

// PlaySports 海边可以冲浪
func (s *seaside) PlaySports() string {
  return fmt.Sprintf("surfing")
}

// mountain 山
type mountain struct {
  namedLocation
}

// NewMountain 创建指定名字的山，比如泰山
func NewMountain(name string) *mountain {
  return &mountain{
    namedLocation: namedLocation{
      name: name,
    },
  }
}

// PlaySports 可以爬山
func (m *mountain) PlaySports() string {
  return fmt.Sprintf("climbing")
}

// desert 荒漠
type desert struct {
  namedLocation
}

// NewDesert 创建指定名字的荒漠，比如罗布泊
func NewDesert(name string) *desert {
  return &desert{
    namedLocation: namedLocation{
      name: name,
    },
  }
}

// PlaySports 荒漠可以徒步穿越
func (d *desert) PlaySports() string {
  return fmt.Sprintf("trekking")
}
