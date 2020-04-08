# proxygool - Go可扩展代理池

> 默认以Redis作为存储，可基于Store接口扩展

![](https://img.shields.io/badge/language-Go-00B1D6.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/cmatrixprobe/proxygool)](https://goreportcard.com/report/github.com/cmatrixprobe/proxygool)

## 运行方式
1. Docker
2. 常规部署

### Docker

首先需要确保安装了如下环境：

* Docker
* Docker-Compose

然后在项目根目录执行：

```shell script
docker-compose up -d
```

### 常规部署

首先配置相应的Go及Redis环境

然后在项目根目录执行：

```shell script
go build
./proxygool
```

或者：

```shell script
go run main.go
```

## 获取代理

* Redis client
* Web url

### Redis client

```shell script
redis-cli
```

* 查看爬取的所有代理

```shell script
SMEMBERS proxypool
```

* 查看代理的详细信息

```shell script
HGETALL proxyinfo
```

### Web url

* 随机获取一条代理

```shell script
curl localhost:8888
```

* 随机获取一条https代理

```shell script
curl localhost:8888/https
```

程序将定时检测心跳并删除无效代理

## 配置项

* 常规部署需要将docker置为false
* pages控制不同代理网站的爬取页数
* fetch.proxy控制是否用代理池爬取

```yaml
docker: true
logger:
  filename: proxygool.log
  level: 4
server:
  host: 0.0.0.0
  port: 8888
redis:
  network: tcp
  host: 127.0.0.1
  port: 6379
  password:
  MaxIdle: 100
  MaxActive: 0
  IdleTimeout: 5m
  testFrequency: 1m
  Wait: true
  proxyPool: proxypool
  proxyInfo: proxyinfo
fetch:
  proxy: true
xicidaili:
  pages: 5
kuaidaili:
  pages: 5
```

## 代理站点

1. [齐云代理](https://www.7yip.cn/)
2. [66代理](http://www.66ip.cn/)
3. [89免费代理](http://www.89ip.cn/)
4. [云代理](http://www.ip3366.net/)
5. [快代理](https://www.kuaidaili.com/)
6. [ProxyListPlus](https://list.proxylistplus.com/)
7. [西刺代理](https://www.xicidaili.com/)

* 除了以上已实现的代理，可以在一分钟内轻松扩展代理接口

```go
func XXX() *model.Request {
    req := model.NewRequest()
    req.WebName = "xxx"
    req.WebURL = "http://www.xxx.cn/index_"
    req.TrRegular = ".table tbody tr"
    req.Pages = viper.GetInt("xxx.pages")
    req.HostIndex = 0
    req.PortIndex = 1
    req.ProtIndex = 3
    req.Protocol = func(s string) string {
	if s == "no" {
	    return "http"
	}
	return "https"
    }
    return req
    req.Trim = true
}

```

* 在spider/site中新增函数并配置到spider/run.go

```go
requests = []*model.Request{
    site.Xici(),
    site.Kuai(),
    site.IP3366(),
    site.Qiyun(),
    //site.PLP(),
    //site.PLPSSL(),
    site.IP66(),
    site.IP89(),
    site.XXX(),
}
```

## 扩展存储方式

1. 实现store.Store接口
2. 调用store.SetCustomStore(s Store)