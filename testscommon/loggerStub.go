package testscommon

import logger "github.com/subrahamanyam341/andes-logger-123"

// LoggerStub -
type LoggerStub struct {
	TraceCalled      func(message string, args ...interface{})
	DebugCalled      func(message string, args ...interface{})
	InfoCalled       func(message string, args ...interface{})
	WarnCalled       func(message string, args ...interface{})
	ErrorCalled      func(message string, args ...interface{})
	LogIfErrorCalled func(err error, args ...interface{})
	GetLevelCalled   func() logger.LogLevel
}

// Trace -
func (stub *LoggerStub) Trace(message string, args ...interface{}) {
	if stub.TraceCalled != nil {
		stub.TraceCalled(message, args)
	}
}

// Debug -
func (stub *LoggerStub) Debug(message string, args ...interface{}) {
	if stub.DebugCalled != nil {
		stub.DebugCalled(message, args)
	}
}

// Info -
func (stub *LoggerStub) Info(message string, args ...interface{}) {
	if stub.InfoCalled != nil {
		stub.InfoCalled(message, args)
	}
}

// Warn -
func (stub *LoggerStub) Warn(message string, args ...interface{}) {
	if stub.WarnCalled != nil {
		stub.WarnCalled(message, args)
	}
}

// Error -
func (stub *LoggerStub) Error(message string, args ...interface{}) {
	if stub.ErrorCalled != nil {
		stub.ErrorCalled(message, args)
	}
}

// LogIfError -
func (stub *LoggerStub) LogIfError(err error, args ...interface{}) {
	if stub.LogIfErrorCalled != nil {
		stub.LogIfErrorCalled(err, args)
	}
}

// GetLevel -
func (stub *LoggerStub) GetLevel() logger.LogLevel {
	if stub.GetLevelCalled != nil {
		return stub.GetLevelCalled()
	}

	return logger.LogTrace
}

// IsInterfaceNil -
func (stub *LoggerStub) IsInterfaceNil() bool {
	return false
}
