# Библиотека для организации логгирования в приложениях на Go


## Установка

Добавить в **Gopkg.toml** проекта строки

```toml
[[constraint]]
  name = "github.com/ekonomizer/golog"
  source = "ssh://git@github.com:ekonomizer/golog.git"
  version = "0.1.3"
```

## Описание

Предельно простой логгер реализованный поверх logutils (https://github.com/hashicorp/logutils), позволяющий писать в `os.Stderr`, `syslog` и вообще куда угодно, если оно реализует интерфейс `io.Writer`.

Параметры логгера задаются один раз вызовом метода `Init` и передачи в него структуры `logger.Params`:

```go
// Params для логгера
type Params struct {
    // Куда писать логи
    Writer   io.Writer
    // Уровни логгирования
    Levels   []string
    // Минимальный уровень логов для отображения
    MinLevel string
}
```

Интерфейс:

```go
// Logger - интерфейс логгера
type Logger interface {
    Debug(msg string)
    Debugf(msg string, args ...interface{})

    Info(msg string)
    Infof(format string, args ...interface{})

    Warn(msg string)
    Warnf(format string, args ...interface{})

    Error(msg string)
    Errorf(format string, args ...interface{})

    Fatal(msg string)
    Fatalf(format string, args ...interface{})
}
```

### Формат логов

```bash
# <data> <time> <severity> <module>: <message>

2018/05/08 09:28:49 [INFO] service: Starting...
```

## Примеры использования

### stderr

```go
import (
    "fmt"
    "os"

    "github.com/ekonomizer/golog"
)

type Service struct {
    log logger.Logger
}

func NewService() *Service {
    logger.Init(logger.Params{
        Writer: os.Stderr,
    })

    return &Service{
        log: logger.NewLogger("service"),
    }
}

func (s *Service) SayHi(name string) {
    s.log.Infof("Saying hi to %s", name)
    fmt.Printf("Hi, %s", name)
}

func Example() {
    service := NewService()
    service.SayHi("Alex")

    // Вывод:
    // 2018/03/23 16:03:18 [INFO] service: Saying hi to Alex
    // Hi, Alex
}
```

## GELF (graylog)

```go
import "gopkg.in/Graylog2/go-gelf.v1/gelf"

gelfWriter, err := gelf.NewWriter("localhost:12201")
if err != nil {
    return errors.Wrap(err, "unable to create gelf writer")
}
```

### syslog

```go
import (
    "log/syslog"
    "github.com/pkg/errors"

    "github.com/ekonomizer/golog"
)

logWriter, err := syslog.New(syslog.LOG_NOTICE, "service_name")

if err != nil {
    return errors.Wrap(err, "unable to create syslog writer")
}

logger.Init(logger.Params{
    Writer: logWriter,
})

log := logger.NewLogger("service")

log.Info("I'm going to syslog")
```

### rsyslog (udp)

```go
logWriter, err := syslog.Dial("udp", "rsyslog:514", syslog.LOG_NOTICE, "service_name")

if err != nil {
    return errors.Wrap(err, "unable to create syslog writer")
}
```

### rsyslog (tcp)

```go
logWriter, err := syslog.Dial("tcp", "rsyslog:10514", syslog.LOG_NOTICE, "service_name")

if err != nil {
    return errors.Wrap(err, "unable to create syslog writer")
}
```
