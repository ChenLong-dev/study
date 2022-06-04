<!--
 * @Author: ChenLong longchen2008@126.com
 * @Date: 2022-06-05 00:07:53
 * @LastEditors: ChenLong longchen2008@126.com
 * @LastEditTime: 2022-06-05 00:19:03
 * @FilePath: \study\03_golang_tech_stack\special\viper\readme.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# Golang配置管理库viper

viper是一个golang配置管理库，很多项目都使用viper来构建，例如：docker、Hugo等等

## 安装viper

```bash
go get github.com/spf13/viper
```

## viper支持的配置很多

- 从 JSON、TOML、YAML、HCL、envfile 和 Java 属性配置文件读取
- 实时监视和重新阅读配置文件
- 从环境变量中读取
- 从远程配置系统（etcd 或 Consul）读取，并观察变化
- 从命令行标志读取
- 从缓冲区读取

## 读取ini配置文件

### ini配置文件

```bash
[db]
username=admin
password=123
driver=mysql
port=3306
```

### 读取

```bash
func ReadIni() {
    v := viper.New()
    v.AddConfigPath("./conf") // 路径
    v.SetConfigName("config") // 名称
    v.SetConfigType("ini")    // 类型

    err := v.ReadInConfig() // 读配置
    if err != nil {
        panic(err)
    }

    s := v.GetString("db.username")
    fmt.Printf("s: %v\n", s)
}
```

## 读yaml配置文件

### yaml配置文件

```bash
db: 
  username: admin
  password: 123
  port: 3306
  driver: mysql

```

### 读取

```bash
func ReadYml() {
    v := viper.New()
    v.AddConfigPath("./conf") // 路径
    v.SetConfigName("config") // 名称
    v.SetConfigType("yaml")   // 类型

    err := v.ReadInConfig() // 读配置
    if err != nil {
        panic(err)
    }

    s := v.GetString("db.username")
    fmt.Printf("s: %v\n", s)
}
```
