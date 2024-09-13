package log

import (
	"FinancialAssistanceScheme/controller"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"io"
	"log"
	"time"
)

type JSONLog struct {
	IP           string `json:"ip"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	Status       int    `json:"status"`
	ResponseTime int64  `json:"response_time_ms"`
	Timestamp    string `json:"timestamp"`
}

func AccessLogMiddleware(ctx iris.Context) {
	// Get request details
	method := ctx.Method()
	path := ctx.Path()
	ip := ctx.RemoteAddr()

	start := time.Now()
	//
	// Continue to the next handler
	ctx.Next()
	//
	// Get response details after the request is processed
	status := ctx.GetStatusCode()
	//ctx.Values().Set("status_message", "OK")
	duration := time.Since(start)
	//
	//// Log the request and response details
	//log.Printf("[%s] %s %s - %d %s", ip, method, path, status, duration)

	// Create a log entry
	logEntry := JSONLog{
		IP:           ip,
		Method:       method,
		Path:         path,
		Status:       status,
		ResponseTime: duration.Milliseconds(),
		Timestamp:    time.Now().Format(time.RFC3339),
	}

	// Marshal the log entry into JSON
	logData, err := json.Marshal(logEntry)
	if err != nil {
		log.Printf("Error marshaling JSON log: %v", err)
		return
	}

	// Log to console (or you can write to a file)
	log.Println(string(logData))

}

func InitAccessLog(w io.Writer) *accesslog.AccessLog {
	ac := accesslog.New(w)
	ac.Async = true

	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.IP = true
	ac.LatencyRound = time.Millisecond //按毫秒进行Round
	ac.BytesReceivedBody = true
	ac.BytesSentBody = false
	ac.BytesReceived = false
	ac.BytesSent = true // must be true, so we can touch response
	ac.BodyMinify = false
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler

	ac.AddFields(func(ctx iris.Context, fields *accesslog.Fields) {
		var resp controller.BaseResponse
		if err := json.Unmarshal(ctx.Recorder().Body(), &resp); err == nil {
			fields.Set("api.code", resp.Code)
			fields.Set("api.message", resp.Message)
		}
	})

	ac.SetFormatter(&accesslog.JSON{
		HumanTime: true,
	})

	return ac
}
