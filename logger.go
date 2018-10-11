package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/logutils"
)

// Наиболее простой логгер, на основе logutils

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

// Params для логгера
type Params struct {
	Writer   io.Writer
	Levels   []string
	MinLevel string
	Package  string
}

type logger struct {
	packageName string
}

// Init - инициализация логгера
func Init(params Params) {
	levels := []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	minLevel := logutils.LogLevel("INFO")

	var writer io.Writer = os.Stderr

	// Установка уровней логгирования (по умолчанию ["DEBUG", "INFO", "WARN", "ERROR", "FATAL"])
	if params.Levels != nil {
		levels = []logutils.LogLevel{}

		for _, level := range params.Levels {
			levels = append(levels, logutils.LogLevel(strings.ToUpper(level)))
		}
	}

	// Установка минимального уровня логгирования (по умолчанию INFO)
	if params.MinLevel != "" {
		minLevel = logutils.LogLevel(strings.ToUpper(params.MinLevel))
	}

	// Установка writer для логов (по умолчанию os.Stderr)
	if params.Writer != nil {
		writer = params.Writer
	}

	filter := &logutils.LevelFilter{
		Levels:   levels,
		MinLevel: minLevel,
		Writer:   writer,
	}

	log.SetOutput(filter)
}

// NewLogger - возвращает объект логгера с заданными параметрами
func NewLogger(packageName string) Logger {

	return logger{packageName: packageName + ":"}
}

func (l logger) Debug(msg string) {
	log.Println("[DEBUG]", l.packageName, msg)
}

func (l logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l logger) Info(msg string) {
	log.Println("[INFO]", l.packageName, msg)
}

func (l logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l logger) Warn(msg string) {
	log.Println("[WARNING]", l.packageName, msg)
}

func (l logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l logger) Error(msg string) {
	log.Println("[ERROR]", l.packageName, msg)
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l logger) Fatal(msg string) {
	log.Println("[FATAL]", l.packageName, msg)
}

func (l logger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}
