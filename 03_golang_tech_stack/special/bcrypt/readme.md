<!--
 * @Author: ChenLong longchen2008@126.com
 * @Date: 2022-06-04 23:43:48
 * @LastEditors: ChenLong longchen2008@126.com
 * @LastEditTime: 2022-06-04 23:50:41
 * @FilePath: \study\03_golang_tech_stack\special\bcrypt\readme.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# Golang加密和解密应用

## 知识图表

- mysql
- gorm
- gin
- bcrypt
  
## 实现步骤

### 安装库

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get github.com/gin-gonic/gin
```

### 导包

```bash
import (
    "fmt"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)
```

### 加密解密算法

```bash
// 生成密码
func GenPwd(pwd string) ([]byte, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost) //加密处理
    return hash, err
}

// 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
    if err != nil {
        return false
    } else {
        return true
    }
}
```

### 首页

```bash
func Index(c *gin.Context) {
    c.HTML(200, "index.html", nil)
}
```

### 用户模型gorm

```bash
type User struct {
    gorm.Model
    Username string `json:"username"`
    Password string `json:"password"`
}

var db *gorm.DB

func init() {
    dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&    parseTime=True&loc=Local"
    db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    // 迁移 schema
    db.AutoMigrate(&User{})

}
```

### 用户注册

```bash
func GoRegister(c *gin.Context) {
    c.HTML(200, "register.html", nil)
}

func Register(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    b, _ := GenPwd(password)

    user := User{
        Username: username,
        Password: string(b),
    }
    db.Create(&user)
    c.Redirect(301, "/")
}
```

### 用户登录

```bash
func GoLogin(c *gin.Context) {
    c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    fmt.Printf("表单密码: %v\n", password)

    // 查询
    var u User
    db.Where("username=?", username).First(&u)

    if u.Username == "" {
        c.HTML(200, "login.html", "用户名不存在！")
        fmt.Println("用户名不存在！")
    } else {
        // 比较
        if ComparePwd(u.Password, password) {
            fmt.Println("登录成功")
            c.Redirect(301, "/")
        } else {
            fmt.Println("密码错误")
            c.HTML(200, "login.html", "密码错误")
        }
    }
}
```

### 路由

```bash
func main() {

    e := gin.Default()
    e.LoadHTMLGlob("*.html")

    e.GET("/login", GoLogin)
    e.POST("/login", Login)

    e.GET("/", Index)
    e.POST("/register", Register)
    e.GET("/register", GoRegister)
    e.Run()

}
```

### html

#### html首页

```bash
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>golang技术栈-老郭-用户登录密码加密与解密</title>
</head>
<body>

    <h1>golang技术栈-老郭-用户登录密码加密与解密</h1>

    <a href="/login">登录</a><br>
    <a href="/register">注册</a><br>

    
</body>
```

#### html注册

```bash
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>golang技术栈-老郭-用户注册</title>
</head>
<body>

    <h1>golang技术栈-老郭-用户注册</h1>

    <form action="/register" method="post">
        Username: <input type="text" name="username"><br>
        Password：<input type="password" name="password"><br>
        <input type="submit" value="注册">
    </form>
    
</body>
</html>
```

#### html登录

```bash
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>golang技术栈-老郭-用户登录</title>
</head>
<body>

    <h1>golang技术栈-老郭-用户登录</h1>

    <form action="/login" method="post">
        Username: <input type="text" name="username"><br>
        Password：<input type="password" name="password"><br>
        <input type="submit" value="登录">
    </form>
    
</body>
</html>
```
