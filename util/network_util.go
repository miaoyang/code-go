package util

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

// GetIp 获取请求的IP
func GetIp(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if len(ip) > 0 {
		return ip
	}

	ip = r.Header.Get("X-Forwarded-For")
	if len(ip) > 0 {
		ips := strings.Split(ip, ", ")
		lastIndex := len(ips) - 1
		return ips[lastIndex]
	}

	ip = r.RemoteAddr
	log.Println("ip ", ip)
	if strings.Contains(ip, ":") {
		ip = strings.Split(ip, ":")[0]
	}

	log.Println("get ip: ", ip)
	if !ValidateIP(ip) {
		log.Println("validate ip fail")
		return ""
	}
	return ip
}

// ValidateIP 校验 IP 地址
func ValidateIP(ip string) bool {
	// 定义 IP 地址的正则表达式
	regex := `^(\d{1,3}\.){3}\d{1,3}$`

	// 编译正则表达式
	re := regexp.MustCompile(regex)

	// 校验 IP 地址
	return re.MatchString(ip)
}

// GetIpAddress get ip from request
func GetIpAddress(r *http.Request) string {
	var ipAddress string

	ipAddress = r.RemoteAddr
	log.Println("ipAddress: ", ipAddress)
	if ipAddress != "" {
		ipAddress = strings.Split(ipAddress, ":")[0]
	}

	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		for _, ip := range strings.Split(r.Header.Get(h), ",") {
			if ip != "" {
				ipAddress = ip
			}
		}
	}
	return ipAddress
}
