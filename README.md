## 说明
基于 zap 封装的日志库，只封装了一些常用的操作。

## 使用
包级函数直接调用，支持 `Debug`、`Info`、`Warn`、`Error`、`Fatal` 五种级别日志。
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

默认只会显示 `Info` 及以上级别的日志，若要显示 `Debug` 或其他级别的日志，请调用 `log.SetOptions(log.WithLevel(log.DebugLevel))`，例如：
```go
package main

import (
	"github.com/fengh0409/log"
)

func main() {
	defer log.Sync()
	log.SetOptions(log.WithLevel(log.DebugLevel))
	// 以上 log.DebugLevel 是 int 类型，若是命令行参数传进来的字符串，如 info，使用以下方式：
	// log.SetOptions(log.WithLevelString("debug"))

	log.Debug("this is debug level log")
	log.Info("this is info level log")
}
```

输出：
```
2022-04-29T22:04:19.204+0800	DEBUG	mytest/main.go:11	this is debug level log
2022-04-29T22:04:19.205+0800	INFO	mytest/main.go:12	this is info level log
```


## 日志写入文件
默认情况下，日志写入到标准错误输出，若要写入文件，请调用 `log.SetOptions(log.WithFileWriter())`

```go
package main

import (
	"github.com/fengh0409/log"
)

func main() {
	defer log.Sync()
	log.SetOptions(log.WithFileWriter(log.WithFilename("/tmp/mytest.log")))

	log.Info("show some message")
	log.Error("show some error message")
}
```

查看 `/tmp/mytest.log` 文件内容：
```
2022-05-05T21:31:11.680+0800	INFO	mytest/main.go:11	show some message
2022-05-05T21:31:11.681+0800	ERROR	mytest/main.go:12	show some error message
```

日志文件支持自动切割，`log.WithFileWriter()` 可以传入以下函数进行配置：
* log.WithFilename("/tmp/log/lumberjack.log") 	设置日志文件名，默认写入到 `/tmp/log/lumberjack.log`
* log.WithMaxSize(200) 						  	设置日志文件最大容量，单位 MB，默认 200MB
* log.WithMaxAge(7) 							设置日志文件保留最长时间，单位 天，默认 7 天
* log.WithBackups(10) 							设置日志文件最大个数，默认 10 个
* log.WithCompress(true) 						设置是否压缩归档的日志文件，默认是


## 结构化日志
输出 json 格式的结构化日志，输出到标准输出，打印 `Info` 及以上级别日志。
```go
package main

import (
	"os"

	"github.com/fengh0409/log"
)

func main() {
	logger := log.New(
		log.WithWriter(os.Stdout),
		log.WithEncoding("json"),
	).Build()
	defer logger.Sync()

	logger.Info("show some message", log.String("hello", "world"))
	logger.Error("show some error message", log.Int("code", 404))
}
```

输出：
```
{"level":"info","ts":1651757350.029574,"caller":"mytest/main.go:16","msg":"show some message","hello":"world"}
{"level":"error","ts":1651757350.0296729,"caller":"mytest/main.go:17","msg":"show some error message","code":404}
```
