# Golang反射

## 反射简介

反射可以在运行期间，操作人员类型的对象。可以通过TypeOf方法获得对象类型。通过ValueOf获得对象值。

## Type接口源码分析

Type接口是对go类型的一种描述，不是所有方法都能调用，要根据具体类型来判断。

```bash
go doc reflect Type
```

```bash
type Type interface {
    
    // 内存字节对齐
    Align() int

    // 结构体内存字节对齐
    FieldAlign() int

    // 结构体实现的方法，公共方法首字母大小
    Method(int) Method

    // 根据方面名获得方法
    MethodByName(string) (Method, bool)

    // 方法数量
    NumMethod() int

        
    // 类型名称
    Name() string

    // 包路径
    PkgPath() string
    
    // 需要存储的字节数
    Size() uintptr
    
    // 类型的字符串描述
    String() string
    
    // Kind returns the specific kind of this type.
    Kind() Kind
    
    // 判断是否实现某个接口 Implements reports whether the type implements  the interface type u.
    Implements(u Type) bool
    
    // 判断该值是否可以赋值给u AssignableTo reports whether a value of the  type is assignable to type u.
    AssignableTo(u Type) bool
    
    // 是否可以转换为u类型ConvertibleTo reports whether a value of the type     is convertible to type u.
    // Even if ConvertibleTo returns true, the conversion may still panic.
    // For example, a slice of type []T is convertible to *[N]T,
    // but the conversion will panic if its length is less than N.
    ConvertibleTo(u Type) bool

    // 是否可以比较 Comparable reports whether values of this type are  comparable.
    // Even if Comparable returns true, the comparison may still panic.
    // For example, values of interface type are comparable,
    // but the comparison will panic if their dynamic type is not   comparable.
    Comparable() bool

    // Methods applicable only to some types, depending on Kind.
    // The methods allowed for each kind are:
    //
    // Int*, Uint*, Float*, Complex*: Bits
    // Array: Elem, Len
    // Chan: ChanDir, Elem
    // Func: In, NumIn, Out, NumOut, IsVariadic.
    // Map: Key, Elem
    // Ptr: Elem
    // Slice: Elem
    // Struct: Field, FieldByIndex, FieldByName, FieldByNameFunc, NumField

    // Bits returns the size of the type in bits.
    // It panics if the type's Kind is not one of the
    // sized or unsized Int, Uint, Float, or Complex kinds.
    // 类型的字节长度
    Bits() int

    // ChanDir returns a channel type's direction.
    // 通道类型的方向 It panics if the type's Kind is not Chan.
    ChanDir() ChanDir

    // For concreteness, if t represents func(x int, y ... float64), then
    //
    // t.NumIn() == 2
    // t.In(0) is the reflect.Type for "int"
    // t.In(1) is the reflect.Type for "[]float64"
    // t.IsVariadic() == true
    //
    // 函数类型是否是可变参数
    IsVariadic() bool

    // 返回类型的元素类型
    // It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
    Elem() Type

    // 返回结构体的第几个字段
    // It panics if the type's Kind is not Struct.
    // It panics if i is not in the range [0, NumField()).
    Field(i int) StructField

    // 根据索引获得字段
    // to the index sequence. It is equivalent to calling Field
    // successively for each index i.
    // It panics if the type's Kind is not Struct.
    FieldByIndex(index []int) StructField

    // 根据名称获得字段
    // and a boolean indicating if the field was found.
    FieldByName(name string) (StructField, bool)

    // FieldByNameFunc returns the struct field with a name
    // that satisfies the match function and a boolean indicating if
    // the field was found.
    //
    // FieldByNameFunc considers the fields in the struct itself
    // and then the fields in any embedded structs, in breadth first order,
    // stopping at the shallowest nesting depth containing one or more
    // fields satisfying the match function. If multiple fields at that     depth
    // satisfy the match function, they cancel each other
    // and FieldByNameFunc returns no match.
    // This behavior mirrors Go's handling of name lookup in
    // structs containing embedded fields.
    FieldByNameFunc(match func(string) bool) (StructField, bool)

    // 函数类型的第几个参数
    // It panics if the type's Kind is not Func.
    // It panics if i is not in the range [0, NumIn()).
    In(i int) Type

    // 返回map类型的key
    // It panics if the type's Kind is not Map.
    Key() Type

    // 返回数组类型的长度
    // It panics if the type's Kind is not Array.
    Len() int

    // 返回结构体类型的字段数
    // It panics if the type's Kind is not Struct.
    NumField() int

    // 返回函数类型的输入参数数量
    // It panics if the type's Kind is not Func.
    NumIn() int

    // 函数类型返回值数量
    // It panics if the type's Kind is not Func.
    NumOut() int

    // 函数类型的第几个返回值
    // It panics if the type's Kind is not Func.
    // It panics if i is not in the range [0, NumOut()).
    Out(i int) Type

    common() *rtype
    uncommon() *uncommonType
}
```

