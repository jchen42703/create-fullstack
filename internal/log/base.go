package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Sets the logger level based on the environment variable.
// Completely ignore if no env variable set.
func SetLoggerLevel(envVarName string, atom zap.AtomicLevel) error {
	lvl := os.Getenv(envVarName)
	// Expects the string representation of log levels
	// I.e. debug, warn, info, error, panic, fatal
	logLvl, err := zapcore.ParseLevel(lvl)
	// Just exit early with error
	if err != nil {
		return err
	}

	atom.SetLevel(logLvl)
	return nil
}

// Creates a logger instance, the logger file, and sets the logger level based on the
// CREATE_FULLSTACK_LOG_LVL environment.
// The logger logs to both a logger file (DEBUG level) and stdout (Info).
func CreateLogger(logFilePath string) (*zap.Logger, error) {
	// Create logger
	atom := zap.NewAtomicLevel()

	// To keep the example deterministic, disable timestamps in the output.
	encoderCfg := zap.NewProductionEncoderConfig()

	// Creates log file if does not already exist
	f, err := os.Create(logFilePath)
	if err != nil {
		return nil, fmt.Errorf("CreateLogger: %s", err.Error())
	}

	// File + JSON logger
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), zapcore.AddSync(f), zap.DebugLevel),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			atom,
		))

	logger := zap.New(core)
	err = SetLoggerLevel("CFS_LOG_LVL", atom)
	if err != nil {
		logger.Debug("using default logger level: WARN")
	}

	return logger, nil
}
