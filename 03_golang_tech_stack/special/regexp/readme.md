<!--
 * @Author: ChenLong longchen2008@126.com
 * @Date: 2022-06-04 17:36:13
 * @LastEditors: ChenLong longchen2008@126.com
 * @LastEditTime: 2022-06-04 17:57:00
 * @FilePath: \study\03_golang_tech_stack\regexp\readme.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# Golang正则表达式

正则表达式，（英语：Regular Expression，在代码中常简写为regex、regexp或RE），正则表达式通常被用来检索、替换那些符合某个模式(规则)的文本。例如：用户注册，邮箱验证、手机号码验证、爬虫字符串匹配等。

## Golang正则表达式语法

```bash
go doc regexp/synt
```

## 单一

```bash
.                   匹配任意一个字符，如果设置 s = true，则可以匹配换行符

[字符类]            匹配“字符类”中的一个字符，“字符类”见后面的说明
[^字符类]           匹配“字符类”外的一个字符，“字符类”见后面的说明

\小写Perl标记       匹配“Perl类”中的一个字符，“Perl类”见后面的说明
\大写Perl标记       匹配“Perl类”外的一个字符，“Perl类”见后面的说明

[:ASCII类名:]       匹配“ASCII类”中的一个字符，“ASCII类”见后面的说明
[:^ASCII类名:]      匹配“ASCII类”外的一个字符，“ASCII类”见后面的说明

\pUnicode普通类名   匹配“Unicode类”中的一个字符(仅普通类)，“Unicode类”见后面的说明
\PUnicode普通类名   匹配“Unicode类”外的一个字符(仅普通类)，“Unicode类”见后面的说明

\p{Unicode类名}     匹配“Unicode类”中的一个字符，“Unicode类”见后面的说明
\P{Unicode类名}     匹配“Unicode类”外的一个字符，“Unicode类”见后面的说明
```

## 可以将“命名字符类”作为“字符类”的元素

```bash
[\d]           匹配数字 (相当于 \d)
[^\d]          匹配非数字 (相当于 \D)
[\D]           匹配非数字 (相当于 \D)
[^\D]          匹配数字 (相当于 \d)
[[:name:]]     命名的“ASCII 类”包含在“字符类”中 (相当于 [:name:])
[^[:name:]]    命名的“ASCII 类”不包含在“字符类”中 (相当于 [:^name:])
[\p{Name}]     命名的“Unicode 类”包含在“字符类”中 (相当于 \p{Name})
[^\p{Name}]    命名的“Unicode 类”不包含在“字符类”中 (相当于 \P{Name})
```

### 说明

“字符类”取值如下（“字符类”包含“Perl类”、“ASCII类”、“Unicode类”）：

```bash
x                    单个字符
A-Z                  字符范围(包含首尾字符)
\小写字母            Perl类
[:ASCII类名:]        ASCII类
\p{Unicode脚本类名}  Unicode类 (脚本类)
\pUnicode普通类名    Unicode类 (普通类)
```

#### “Perl 类”取值如下

```bash
\d             数字 (相当于 [0-9])
\D             非数字 (相当于 [^0-9])
\s             空白 (相当于 [\t\n\f\r ])
\S             非空白 (相当于[^\t\n\f\r ])
\w             单词字符 (相当于 [0-9A-Za-z_])
\W             非单词字符 (相当于 [^0-9A-Za-z_])
```

#### “ASCII 类”取值如下

```bash
[:alnum:]      字母数字 (相当于 [0-9A-Za-z])
[:alpha:]      字母 (相当于 [A-Za-z])
[:ascii:]      ASCII 字符集 (相当于 [\x00-\x7F])
[:blank:]      空白占位符 (相当于 [\t ])
[:cntrl:]      控制字符 (相当于 [\x00-\x1F\x7F])
[:digit:]      数字 (相当于 [0-9])
[:graph:]      图形字符 (相当于 [!-~])
[:lower:]      小写字母 (相当于 [a-z])
[:print:]      可打印字符 (相当于 [ -~] 相当于 [ [:graph:]])
[:punct:]      标点符号 (相当于 [!-/:-@[-反引号{-~])
[:space:]      空白字符(相当于 [\t\n\v\f\r ])
[:upper:]      大写字母(相当于 [A-Z])
[:word:]       单词字符(相当于 [0-9A-Za-z_])
[:xdigit:]     16 進制字符集(相当于 [0-9A-Fa-f])
```

