package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logFormat string

func (l logFormat) BeforeResolve() error {
	switch l {
	case "console":
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	return nil
}

func (l logFormat) AfterApply() error {
	switch l {
	case "console":
	case "json":
		zapConfig.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		zapConfig.Encoding = "json"
	}
	return nil
}

type timeFormat string

func (l timeFormat) AfterApply() error {
	switch l {
	case "":
	case "local-time":

	}
	return nil
}

type logLevel string

func (l logLevel) IsDebug() bool {
	switch l {
	case "-2":
		return true
	}
	return false
}

func (l logLevel) BeforeApply() error {
	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	return nil
}

func (l logLevel) AfterApply() error {
	var level zapcore.Level
	switch l {
	case "-2", "-1", "debug", "DEBUG":
		level = zapcore.DebugLevel
	case "0", "info", "INFO":
		level = zapcore.InfoLevel
	case "1", "warn", "WARN":
		level = zapcore.WarnLevel
	case "2", "error", "ERROR":
		level = zapcore.ErrorLevel
	case "3", "panic", "PANIC":
		level = zapcore.PanicLevel
	case "4", "fatal", "FATAL":
		level = zapcore.PanicLevel
	default:
		fmt.Printf("Invalid level %s. Defaulting to Info", l)
		level = zapcore.InfoLevel
	}
	zapConfig.Level = zap.NewAtomicLevelAt(level)
	return nil
}
