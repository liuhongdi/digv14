package util

import (
	"github.com/gin-gonic/gin"
	"net"
	//"net/http"
	"strings"
)

func GetRealIp(r *gin.Context) string {
	ip, _, err := net.SplitHostPort(r.Request.RemoteAddr)
	if err != nil {
		ip = r.Request.RemoteAddr
	}
	if ip != "127.0.0.1" {
		return ip
	}
	// Check if behide nginx or apache
	xRealIP := r.Request.Header.Get("X-Real-Ip")
	xForwardedFor :=  r.Request.Header.Get("X-Forwarded-For")

	for _, address := range strings.Split(xForwardedFor, ",") {
		address = strings.TrimSpace(address)
		if address != "" {
			return address
		}
	}

	if xRealIP != "" {
		return xRealIP
	}
	return ip
}