## 复合

```bash
xy             匹配 xy（x 后面跟随 y）
x|y            匹配 x 或 y (优先匹配 x)
```

## 重复

```bash
x*              匹配零个或多个 x，优先匹配更多(贪婪)
x+              匹配一个或多个 x，优先匹配更多(贪婪)
x?              匹配零个或一个 x，优先匹配一个(贪婪)
x{n,m}          匹配 n 到 m 个 x，优先匹配更多(贪婪)
x{n,}           匹配 n 个或多个 x，优先匹配更多(贪婪)
x{n}            只匹配 n 个 x
x*?             匹配零个或多个 x，优先匹配更少(非贪婪)
x+?             匹配一个或多个 x，优先匹配更少(非贪婪)
x??             匹配零个或一个 x，优先匹配零个(非贪婪)
x{n,m}?         匹配 n 到 m 个 x，优先匹配更少(非贪婪)
x{n,}?          匹配 n 个或多个 x，优先匹配更少(非贪婪)
x{n}?           只匹配 n 个 x
```
贪婪模式在整个表达式匹配成功的前提下，尽可能多的匹配，而非贪婪模式在整个表达式匹配成功的前提下，尽可能少的匹配。

## 分组

```bash
(子表达式)            被捕获的组，该组被编号 (子匹配)
(?P<命名>子表达式)    被捕获的组，该组被编号且被命名 (子匹配)
(?:子表达式)          非捕获的组 (子匹配)
(?标记)               在组内设置标记，非捕获，标记影响当前组后的正则表达式
(?标记:子表达式)      在组内设置标记，非捕获，标记影响当前组内的子表达式

标记的语法是：
xyz  (设置 xyz 标记)
-xyz (清除 xyz 标记)
xy-z (设置 xy 标记, 清除 z 标记)

可以设置的标记有：
i              不区分大小写 (默认为 false)
m              多行模式：让 ^ 和 $ 匹配整个文本的开头和结尾，而非行首和尾(默认为 false)
s              让 . 匹配 \n (默认为 false)
U              非贪婪模式：交换 x* 和 x*? 等的含义 (默认为 false)
```

## 位置标记

```bash
^              如果标记 m=true 则匹配行首，否则匹配整个文本的开头（m 默认为 false）
$              如果标记 m=true 则匹配行尾，否则匹配整个文本的结尾（m 默认为 false）
\A             匹配整个文本的开头，忽略 m 标记
\b             匹配单词边界
\B             匹配非单词边界
\z             匹配整个文本的结尾，忽略 m 标记
```

## 转义序列

```bash
\a             匹配响铃符    （相当于 \x07）
               注意：正则表达式中不能使用 \b 匹配退格符，因为 \b 被用来匹配单词边界，
               可以使用 \x08 表示退格符。
\f             匹配换页符    （相当于 \x0C）
\t             匹配横向制表符（相当于 \x09）
\n             匹配换行符    （相当于 \x0A）
\r             匹配回车符    （相当于 \x0D）
\v             匹配纵向制表符（相当于 \x0B）
\123           匹配 8  進制编码所代表的字符（必须是 3 位数字）
\x7F           匹配 16 進制编码所代表的字符（必须是 3 位数字）
\x{10FFFF}     匹配 16 進制编码所代表的字符（最大值 10FFFF  ）
\Q...\E        匹配 \Q 和 \E 之间的文本，忽略文本中的正则语法

\\             匹配字符 \
\^             匹配字符 ^
\$             匹配字符 $
\.             匹配字符 .
\*             匹配字符 *
\+             匹配字符 +
\?             匹配字符 ?
\{             匹配字符 {
\}             匹配字符 }
\(             匹配字符 (
\)             匹配字符 )
\[             匹配字符 [
\]             匹配字符 ]
\|             匹配字符 |
```

## Golang正则入门实例

```bash
package main

import (
    "fmt"
    "regexp"
)

func main() {
    // 匹配
    match, _ := regexp.MatchString("H(.*)!", "Hello world!")
    fmt.Println(match) // true
    fmt.Println("---------------")
    // 查找
    re := regexp.MustCompile(`foo.?`)
    fmt.Printf("%q\n", re.FindString("seafood fool"))
    fmt.Printf("%q\n", re.FindString("meat"))
    fmt.Println("---------------")
}
```

