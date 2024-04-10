package zlog

import (
	"fmt"
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogger *zap.Logger
)

func init() {
	execName, _ := os.Executable()
	appName := path.Base(execName)
	if appName == "main" {
		appName = fmt.Sprintf("go-run-%v", time.Now().Format("20060102T150405"))
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dir := path.Join(homeDir, "go-app-logs")
	fmt.Println("dir = ", dir)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		dirErr := os.MkdirAll(dir, 0700) // Create your file
		if dirErr != nil {
			panic(dirErr)
		}
	}

	fname := path.Join(dir, appName+".log")
	fmt.Println("log filename = ", fname)
	InitZapLogger(fname)
}

func InitZapLogger(fileName string) {
	// Create a file to write logs to
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	// Configure encoder with custom time format
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	// Create file writer
	fileWriter := zapcore.AddSync(file)

	// Create console writer
	consoleWriter := zapcore.Lock(os.Stdout)

	// Set encoder for file and console writers
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Configure core with multiple outputs
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWriter, zap.InfoLevel),
		zapcore.NewCore(consoleEncoder, consoleWriter, zap.InfoLevel),
	)

	// Create logger
	defaultLogger = zap.New(core)

	// Example logging
	Infof("InitZapLogger: file:%v", fileName)
}

func Infof(format string, args ...any) {
	defaultLogger.Info(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...any) {
	defaultLogger.Warn(fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...any) {
	defaultLogger.Error(fmt.Sprintf(format, args...))
}

func Debugf(format string, args ...any) {
	defaultLogger.Debug(fmt.Sprintf(format, args...))
}
