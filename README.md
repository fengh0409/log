## 说明
基于 zap 封装的日志库，只封装了一些常用的操作。

## 使用
包级函数直接调用，支持 `Debug`、`Info`、`Warn`、`Error`、`Fatal` 五种级别日志.
```go
package main

import (
	"github.com/fengh0409/log"
)

func main() {
	defer log.Sync()
	log.Debug("this is debug level log")
	log.Info("this is info level log", " more args")
	log.Warn("this is warn level log")
	log.Error("this is error level log")
	//log.Fatal("this is fatal level log") // will exit
	log.Debugf("this is %s level log", "debugf")
	log.Infof("this is %s level log", "infof")
	log.Warnf("this is %s level log", "warnf")
	log.Errorf("this is %s level log", "errorf")
	//log.Fatalf("this is %s level log", "fatalf") //will exit
}
```

输出：
```
2022-04-29T22:02:42.744+0800	INFO	mytest/main.go:10	this is info level log more args
2022-04-29T22:02:42.744+0800	WARN	mytest/main.go:11	this is warn level log
2022-04-29T22:02:42.744+0800	ERROR	mytest/main.go:12	this is error level log
2022-04-29T22:02:42.744+0800	INFO	mytest/main.go:15	this is infof level log
2022-04-29T22:02:42.744+0800	WARN	mytest/main.go:16	this is warnf level log
2022-04-29T22:02:42.744+0800	ERROR	mytest/main.go:17	this is errorf level log
```

默认只会显示 `Info` 及以上级别的日志，若要显示 `Debug` 或其他级别的日志，请调用 `1og.SetLogLeve1(1og.DebugLeve1)`，例如：
```go
package main

import (
	"github.com/fengh0409/log"
)

func main() {
	defer log.Sync()
	log.SetLogLevel(log.DebugLevel)

	log.Debug("this is debug level log")
	log.Info("this is info level log")
}

```

输出：
```
2022-04-29T22:04:19.204+0800	DEBUG	mytest/main.go:11	this is debug level log
2022-04-29T22:04:19.205+0800	INFO	mytest/main.go:12	this is info level log
```

## 定制化
输出 json 格式的结构化日志，输出到标准输出，打印 `Info` 及以上级别日志。
```go
package main

import (
	"os"

	"github.com/fengh0409/log"
)

func main() {
	logger := log.New(os.Stderr, log.InfoLevel, log.JSONEncoder)
	defer logger.Sync()

	logger.Info("show some message", log.String("hello", "world"))
	logger.Error("show some error message", log.Int("code", 404))
}

```

输出：
```
{"level":"info","ts":1651241199.771987,"caller":"mytest/main.go:13","msg":"show some message","hello":"world"}
{"level":"error","ts":1651241199.772078,"caller":"mytest/main.go:14","msg":"show some error message","code":404}

```