### 运行结果

```bash
true
---------------
"food"
""
---------------
```

## 综合实例演示

```bash
// 判断在 b 中能否找到正则表达式 pattern 所匹配的子串
// pattern：要查找的正则表达式
// b：要在其中进行查找的 []byte
// matched：返回是否找到匹配项
// err：返回查找过程中遇到的任何错误
// 此函数通过调用 Regexp 的方法实现
func Match(pattern string, b []byte) (matched bool, err error)

func main() {
    fmt.Println(regexp.Match("H.* ", []byte("Hello World!")))
    // true 
}

------------------------------------------------------------

// 判断在 r 中能否找到正则表达式 pattern 所匹配的子串
// pattern：要查找的正则表达式
// r：要在其中进行查找的 RuneReader 接口
// matched：返回是否找到匹配项
// err：返回查找过程中遇到的任何错误
// 此函数通过调用 Regexp 的方法实现
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)

func main() {
    r := bytes.NewReader([]byte("Hello World!"))
    fmt.Println(regexp.MatchReader("H.* ", r))
    // true 
}

------------------------------------------------------------

// 判断在 s 中能否找到正则表达式 pattern 所匹配的子串
// pattern：要查找的正则表达式
// r：要在其中进行查找的字符串
// matched：返回是否找到匹配项
// err：返回查找过程中遇到的任何错误
// 此函数通过调用 Regexp 的方法实现
func MatchString(pattern string, s string) (matched bool, err error)

func main() {
    fmt.Println(regexp.Match("H.* ", "Hello World!"))
    // true 
}

------------------------------------------------------------

// QuoteMeta 将字符串 s 中的“特殊字符”转换为其“转义格式”
// 例如，QuoteMeta（`[foo]`）返回`\[foo\]`。
// 特殊字符有：\.+*?()|[]{}^$
// 这些字符用于实现正则语法，所以当作普通字符使用时需要转换
func QuoteMeta(s string) string

func main() {
    fmt.Println(regexp.QuoteMeta("(?P:Hello) [a-z]"))
    // \(\?P:Hello\) \[a-z\]
}

------------------------------------------------------------

// Regexp 结构表示一个编译后的正则表达式
// Regexp 的公开接口都是通过方法实现的
// 多个 goroutine 并发使用一个 RegExp 是安全的
type Regexp struct {
    // 私有字段
}

// 通过 Complite、CompilePOSIX、MustCompile、MustCompilePOSIX 
// 四个函数可以创建一个 Regexp 对象

------------------------------------------------------------

// Compile 用来解析正则表达式 expr 是否合法，如果合法，则返回一个 Regexp 对象
// Regexp 对象可以在任意文本上执行需要的操作
func Compile(expr string) (*Regexp, error)

func main() {
    reg, err := regexp.Compile(`\w+`)
    fmt.Printf("%q,%v\n", reg.FindString("Hello World!"), err)
    // "Hello",
}

------------------------------------------------------------

// CompilePOSIX 的作用和 Compile 一样
// 不同的是，CompilePOSIX 使用 POSIX 语法，
// 同时，它采用最左最长方式搜索，
// 而 Compile 采用最左最短方式搜索
// POSIX 语法不支持 Perl 的语法格式：\d、\D、\s、\S、\w、\W
func CompilePOSIX(expr string) (*Regexp, error)

func main() {
    reg, err := regexp.CompilePOSIX(`[[:word:]]+`)
    fmt.Printf("%q,%v\n", reg.FindString("Hello World!"), err)
    // "Hello"
}

------------------------------------------------------------

// MustCompile 的作用和 Compile 一样
// 不同的是，当正则表达式 str 不合法时，MustCompile 会抛出异常
// 而 Compile 仅返回一个 error 值
func MustCompile(str string) *Regexp

func main() {
    reg := regexp.MustCompile(`\w+`)
    fmt.Println(reg.FindString("Hello World!"))
    // Hello
}

------------------------------------------------------------

// MustCompilePOSIX 的作用和 CompilePOSIX 一样
// 不同的是，当正则表达式 str 不合法时，MustCompilePOSIX 会抛出异常
// 而 CompilePOSIX 仅返回一个 error 值
func MustCompilePOSIX(str string) *Regexp

func main() {
    reg := regexp.MustCompilePOSIX(`[[:word:]].+ `)
    fmt.Printf("%q\n", reg.FindString("Hello World!"))
    // "Hello "
}

------------------------------------------------------------

// 在 b 中查找 re 中编译好的正则表达式，并返回第一个匹配的内容
func (re *Regexp) Find(b []byte) []byte

func main() {
    reg := regexp.MustCompile(`\w+`)
    fmt.Printf("%q", reg.Find([]byte("Hello World!")))
    // "Hello"
}

------------------------------------------------------------

// 在 s 中查找 re 中编译好的正则表达式，并返回第一个匹配的内容
func (re *Regexp) FindString(s string) string

func main() {
    reg := regexp.MustCompile(`\w+`)
    fmt.Println(reg.FindString("Hello World!"))
    // "Hello"
}

------------------------------------------------------------

// 在 b 中查找 re 中编译好的正则表达式，并返回所有匹配的内容
// {{匹配项}, {匹配项}, ...}
// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
func (re *Regexp) FindAll(b []byte, n int) [][]byte

func main() {
    reg := regexp.MustCompile(`\w+`)
    fmt.Printf("%q", reg.FindAll([]byte("Hello World!"), -1))
    // ["Hello" "World"]
}

------------------------------------------------------------

// 在 s 中查找 re 中编译好的正则表达式，并返回所有匹配的内容
// {匹配项, 匹配项, ...}
// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
func (re *Regexp) FindAllString(s string, n int) []string

func main() {
    reg := regexp.MustCompile(`\w+`)
    fmt.Printf("%q", reg.FindAllString("Hello World!", -1))
    // ["Hello" "World"]
}

------------------------------------------------------------

// 在 b 中查找 re 中编译好的正则表达式，并返回第一个匹配的位置
// {起始位置, 结束位置}
func (re *Regexp) FindIndex(b []byte) (loc []int)

func main() {
    reg := regexp.MustCompile(`\w+`)
    fmt.Println(reg.FindIndex([]byte("Hello World!")))
    // [0 5]
}

------------------------------------------------------------

// 在 s 中查找 re 中编译好的正则表达式，并返回第一个匹配的位置
// {起始位置, 结束位置}
func (re *Regexp) FindStringIndex(s string) (loc []int)

func main() {
    reg := regexp.MustCompile(`\w+`)
    fmt.Println(reg.FindStringIndex("Hello World!"))
    // [0 5]
}

------------------------------------------------------------

// 在 r 中查找 re 中编译好的正则表达式，并返回第一个匹配的位置
// {起始位置, 结束位置}
func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)

func main() {
    r := bytes.NewReader([]byte("Hello World!"))
    reg := regexp.MustCompile(`\w+`)
    fmt.Println(reg.FindReaderIndex(r))
    // [0 5]
}

------------------------------------------------------------

// 在 b 中查找 re 中编译好的正则表达式，并返回所有匹配的位置
// {{起始位置, 结束位置}, {起始位置, 结束位置}, ...}
// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int

func main() {
    reg := regexp.MustCompile(`\w+`)
    fmt.Println(reg.FindAllIndex([]byte("Hello World!"), -1))
    // [[0 5] [6 11]]
}

------------------------------------------------------------

// 在 s 中查找 re 中编译好的正则表达式，并返回所有匹配的位置
// {{起始位置, 结束位置}, {起始位置, 结束位置}, ...}
// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int

func main() {
    reg := regexp.MustCompile(`\w+`)
    fmt.Println(reg.FindAllStringIndex("Hello World!", -1))
    // [[0 5] [6 11]]
}

------------------------------------------------------------

// 在 b 中查找 re 中编译好的正则表达式，并返回第一个匹配的内容
// 同时返回子表达式匹配的内容
// {{完整匹配项}, {子匹配项}, {子匹配项}, ...}
func (re *Regexp) FindSubmatch(b []byte) [][]byte

func main() {
    reg := regexp.MustCompile(`(\w)(\w)+`)
    fmt.Printf("%q", reg.FindSubmatch([]byte("Hello World!")))
    // ["Hello" "H" "o"]
}

------------------------------------------------------------

// 在 s 中查找 re 中编译好的正则表达式，并返回第一个匹配的内容
// 同时返回子表达式匹配的内容
// {完整匹配项, 子匹配项, 子匹配项, ...}
func (re *Regexp) FindStringSubmatch(s string) []string

func main() {
    reg := regexp.MustCompile(`(\w)(\w)+`)
    fmt.Printf("%q", reg.FindStringSubmatch("Hello World!"))
    // ["Hello" "H" "o"]
}

------------------------------------------------------------

// 在 b 中查找 re 中编译好的正则表达式，并返回所有匹配的内容
// 同时返回子表达式匹配的内容
// {
//     {{完整匹配项}, {子匹配项}, {子匹配项}, ...},
//     {{完整匹配项}, {子匹配项}, {子匹配项}, ...},
//     ...
// }
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte

func main() {
    reg := regexp.MustCompile(`(\w)(\w)+`)
    fmt.Printf("%q", reg.FindAllSubmatch([]byte("Hello World!"), -1))
    // [["Hello" "H" "o"] ["World" "W" "d"]]
}

------------------------------------------------------------

// 在 s 中查找 re 中编译好的正则表达式，并返回所有匹配的内容
// 同时返回子表达式匹配的内容
// {
//     {完整匹配项, 子匹配项, 子匹配项, ...},
//     {完整匹配项, 子匹配项, 子匹配项, ...},
//     ...
// }
// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string

func main() {
    reg := regexp.MustCompile(`(\w)(\w)+`)
    fmt.Printf("%q", reg.FindAllStringSubmatch("Hello World!", -1))
    // [["Hello" "H" "o"] ["World" "W" "d"]]
}

------------------------------------------------------------

// 在 b 中查找 re 中编译好的正则表达式，并返回第一个匹配的位置
// 同时返回子表达式匹配的位置
// {完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...}
func (re *Regexp) FindSubmatchIndex(b []byte) []int

func main() {
    reg := regexp.MustCompile(`(\w)(\w)+`)
    fmt.Println(reg.FindSubmatchIndex([]byte("Hello World!")))
    // [0 5 0 1 4 5]
}

------------------------------------------------------------

// 在 s 中查找 re 中编译好的正则表达式，并返回第一个匹配的位置
// 同时返回子表达式匹配的位置
// {完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...}
func (re *Regexp) FindStringSubmatchIndex(s string) []int

func main() {
    reg := regexp.MustCompile(`(\w)(\w)+`)
    fmt.Println(reg.FindStringSubmatchIndex("Hello World!"))
    // [0 5 0 1 4 5]
}

------------------------------------------------------------

// 在 r 中查找 re 中编译好的正则表达式，并返回第一个匹配的位置
// 同时返回子表达式匹配的位置
// {完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...}
func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int

func main() {
    r := bytes.NewReader([]byte("Hello World!"))
    reg := regexp.MustCompile(`(\w)(\w)+`)
    fmt.Println(reg.FindReaderSubmatchIndex(r))
    // [0 5 0 1 4 5]
}

------------------------------------------------------------

// 在 b 中查找 re 中编译好的正则表达式，并返回所有匹配的位置
// 同时返回子表达式匹配的位置
// {
//     {完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...}, 
//     {完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...}, 
//     ...
// }
// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int

func main() {
    reg := regexp.MustCompile(`(\w)(\w)+`)
    fmt.Println(reg.FindAllSubmatchIndex([]byte("Hello World!"), -1))
    // [[0 5 0 1 4 5] [6 11 6 7 10 11]]
}

------------------------------------------------------------

// 在 s 中查找 re 中编译好的正则表达式，并返回所有匹配的位置
// 同时返回子表达式匹配的位置
// {
//     {完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...}, 
//     {完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...}, 
//     ...
// }
// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int

func main() {
    reg := regexp.MustCompile(`(\w)(\w)+`)
    fmt.Println(reg.FindAllStringSubmatchIndex("Hello World!", -1))
    // [[0 5 0 1 4 5] [6 11 6 7 10 11]]
}

------------------------------------------------------------

// 将 template 的内容经过处理后，追加到 dst 的尾部。
// template 中要有 $1、$2、${name1}、${name2} 这样的“分组引用符”
// match 是由 FindSubmatchIndex 方法返回的结果，里面存放了各个分组的位置信息
// 如果 template 中有“分组引用符”，则以 match 为标准，
// 在 src 中取出相应的子串，替换掉 template 中的 $1、$2 等引用符号。
func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte

func main() {
    reg := regexp.MustCompile(`(\w+),(\w+)`)
    src := []byte("Golang,World!")           // 源文本
    dst := []byte("Say: ")                   // 目标文本
    template := []byte("Hello $1, Hello $2") // 模板
    match := reg.FindSubmatchIndex(src)      // 解析源文本
    // 填写模板，并将模板追加到目标文本中
    fmt.Printf("%q", reg.Expand(dst, template, src, match))
    // "Say: Hello Golang, Hello World"
}

------------------------------------------------------------

// 功能同 Expand 一样，只不过参数换成了 string 类型
func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte

func main() {
    reg := regexp.MustCompile(`(\w+),(\w+)`)
    src := "Golang,World!"                    // 源文本
    dst := []byte("Say: ")                    // 目标文本（可写）
    template := "Hello $1, Hello $2"          // 模板
    match := reg.FindStringSubmatchIndex(src) // 解析源文本
    // 填写模板，并将模板追加到目标文本中
    fmt.Printf("%q", reg.ExpandString(dst, template, src, match))
    // "Say: Hello Golang, Hello World"
}

------------------------------------------------------------

// LiteralPrefix 返回所有匹配项都共同拥有的前缀（去除可变元素）
// prefix：共同拥有的前缀
// complete：如果 prefix 就是正则表达式本身，则返回 true，否则返回 false
func (re *Regexp) LiteralPrefix() (prefix string, complete bool)

func main() {
    reg := regexp.MustCompile(`Hello[\w\s]+`)
    fmt.Println(reg.LiteralPrefix())
    // Hello false
    reg = regexp.MustCompile(`Hello`)
    fmt.Println(reg.LiteralPrefix())
    // Hello true
}

------------------------------------------------------------

// 切换到“贪婪模式”
func (re *Regexp) Longest()

func main() {
    text := `Hello World, 123 Go!`
    pattern := `(?U)H[\w\s]+o` // 正则标记“非贪婪模式”(?U)
    reg := regexp.MustCompile(pattern)
    fmt.Printf("%q\n", reg.FindString(text))
    // Hello
    reg.Longest() // 切换到“贪婪模式”
    fmt.Printf("%q\n", reg.FindString(text))
    // Hello Wo
}

------------------------------------------------------------

// 判断在 b 中能否找到匹配项
func (re *Regexp) Match(b []byte) bool

func main() {
    b := []byte(`Hello World`)
    reg := regexp.MustCompile(`Hello\w+`)
    fmt.Println(reg.Match(b))
    // false
    reg = regexp.MustCompile(`Hello[\w\s]+`)
    fmt.Println(reg.Match(b))
    // true
}

------------------------------------------------------------

// 判断在 r 中能否找到匹配项
func (re *Regexp) MatchReader(r io.RuneReader) bool

func main() {
    r := bytes.NewReader([]byte(`Hello World`))
    reg := regexp.MustCompile(`Hello\w+`)
    fmt.Println(reg.MatchReader(r))
    // false
    r.Seek(0, 0)
    reg = regexp.MustCompile(`Hello[\w\s]+`)
    fmt.Println(reg.MatchReader(r))
    // true
}

------------------------------------------------------------

// 判断在 s 中能否找到匹配项
func (re *Regexp) MatchString(s string) bool

func main() {
    s := `Hello World`
    reg := regexp.MustCompile(`Hello\w+`)
    fmt.Println(reg.MatchString(s))
    // false
    reg = regexp.MustCompile(`Hello[\w\s]+`)
    fmt.Println(reg.MatchString(s))
    // true
}

------------------------------------------------------------

// 统计正则表达式中的分组个数（不包括“非捕获的分组”）
func (re *Regexp) NumSubexp() int

func main() {
    reg := regexp.MustCompile(`(?U)(?:Hello)(\s+)(\w+)`)
    fmt.Println(reg.NumSubexp())
    // 2
}

------------------------------------------------------------

// 在 src 中搜索匹配项，并替换为 repl 指定的内容
// 全部替换，并返回替换后的结果
func (re *Regexp) ReplaceAll(src, repl []byte) []byte

func main() {
    b := []byte("Hello World, 123 Go!")
    reg := regexp.MustCompile(`(Hell|G)o`)
    rep := []byte("${1}ooo")
    fmt.Printf("%q\n", reg.ReplaceAll(b, rep))
    // "Hellooo World, 123 Gooo!"
}

------------------------------------------------------------

// 在 src 中搜索匹配项，并替换为 repl 指定的内容
// 全部替换，并返回替换后的结果
func (re *Regexp) ReplaceAllString(src, repl string) string

func main() {
    s := "Hello World, 123 Go!"
    reg := regexp.MustCompile(`(Hell|G)o`)
    rep := "${1}ooo"
    fmt.Printf("%q\n", reg.ReplaceAllString(s, rep))
    // "Hellooo World, 123 Gooo!"
}

------------------------------------------------------------

// 在 src 中搜索匹配项，并替换为 repl 指定的内容
// 如果 repl 中有“分组引用符”（$1、$name），则将“分组引用符”当普通字符处理
// 全部替换，并返回替换后的结果
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte

func main() {
    b := []byte("Hello World, 123 Go!")
    reg := regexp.MustCompile(`(Hell|G)o`)
    rep := []byte("${1}ooo")
    fmt.Printf("%q\n", reg.ReplaceAllLiteral(b, rep))
    // "${1}ooo World, 123 ${1}ooo!"
}

------------------------------------------------------------

// 在 src 中搜索匹配项，并替换为 repl 指定的内容
// 如果 repl 中有“分组引用符”（$1、$name），则将“分组引用符”当普通字符处理
// 全部替换，并返回替换后的结果
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string

func main() {
    s := "Hello World, 123 Go!"
    reg := regexp.MustCompile(`(Hell|G)o`)
    rep := "${1}ooo"
    fmt.Printf("%q\n", reg.ReplaceAllLiteralString(s, rep))
    // "${1}ooo World, 123 ${1}ooo!"
}

------------------------------------------------------------

// 在 src 中搜索匹配项，然后将匹配的内容经过 repl 处理后，替换 src 中的匹配项
// 如果 repl 的返回值中有“分组引用符”（$1、$name），则将“分组引用符”当普通字符处理
// 全部替换，并返回替换后的结果
func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte

func main() {
    s := []byte("Hello World!")
    reg := regexp.MustCompile("(H)ello")
    rep := []byte("$0$1")
    fmt.Printf("%s\n", reg.ReplaceAll(s, rep))
    // HelloH World!

    fmt.Printf("%s\n", reg.ReplaceAllFunc(s,
        func(b []byte) []byte {
            rst := []byte{}
            rst = append(rst, b...)
            rst = append(rst, "$1"...)
            return rst
        }))
    // Hello$1 World!
}
k
------------------------------------------------------------

// 在 src 中搜索匹配项，然后将匹配的内容经过 repl 处理后，替换 src 中的匹配项
// 如果 repl 的返回值中有“分组引用符”（$1、$name），则将“分组引用符”当普通字符处理
// 全部替换，并返回替换后的结果
func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string

func main() {
    s := "Hello World!"
    reg := regexp.MustCompile("(H)ello")
    rep := "$0$1"
    fmt.Printf("%s\n", reg.ReplaceAllString(s, rep))
    // HelloH World!
    fmt.Printf("%s\n", reg.ReplaceAllStringFunc(s,
        func(b string) string {
            return b + "$1"
        }))
    // Hello$1 World!
}

------------------------------------------------------------

// 在 s 中搜索匹配项，并以匹配项为分割符，将 s 分割成多个子串
// 最多分割出 n 个子串，第 n 个子串不再进行分割
// 如果 n < 0，则分割所有子串
// 返回分割后的子串列表
func (re *Regexp) Split(s string, n int) []string

func main() {
    s := "Hello World\tHello\nGolang"
    reg := regexp.MustCompile(`\s`)
    fmt.Printf("%q\n", reg.Split(s, -1))
    // ["Hello" "World" "Hello" "Golang"]
}

------------------------------------------------------------

// 返回 re 中的“正则表达式”字符串
func (re *Regexp) String() string

func main() {
    re := regexp.MustCompile("Hello.*$")
    fmt.Printf("%s\n", re.String())
    // Hello.*$
}

------------------------------------------------------------

// 返回 re 中的分组名称列表，未命名的分组返回空字符串
// 返回值[0] 为整个正则表达式的名称
// 返回值[1] 是分组 1 的名称
// 返回值[2] 是分组 2 的名称
// ……
func (re *Regexp) SubexpNames() []string

func main() {
    re := regexp.MustCompile("(?PHello) (World)")
    fmt.Printf("%q\n", re.SubexpNames())
    // ["" "Name1" ""]
}
```
