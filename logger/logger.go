package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Options struct {
	Url        string
	Db         string
	Collection string
}

type Logger struct {
	url        string
	db         string
	collection string
}

func NewLogger(options *Options) *Logger {
	return &Logger{
		url:        options.Url,
		db:         options.Db,
		collection: options.Collection,
	}
}

func (l *Logger) UpdateLog(mgoId, step, field string) {
	// 构建请求体数据
	data := make(map[string]interface{})
	key := fmt.Sprintf("%s.%s", step, field)
	data[key] = time.Now().Format("2006-01-02 15:04:05")
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON编码失败:", err)
		return
	}

	fullUrl := fmt.Sprintf("%s/update/%s/%s/%s", l.url, l.db, l.collection, mgoId)
	// 发送 POST 请求
	resp, err := http.Post(fullUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
}
