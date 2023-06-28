package logger

import "testing"

func TestLogger_UpdateLog(t *testing.T) {
	option := &Options{
		Url:        "http://127.0.0.1:8000",
		Db:         "logs",
		Collection: "task",
	}
	lgr := NewLogger(option)
	lgr.UpdateLog("649a8f81502ebd74179b8621", "nc2db", "test_time")
	t.Log(option.Url)
}
