package bridge

// Traffic 交通工具
type Traffic interface {
   Transport() string
}

// airplane 飞机
type airplane struct{}

// Transport 坐飞机
func (a *airplane) Transport() string {
   return "by airplane"
}

// car 汽车
type car struct{}

// Transport 坐汽车
func (t *car) Transport() string {
   return "by car"
}
