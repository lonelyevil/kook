package kook

import (
	"net"
	"time"
)

//const (
//	LogFatal = iota
//	LogError
//	LogWarning
//	LogInfo
//	LogDebug
//	LogTrace
//)

//var LogFunc func(level, caller int, format string, a ...interface{})
//
//func mlog(level, caller int, format string, a ...interface{}) {
//	if LogFunc != nil {
//		LogFunc(level, caller, format, a...)
//		return
//	}
//	pc, file, line, _ := runtime.Caller(caller)
//
//	files := strings.Split(file, "/")
//	file = files[len(files)-1]
//
//	name := runtime.FuncForPC(pc).Name()
//	fns := strings.Split(name, ".")
//	name = fns[len(fns)-1]
//
//	msg := fmt.Sprintf(format, a...)
//
//	log.Printf("[DG%d] %s:%d:%s() %s\n", level, file, line, name, msg)
//}
//
//func (s *Session) log(level int, format string, a ...interface{})  {
//	if level > s.LogLevel {
//		return
//	}
//	mlog(level, 2, format, a...)
//}

func addCaller(e Entry) Entry {
	return e.Caller(3)
}

// Logger is the interface which external loggers should meet.
type Logger interface {
	Trace() Entry
	Debug() Entry
	Info() Entry
	Warn() Entry
	Error() Entry
	Fatal() Entry
}

// Entry is the interface which external loggers' entry should meet.
type Entry interface {
	Bool(key string, b bool) Entry
	Bytes(key string, val []byte) Entry
	Caller(depth int) Entry
	Dur(key string, d time.Duration) Entry
	Err(key string, err error) Entry
	Float64(key string, f float64) Entry
	IPAddr(key string, ip net.IP) Entry
	Int(key string, i int) Entry
	Int64(key string, i int64) Entry
	Interface(key string, i interface{}) Entry
	Msg(msg string)
	Msgf(f string, i ...interface{})
	Str(key string, s string) Entry
	Strs(key string, s []string) Entry
	Time(key string, t time.Time) Entry
}

// mock types are only for possible testing usage, do not use it in real project.
// type mockLogger struct {
// }

// func (m mockLogger) Trace() Entry {
// 	return mockEntry{}
// }

// func (m mockLogger) Debug() Entry {
// 	return mockEntry{}
// }

// func (m mockLogger) Info() Entry {
// 	return mockEntry{}
// }

// func (m mockLogger) Warn() Entry {
// 	return mockEntry{}
// }

// func (m mockLogger) Error() Entry {
// 	return mockEntry{}
// }

// func (m mockLogger) Fatal() Entry {
// 	return mockEntry{}
// }

// type mockEntry struct {
// }

// func (m mockEntry) Strs(_ string, _ []string) Entry {
// 	return m
// }

// func (m mockEntry) Bool(_ string, _ bool) Entry {
// 	return m
// }

// func (m mockEntry) Bytes(_ string, _ []byte) Entry {
// 	return m
// }

// func (m mockEntry) Caller(_ int) Entry {
// 	return m
// }

// func (m mockEntry) Dur(_ string, _ time.Duration) Entry {
// 	return m
// }

// func (m mockEntry) Err(_ string, _ error) Entry {
// 	return m
// }

// func (m mockEntry) Float64(_ string, _ float64) Entry {
// 	return m
// }

// func (m mockEntry) IPAddr(_ string, _ net.IP) Entry {
// 	return m
// }

// func (m mockEntry) Int(_ string, _ int) Entry {
// 	return m
// }

// func (m mockEntry) Int64(_ string, _ int64) Entry {
// 	return m
// }

// func (m mockEntry) Interface(_ string, _ interface{}) Entry {
// 	return m
// }

// func (m mockEntry) Msg(_ string) {
// }

// func (m mockEntry) Msgf(_ string, _ ...interface{}) {
// }

// func (m mockEntry) Str(_ string, _ string) Entry {
// 	return m
// }

// func (m mockEntry) Time(_ string, _ time.Time) Entry {
// 	return m
// }
