package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type debugFlag bool

func (d debugFlag) BeforeApply() error {
	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	return nil
}

func (d debugFlag) AfterApply() error {
	if d {
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}
	return nil
}

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
