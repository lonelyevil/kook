package plog

import (
	"github.com/lonelyevil/khl"
	"github.com/phuslu/log"
	"net"
	"time"
)

type LoggerAdapter struct {
	l *log.Logger
}

func NewLogger(l *log.Logger) *LoggerAdapter {
	return &LoggerAdapter{l: l}
}

func (l LoggerAdapter) Trace() khl.Entry {
	return &EntryAdapter{l.l.Trace()}
}

func (l LoggerAdapter) Debug() khl.Entry {
	return &EntryAdapter{l.l.Debug()}
}

func (l LoggerAdapter) Info() khl.Entry {
	return &EntryAdapter{l.l.Info()}
}

func (l LoggerAdapter) Warn() khl.Entry {
	return &EntryAdapter{l.l.Warn()}
}

func (l LoggerAdapter) Error() khl.Entry {
	return &EntryAdapter{l.l.Error()}
}

func (l LoggerAdapter) Fatal() khl.Entry {
	return &EntryAdapter{l.l.Fatal()}
}

type EntryAdapter struct {
	e *log.Entry
}

func (e *EntryAdapter) Bool(key string, b bool) khl.Entry {
	e.e = e.e.Bool(key, b)
	return e
}

func (e *EntryAdapter) Bytes(key string, val []byte) khl.Entry {
	e.e = e.e.Bytes(key, val)
	return e
}

func (e *EntryAdapter) Caller(depth int) khl.Entry {
	e.e = e.e.Caller(depth, false)
	return e
}

func (e *EntryAdapter) Dur(key string, d time.Duration) khl.Entry {
	e.e = e.e.Dur(key, d)
	return e
}

func (e *EntryAdapter) Err(key string, err error) khl.Entry {
	e.e = e.e.Err(err)
	return e
}

func (e *EntryAdapter) Float64(key string, f float64) khl.Entry {
	e.e = e.e.Float64(key, f)
	return e
}

func (e *EntryAdapter) IPAddr(key string, ip net.IP) khl.Entry {
	e.e = e.e.IPAddr(key, ip)
	return e
}

func (e *EntryAdapter) Int(key string, i int) khl.Entry {
	e.e = e.e.Int(key, i)
	return e
}

func (e *EntryAdapter) Int64(key string, i int64) khl.Entry {
	e.e = e.e.Int64(key, i)
	return e
}

func (e *EntryAdapter) Interface(key string, i interface{}) khl.Entry {
	e.e = e.e.Interface(key, i)
	return e
}

func (e *EntryAdapter) Msg(msg string) {
	e.e.Msg(msg)
}

func (e *EntryAdapter) Msgf(f string, i ...interface{}) {
	e.e.Msgf(f, i)
}

func (e *EntryAdapter) Str(key string, s string) khl.Entry {
	e.e = e.e.Str(key, s)
	return e
}

func (e *EntryAdapter) Strs(key string, s []string) khl.Entry {
	e.e = e.e.Strs(key, s)
	return e
}

func (e *EntryAdapter) Time(key string, t time.Time) khl.Entry {
	e.e = e.e.Time(key, t)
	return e
}
