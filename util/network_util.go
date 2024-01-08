package util

import (
	"fmt"
	"log"
	"math/rand"
	"net"
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

// GetAvailablePort 查询可用端口
func GetAvailablePort(minPort, maxPort, defaultPort int) int {
	if isPortAvailable(defaultPort) {
		return defaultPort
	}
	if minPort >= maxPort {
		log.Printf("minPort>=maxPort, %d, %d\n", minPort, maxPort)
		return -1
	}

	portRange := maxPort - minPort
	var searchCount = 0
	for searchCount <= portRange {
		nextPort := findRandomPort(minPort, maxPort)
		if isPortAvailable(nextPort) {
			return nextPort
		}
	}
	return -1
}

// isPortAvailable 端口是否可用
func isPortAvailable(port int) bool {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Printf("port: %d is not available\n", port)
		return false
	}
	defer listener.Close()
	return true
}

// 查询下一个可用的端口 findRandomPort
func findRandomPort(minPort, maxPort int) int {
	portRange := maxPort - minPort
	// intn [0,range)
	nextPort := minPort + rand.Intn(portRange+1)
	return nextPort
}
