package consumer

//
//import (
//	"github.com/sirupsen/logrus"
//)
//
//// A Logger is a minimal interface to as a adaptor for external logging library to consumer
//type Logger interface {
//	Log(...interface{})
//	Info(...interface{})
//	Debug(...interface{})
//	SetLevel(...interface{})
//}
//
//// noopLogger implements logger interface with discard
//type noopLogger struct {
//	logger *logrus.Logger
//}
//
//// Log using stdlib logger. See log.Println.
//func (l noopLogger) Debug(args ...interface{}) {
//	l.logger.Debug(args...)
//}
//
//func (l noopLogger) Info(args ...interface{}) {
//	l.logger.Info(args...)
//}
//
//func (l noopLogger) Log(args ...interface{}) {
//	l.logger.Debug(args...)
//}
