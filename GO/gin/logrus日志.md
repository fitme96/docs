```go
package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	//设定字段
	// log := logrus.WithField("app", "study")

	// log.Warnln("test")

	// logrus.SetLevel(logrus.InfoLevel)
	// logrus.Warnln("警告")
	// logrus.Infoln("信息")

	// fmt.Println(logrus.GetLevel())

	//自定义格式

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("警告")
	logrus.Debugln("警告")
	logrus.Infoln("警告")
	logrus.Errorln("警告")

}

```

颜色

```go
package main

import "fmt"

func main() {
	fmt.Println("\033[31m 测试\033[0m") //红色
	fmt.Println("\033[30m 测试\033[0m") //黑色
}

```


写入文件

```go
package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	file, _ := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	logrus.SetOutput(file)
	logrus.SetLevel(logrus.DebugLevel)

	logrus.SetOutput(io.MultiWriter(file, os.Stdout))

	logrus.Info("信息")
	logrus.Debugln("debug")
}


```


hook
```go
package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Myhook struct {
}

func (hook Myhook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook Myhook) Fire(entry *logrus.Entry) error {
	file, _ := os.OpenFile("err.og", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	line, _ := entry.String()
	file.Write([]byte(line))
	return nil
}

func main() {
	logrus.AddHook(&Myhook{})
	logrus.Warn("warn")
	logrus.Error("error")

}

```