## Value源码分析

```bash
go doc reflect Value
```

## value.go中的函数

调用方法 reflect.XXX

```bash
func Append(s Value, x ...Value) Value
添加值到切片
func AppendSlice(s, t Value) Value
添加一个切片到切片
func Indirect(v Value) Value
返回v的指针值
func MakeChan(typ Type, buffer int) Value
创建一个通道，返回Value类型
func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
// 创建一个函数
func MakeMap(typ Type) Value
创建一个map
func MakeMapWithSize(typ Type, n int) Value
// 创建一个map
func MakeSlice(typ Type, len, cap int) Value
// 创建一个切片
func New(typ Type) Value
创建一个类型 
func NewAt(typ Type, p unsafe.Pointer) Value
在某个指针地址处创建一个类型
func ValueOf(i interface{}) Value
值类型
func Zero(typ Type) Value
零值描述
```

## Value结构体的方法

```bash
Addr() Value
通常用于获取一个指向结构体字段或slice元素为了调用一个方法,需要一个指针receiver。
Bool() bool
返回底层的值,如果v的kind不是bool则会产生panic
Bytes() []byte
返回底层的值,如果v的底层值不是一个字节切片,则会产生panic
b := []byte{'a', 'b'}
fmt.Println(reflect.ValueOf(b).Bytes())

CanAddr() bool
检查v是否是可寻址的
CanSet() bool
检查值是否可被设置,只有可寻址的才能被设置
b := 555
p:=reflect.ValueOf(&b)
v := p.Elem()  
//反射对象 p并不是可寻址的，但是并不希望设置p，（实际上）是 *p。为了获得 p 指向的内容，调用值上的 Elem 方法，从指针间接指向，然后保存反射值的结果叫做 v
v.SetInt(666)
fmt.Println(b)

Call(in []Value) []Value

反射函数的值.并调用
func test(a string) string {
    return a
}
func main() {
    a := "sssssss"
    args := []reflect.Value{reflect.ValueOf(a)}
    c := reflect.ValueOf(test).Call(args)
    fmt.Println(c)
}
CallSlice(in []Value) []Value
同上
Close()
关闭channel,如果不是chan则产生panic
Complex() complex128
返回底层的值,如果值不是一个复数,则产生一个panic
Elem() Value
返回v包含的值,多被用于设置值时的寻址操作
Field(i int) Value
返回结构中索引字段的Value
type A struct {
    a int
    b byte
    c string
}
func main() {
    a := A{}
    fmt.Println(reflect.ValueOf(a).Field(0).Int())
}
FieldByIndex(index []int) Value
同上不过.提供的是一个切片
FieldByName(name string) Value
通过字段名查找
FieldByNameFunc(match func(string) bool) Value
通过函数名查找
Float() float64
返回底层的值,如果值不是一个float,则产生一个panic
Index(i int) Value
如果kind不是array或者sliece则差生panic,将其中的元素返回为Value
Int() int64
返回底层的值,如果值不是一个int,则产生一个panic
CanInterface() bool
如果接口能被使用,则返回true
Interface() (i interface{})
返回V作为interface{}的当前值
InterfaceData() [2]uintptr
如果kind不是一个接口则会产生panic
IsNil() bool
如果v是一个nil,则返回true
IsValid() bool
如果v代表一个值,则返回true
Kind() Kind
返回v的种类
Len() int
返回v的长度
MapIndex(key Value) Value
如果是一个map,根据key反射其键值的Value
MapKeys() []Value
返回map的所有key
Method(i int) Value
按索引反射结构某个方法的值
NumMethod() int
统计结构方法数量
MethodByName(name string) Value
反射方法的值根据方法名
NumField() int
反射一个结构的字段数
OverflowComplex(x complex128) bool
覆盖复数
OverflowFloat(x float64) bool
覆盖浮点数
overflowFloat32(x float64) bool
OverflowInt(x int64) bool
OverflowUint(x uint64) bool
Pointer() uintptr
反射一个指针的值.返回一个指针的整型值
Recv() (x Value, ok bool)
用于channel
Send(x Value)
用于channel
Set(x Value)
如果v可设置,则设置一个v的值
SetBool(x bool)
如果v可设置,且是bool,则设置一个v的值
SetBytes(x []byte)
SetComplex(x complex128)
SetFloat(x float64)
SetInt(x int64)
SetLen(n int)
SetMapIndex(key, val Value)
SetUint(x uint64)
SetPointer(x unsafe.Pointer)
SetString(x string)
Slice(beg, end int) Value
如果底层是slice.则返回值.
String() string
如果底层是字符串.则返回字符串
TryRecv() (x Value, ok bool)
用于channel,接受
TrySend(x Value) bool
用于channel,发送
Type() Type
返回type
Uint() uint64
如果底层是Uint.则返回uint
UnsafeAddr() uintptr
```

