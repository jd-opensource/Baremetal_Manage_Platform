package cps_log

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

// Level type
type Level uint32

const (
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel Level = iota
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel
)

var globalDefinedKeys = []string{"time", "logid", "self_t", "all_t"}

type Logger struct {
	data        map[string]interface{}
	level       Level
	fileName    string
	definedKeys []string
}

func New(filename string) *Logger {
	if filename == "" {
		log.Fatal("filename must be provided !!!")
		return nil
	}
	l := &Logger{
		data:     make(map[string]interface{}),
		fileName: filename,
	}
	l.definedKeys = append([]string{}, globalDefinedKeys...)
	l.data["time"] = time.Now().Format("2006-01-02T15:04:05")
	return l
}

func NewLoggerWithConfig(filename string, level Level, definedKeys []string) *Logger {
	l := New(filename)
	l.SetLevel(level)
	l.SetStableFields(definedKeys)
	return l
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
}

//SetStableFields 需要固定每条日志的字段时，通过此接口设置，global在前，userdefined在后
func (l *Logger) SetStableFields(fields []string) {
	l.definedKeys = append(l.definedKeys, fields...)
}

func (l *Logger) Point(key string, val interface{}) {
	l.data[key] = val
}

func (l *Logger) GetPoint(key string) interface{} {
	return l.data[key]
}

func (l *Logger) TimeStart(key string) {
	if _, ok := l.data[key]; ok {
		log.Print(fmt.Sprintf("point %s TimeStart more than once, check !!!", key))
	}
	l.data[key] = time.Now()

}

func (l *Logger) TimeEnd(key string) {
	if t, ok := l.data[key]; ok {
		v := time.Now().Sub(t.(time.Time)).Milliseconds()
		l.data[key] = fmt.Sprintf("%f", float64(v)/1000.0)
	} else {
		log.Print(fmt.Sprintf("missing call TimeStart before TimeEnd when point %s, check it!!!", key))
	}

}

func (l *Logger) PushPoint(key string, val interface{}) {
	if _, ok := l.data[key]; !ok {
		l.data[key] = []interface{}{}
	}
	d := l.data[key].([]interface{})
	d = append(d, val)
}

func (l *Logger) Info(v ...interface{}) {
	if l.level > InfoLevel {
		return
	}
	logid, _ := l.data["logid"].(string)
	content := time.Now().Format("2006-01-02T15:04:05") + fmt.Sprintf(" [%s] ", logid) + fmt.Sprintln(v...)
	filename := fmt.Sprintf("%s.INFO", l.fileName)
	WriteFile(filename, strings.TrimSpace(content)+"\r\n")
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level > InfoLevel {
		return
	}
	logid, _ := l.data["logid"].(string)
	content := time.Now().Format("2006-01-02T15:04:05") + fmt.Sprintf(" [%s] ", logid) + fmt.Sprintf(format, v...)
	filename := fmt.Sprintf("%s.INFO", l.fileName)
	WriteFile(filename, strings.TrimSpace(content)+"\r\n")
}

func (l *Logger) Warn(v ...interface{}) {
	if l.level > WarnLevel {
		return
	}
	logid, _ := l.data["logid"].(string)
	content := time.Now().Format("2006-01-02T15:04:05") + fmt.Sprintf(" [%s] ", logid) + fmt.Sprintln(v...)
	filename := fmt.Sprintf("%s.WARN", l.fileName)
	WriteFile(filename, strings.TrimSpace(content)+"\r\n")
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level > WarnLevel {
		return
	}
	logid, _ := l.data["logid"].(string)
	content := time.Now().Format("2006-01-02T15:04:05") + fmt.Sprintf(" [%s] ", logid) + fmt.Sprintf(format, v...)
	filename := fmt.Sprintf("%s.WARN", l.fileName)
	WriteFile(filename, strings.TrimSpace(content)+"\r\n")
}

func (l *Logger) Fatal(v ...interface{}) {
	logid, _ := l.data["logid"].(string)
	content := time.Now().Format("2006-01-02T15:04:05") + fmt.Sprintf(" [%s] ", logid) + fmt.Sprintln(v...)
	filename := fmt.Sprintf("%s.FATAL", l.fileName)
	WriteFile(filename, strings.TrimSpace(content)+"\r\n")
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	logid, _ := l.data["logid"].(string)
	content := time.Now().Format("2006-01-02T15:04:05") + fmt.Sprintf(" [%s] ", logid) + fmt.Sprintf(format, v...)
	filename := fmt.Sprintf("%s.FATAL", l.fileName)
	WriteFile(filename, strings.TrimSpace(content)+"\r\n")
}

func (l *Logger) Flush() {
	if l.level > DebugLevel {
		return
	}
	var beforeStr, afterStr string
	for _, key := range l.definedKeys {
		if val, ok := l.data[key]; ok {
			if key == "time" {
				beforeStr = val.(string)
				continue
			}
			switch reflect.TypeOf(val).Kind() {
			case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
				beforeStr = beforeStr + " " + key + ":" + fmt.Sprintf("%d", val)
			case reflect.String:
				beforeStr = beforeStr + " " + key + ":" + val.(string)
			case reflect.Array, reflect.Slice, reflect.Ptr, reflect.Interface, reflect.Map, reflect.Struct:
				if s, err := json.Marshal(val); err == nil {
					beforeStr = beforeStr + " " + key + ":" + string(s)
				} else {
					log.Print(fmt.Sprintf("val for %s Marshal error: %s", key, err.Error()))
				}
			case reflect.Func, reflect.Chan, reflect.UnsafePointer:
				log.Print(fmt.Sprintf("val for %s not support channel or function, modify it!!!", key))
			default:
				log.Print(fmt.Printf("val for %s unknown type!!!", key))

			}
		}
	}
	for key, val := range l.data {
		if exist, _ := inArray(key, l.definedKeys); !exist {
			switch reflect.TypeOf(val).Kind() {
			case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
				afterStr = afterStr + " " + key + ":" + fmt.Sprintf("%d", val)
			case reflect.String:
				afterStr = afterStr + " " + key + ":" + val.(string)
			case reflect.Array, reflect.Slice, reflect.Ptr, reflect.Interface, reflect.Map, reflect.Struct:
				if s, err := json.Marshal(val); err == nil {
					afterStr = afterStr + " " + key + ":" + string(s)
				} else {
					log.Print(fmt.Sprintf("val for %s Marshal error: %s", key, err.Error()))
				}
			case reflect.Func, reflect.Chan, reflect.UnsafePointer:
				log.Print(fmt.Sprintf("val for %s not support channel or function, modify it!!!", key))
			default:
				log.Print(fmt.Sprintf("val for %s unknown type!!!", key))
			}
		}
	}

	line := strings.TrimSpace(beforeStr) + " " + strings.TrimSpace(afterStr) + "\r\n"
	filename := fmt.Sprintf("%s.DEBUG", l.fileName)
	WriteFile(filename, line)
}

func WriteFile(filename, line string) {
	TimeStamp := time.Now()
	str := filename + "." + TimeStamp.Format("2006-01-02")
	logf, err := os.OpenFile(str, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("open filename error:", err.Error())
	}
	defer logf.Close()
	if _, err := io.WriteString(logf, line); err != nil {
		log.Fatal("write line  error:", err.Error())
	}
}

func inArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}
