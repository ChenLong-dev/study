<!--
 * @Author: ChenLong longchen2008@126.com
 * @Date: 2022-06-04 22:25:02
 * @LastEditors: ChenLong longchen2008@126.com
 * @LastEditTime: 2022-06-04 22:48:56
 * @FilePath: \study\03_golang_tech_stack\special\error\readme.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# Golang错误处理

我们经常会看到golang，标准库中的方法，都返回两个值，其中一个是error类型的。这就是golang错误处理的简洁逻辑。例如：

```bash
func Open(name string) (*File, error) {
    return OpenFile(name, O_RDONLY, 0)
}

os.Open("a.txt")
```

## Golang 错误处理源码分析

### 接口 builtin.go

```bash
// error接口 builtin.go
type error interface {
    Error() string
}
```

### errors.go

```bash
package errors

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
    return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

### wrap.go

```bash
package errors

// 拆开一个被封装的错误error
func Unwrap(err error) error {
    u, ok := err.(interface {
        Unwrap() error
    })
    if !ok {
        return nil
    }
    return u.Unwrap()
}

// 判断被包装的 error 是否是包含指定错误
func Is(err, target error) bool {
    if target == nil {
        return err == target
    }

    isComparable := reflectlite.TypeOf(target).Comparable()
    for {
            if isComparable && err == target {
                return true
            }
            if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
                return true
            }
            // TODO: consider supporting target.Is(err). This would allow
            // user-definable predicates, but also may allow for coping with sloppy
            // APIs, thereby making it easier to get away with them.
            if err = Unwrap(err); err == nil {
                return false
            }
    }
}

// 提取指定类型的错误
func As(err error, target interface{}) bool {
    if target == nil {
        panic("errors: target cannot be nil")
    }
    val := reflectlite.ValueOf(target)
    typ := val.Type()
    if typ.Kind() != reflectlite.Ptr || val.IsNil() {
        panic("errors: target must be a non-nil pointer")
    }
    targetType := typ.Elem()
    if targetType.Kind() != reflectlite.Interface && !targetType.Implements(errorType) {
        panic("errors: *target must be interface or implement error")
    }
    for err != nil {
        if reflectlite.TypeOf(err).AssignableTo(targetType) {
            val.Elem().Set(reflectlite.ValueOf(err))
            return true
        }
        if x, ok := err.(interface{ As(interface{}) bool }); ok && x.As(target) {
            return true
        }
        err = Unwrap(err)
    }
    return false
}

var errorType = reflectlite.TypeOf((*error)(nil)).Elem()
```

## 实例演示

### 返回error

```bash
package main

import (
    "errors"
    "fmt"
)

func Div(a int, b int) (int, error) {
    if b == 0 {
        return -1, errors.New("除数不能为零")
    } else {
        return a / b, nil
    }
}

func main() {
    _, err := Div(1, 0)
    fmt.Printf("err.Error(): %v\n", err.Error())
    fmt.Printf("err: %v\n", err)
}

```

### 返回error（错误拼接）

```bash
# 错误拼接
fmt.Errorf("not found mongodb config: %s", "出现错误")
```

## 自定义 error

### 方式一：fmt.Errorf

使用 %w 参数返回一个被包装的 error

```bash
err1 := errors.New("new error")
err2 := fmt.Errorf("err2: [%w]", err1)
err3 := fmt.Errorf("err3: [%w]", err2)
fmt.Println(err3)
// output
err3: [err2: [new error]]
```

err2 就是一个合法的被包装的 error，同样地，err3 也是一个被包装的 error，如此可以一直套下去。

### 方式二：自定义 struct

```bash
type WarpError struct {
    msg string
    err error
}

func (e *WarpError) Error() string {
    return e.msg
}

func (e *WrapError) Unwrap() error {
    return e.err
}
```

之前看过源码的同学可能已经知道了，这就是 fmt/errors.go 中关于 warp 的结构。 就，很简单。自定义一个实现了 Unwrap 方法的 struct 就可以了。

## 拆开一个被包装的 error

### errors.Unwrap

```bash
err1 := errors.New("new error")
err2 := fmt.Errorf("err2: [%w]", err1)
err3 := fmt.Errorf("err3: [%w]", err2)

fmt.Println(errors.Unwrap(err3))
fmt.Println(errors.Unwrap(errors.Unwrap(err3)))
// output
err2: [new error]
new error
```

## 判断被包装的 error 是否是包含指定错误

### errors.Is

当多层调用返回的错误被一次次地包装起来，我们在调用链上游拿到的错误如何判断是否是底层的某个错误呢？ 它递归调用 Unwrap 并判断每一层的 err 是否相等，如果有任何一层 err 和传入的目标错误相等，则返回 true。

```bash
err1 := errors.New("new error")
err2 := fmt.Errorf("err2: [%w]", err1)
err3 := fmt.Errorf("err3: [%w]", err2)

fmt.Println(errors.Is(err3, err2))
fmt.Println(errors.Is(err3, err1))
// output
true
true
```

## 提取指定类型的错误

### errors.As

这个和上面的 errors.Is 大体上是一样的，区别在于 Is 是严格判断相等，即两个 error 是否相等。 而 As 则是判断类型是否相同，并提取第一个符合目标类型的错误，用来统一处理某一类错误。

```bash
type ErrorString struct {
    s string
}

func (e *ErrorString) Error() string {
    return e.s
}

var targetErr *ErrorString
err := fmt.Errorf("new error:[%w]", &ErrorString{s:"target err"})
fmt.Println(errors.As(err, &targetErr))
// output
true
```