## 实例演示

### 接口值到反射对象

```bash
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x int = 1
    fmt.Println("type: ", reflect.TypeOf(x))
}

type:  int
```

TypeOf函数的定义如下，参数为接口类型，返回值为类型

```bash
func TypeOf(i interface {}) Type
ValueOf`函数的定义如下，参数为接口类型，返回值为`Value
var x int = 1
fmt.Println("value: ", reflect.ValueOf(x))

value:  <int Value>
```

可以通过Kind函数来检查类型，

```bash
fmt.Println("Kind:  ", reflect.ValueOf(x).Kind())
fmt.Println("Kind is Int? ", reflect.ValueOf(x).Kind() == reflect.Int)

Kind:   int
Kind is Int?  true
```

### 反射对象到接口值

通过Interface函数可以实现反射对象到接口值的转换，

```bash
func (v Value) Interface() interface {}

// Interface 以 interface{} 返回 v 的值
y := v.Interface().(float64)
fmt.Println(y)
```

### 修改反射对象

修改反射对象的前提条件是其值必须是可设置的

```bash
var x float64 = 3.4
v := reflect.ValueOf(x)
v.SetFloat(7.3) // Error: panic
```

为了避免这个问题，需要使用CanSet函数来检查该值的设置性，

```bash
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("settability of v: ", v.CanSet())

settability of v: false
```

那么如何才能设置该值呢？ 这里需要考虑一个常见的问题，参数传递，传值还是传引用或地址？ 在上面的例子中，我们使用的是reflect.ValueOf(x)，这是一个值传递，传递的是x的值的一个副本，不是x本身，因此更新副本中的值是不允许的。如果使用reflect.ValueOf(&x)来替换刚才的值传递，就可以实现值的修改。

```bash
var x float64 = 3.4
p := reflect.ValueOf(&x) // 获取x的地址
fmt.Println("settability of p: ", p.CanSet())
v := p.Elem()
fmt.Println("settability of v: ", v.CanSet())
v.SetFloat(7.1)
fmt.Println(v.Interface())
fmt.Println(x)

settability of p: false
settability of v: true
7.1
7.1
```

### 获取结构体标签

首先介绍如何遍历结构体字段内容， 假设结构体如下，

```bash
type T struct {
    A int
    B string
}

t := T{1, "golang"}
```

从而，通过反射来遍历所有的字段内容

```bash
s := reflect.ValueOf(&t).Elem()
typeOfT := s.Type()
for i := 0; i < s.NumField(); i++ {
    f := s.Field(i)
    fmt.Printf("%d %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
}

0 A int = 1
1 B string = golang
```

接下来，如何获取结构体的标签内容?

```bash
func main() {
    type S struct {
        F string `species:"gopher" color:"blue"`
    }

    s := S{}
    st := reflect.TypeOf(s)
    field := st.Field(0)
    fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
}
```

## interface{}到函数反射

一般情况下，为了存储多个函数值，一般采用map来存储。其中key为函数名称，而value为相应的处理函数。 在这里需要定义好函数类型，但是函数的参数以及返回类型就需要是统一的，如下

```bash
package main

import "fmt"

func say(text string) {
    fmt.Println(text)
}

func main() {
    var funcMap = make(map[string]func(string))
    funcMap["say"] = say
    funcMap["say"]("hello")
}
```

如果希望map可以存储任意类型的函数（参数不同，返回值不同），那么就需要用interface{}而不是func(param...)来定义value。

```bash
package main

import "fmt"

func say(text string) {
    fmt.Println(text)
}

func main() {
    var funcMap = make(map[string]interface{})
    funcMap["say"] = say
    funcMap["say"]("hello")
}

cannot call non-function funcMap["say"] (type interface {})
```

直接调用会报错，提示不能调用interface{}类型的函数。

这时，需要利用reflect把函数从interface转换到函数来使用，

```bash
package main

import (
    "fmt"
    "reflect"
)

func say(text string) {
    fmt.Println(text)
}

func Call(m map[string]interface{}, name string, params ... interface{}) (result []reflect.Value) {
    f := reflect.ValueOf(m[name])
    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }
    result = f.Call(in)
    return
}

func main() {
    var funcMap = make(map[string]interface{})
    funcMap["say"] = say
    Call(funcMap, "say", "hello")
}
```
