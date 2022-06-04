# Golang访问权限控制框架casbin

## casbin简介

Casbin 是一个强大的、高效的开源访问控制框架，其权限管理机制支持多种访问控制模型。支持的语言也很多，例如：go、java、node.js、python等等.

### Casbin 是什么？

#### Casbin 可以

1、支持自定义请求的格式，默认的请求格式为{subject, object, action}。

2、具有访问控制模型model和策略policy两个核心概念。

3、支持RBAC中的多层角色继承，不止主体可以有角色，资源也可以具有角色。

4、支持内置的超级用户 例如：root 或 administrator。超级用户可以执行任何操作而无需显式的权限声明。

5、支持多种内置的操作符，如 keyMatch，方便对路径式的资源进行管理，如 /foo/bar 可以映射到 /foo*

#### Casbin 不能

1、身份认证 authentication（即验证用户的用户名和密码），Casbin 只负责访问控制。应该有其他专门的组件负责身份认证，然后由 Casbin 进行访问控制，二者是相互配合的关系。

2、管理用户列表或角色列表。 Casbin 认为由项目自身来管理用户、角色列表更为合适， 用户通常有他们的密码，但是 Casbin 的设计思想并不是把它作为一个存储密码的容器。 而是存储RBAC方案中用户和角色之间的映射关系。

## 实例演示

### 安装包

```bash
go get github.com/casbin/casbin
go get github.com/casbin/gorm-adapter
go get github.com/go-sql-driver/mysql
go get github.com/jinzhu/gorm
```

### 创建model配置文件

```bash
# Request定义
[request_definition]
r = sub, obj, act

# 策略定义
[policy_definition]
p = sub, obj, act

# 角色定义
[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

# 匹配器定义
[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
```

### 权限控制

```bash
package main

import (
    "fmt"

    "github.com/casbin/casbin"
    gormadapter "github.com/casbin/gorm-adapter"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    // 使用 MySQL 数据库初始化一个 gorm 适配器
    a := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/  test_db", true)

    e := casbin.NewEnforcer("rbac_model.conf", a)

    sub := "ghz"   // 想要访问资源的用户。
    obj := "data1" // 将被访问的资源。
    act := "read"  // 用户对资源执行的操作。

    // e.AddPolicy(sub, obj, act)

    // e.ClearPolicy()

    ok := e.Enforce(sub, obj, act)

    if ok {
        // 允许ghz读取data1
        fmt.Println("yes")
    } else {
        // 拒绝请求，抛出异常
        fmt.Println("no")
    }
}
```

### 测试sql

```bash
INSERT INTO  casbin_rule(`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'ghz', 'data1', 'read', NULL, NULL, NULL);
```
