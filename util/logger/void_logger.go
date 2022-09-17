package logger

import (
	"context"
	"time"

	lg "gorm.io/gorm/logger"
)

type VoidLogger struct {
	lg.Interface
}

func (v *VoidLogger) LogMode(lg.LogLevel) lg.Interface {
	return v
}
func (v *VoidLogger) Info(context.Context, string, ...interface{}) {

}
func (v *VoidLogger) Warn(context.Context, string, ...interface{}) {

}
func (v *VoidLogger) Error(context.Context, string, ...interface{}) {

}
func (v *VoidLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
}
