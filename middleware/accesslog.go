package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv14/global"
	"github.com/liuhongdi/digv14/pkg/util"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().UnixNano()
		c.Next()
		endTime := time.Now().UnixNano()
		duration:=endTime-beginTime

		s := "%s %s \"%s %s\" "+
			"%s %d %d %dµs "+
			"\"%s\""

		layout := "2006-01-02 15:04:05"
		timeNow := time.Now().Format(layout)

       global.AccessLogger.Infof(s,
		   util.GetRealIp(c),
		   timeNow,
       	   c.Request.Method,
		   c.Request.RequestURI,
		   c.Request.Proto,
		   bodyWriter.Status(),
		   bodyWriter.body.Len(),
		   duration/1000,
		   c.Request.Header.Get("User-Agent"),

		   );
	}
}

/*
127.0.0.1 - [Sat, 28 Nov 2020 16:16:20 CST] "GET /article/getone/2" HTTP/1.1 200 171 2.201217ms "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:83.0) Gecko/201001
01 Firefox/83.0"
*/

/*
	s := "time:%s, method: %s, status_code: %d, " +
		"begin_time: %d, end_time: %d, "+
		"request:%s, response:%s, uri:%s,"+
		"protocol:%s, host:%s, "+
		"header:%s, length:%d "+
		"ip:%s, duration: %d "
*/
/*
		   beginTime,
		   endTime,
		   c.Request.PostForm.Encode(),
		   bodyWriter.body.String(),
		   c.Request.Host,
		   c.Request.Header,
*/