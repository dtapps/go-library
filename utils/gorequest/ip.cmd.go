package gorequest

import (
	"bufio"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

func getCmdIPV4() string {
	if runtime.GOOS == "windows" {
		return getCmdIPV4Windows()
	} else if runtime.GOOS == "darwin" {
		return getCmdIPV4Mac()
	}
	return getCmdIPV4Linux()
}

func getCmdIPV6() string {
	if runtime.GOOS == "windows" {
		return getCmdIPV6Windows()
	} else if runtime.GOOS == "darwin" {
		return getCmdIPV6Mac()
	}
	return getCmdIPV6Linux()
}

func getCmdIPV4Linux() string {
	// 执行 ifconfig | grep 'inet ' | awk '{print $2}' 命令
	cmd := exec.Command("bash", "-c", "ifconfig | grep 'inet ' | awk '{print $2}'")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// 解析输出
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		ipv4 := scanner.Text()
		if IsIPV4(ipv4) && IsIPv4Public(net.ParseIP(ipv4)) {
			return ipv4
		}
	}

	return ""
}

func getCmdIPV6Linux() string {
	// 执行 ip -6 addr | grep inet6 | awk -F '[ \t]+|/' '$3 == "::1" { next;} $3 ~ /^fe80::/ { next;} /inet6/ {print $3}' 命令
	//cmd := exec.Command("bash", "-c", "ip -6 addr | grep inet6 | awk -F '[ \\t]+|/' '$3 == \"::1\" { next;} $3 ~ /^fe80::/ { next;} /inet6/ {print $3}'")
	//output, err := cmd.Output()
	// 执行 ip -6 addr | grep inet6 | awk -F '[ \t]+|/' '$3 == "::1" { next;} {print $3}' 命令
	cmd := exec.Command("bash", "-c", "ip -6 addr | grep inet6 | awk -F '[ \\t]+|/' '$3 == \"::1\" { next;} {print $3}'")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// 解析输出
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		ipv6 := scanner.Text()
		if IsIPV6(ipv6) && IsIPv6Public(net.ParseIP(ipv6)) {
			return ipv6
		}
	}

	return ""
}

func getCmdIPV4Windows() string {
	// 执行 ipconfig | findstr IPv4 命令
	cmd := exec.Command("cmd", "/c", "ipconfig", "|", "findstr", "IPv4")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// 解析输出
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			ipv4 := fields[len(fields)-1]
			if IsIPV4(ipv4) && IsIPv4Public(net.ParseIP(ipv4)) {
				return ipv4
			}
		}
	}

	return ""
}

func getCmdIPV6Windows() string {
	// 执行 ipconfig | findstr IPv6 命令
	cmd := exec.Command("cmd", "/c", "ipconfig", "|", "findstr", "IPv6")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// 解析输出
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			ipv6 := fields[len(fields)-1]
			if IsIPV6(ipv6) && IsIPv6Public(net.ParseIP(ipv6)) {
				return ipv6
			}
		}
	}

	return ""
}

func getCmdIPV4Mac() string {
	// 执行 ifconfig | grep inet | grep -v inet6 | awk '{print $2}'
	cmd := exec.Command("bash", "-c", "ifconfig | grep inet | grep -v inet6 | awk '{print $2}'")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// 解析输出
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		ipv4 := scanner.Text()
		if IsIPV4(ipv4) && IsIPv4Public(net.ParseIP(ipv4)) {
			return ipv4
		}
	}

	return ""
}

func getCmdIPV6Mac() string {
	// 执行 ifconfig | grep inet6 | awk -F '[ \t]+|/' '$3 == "::1" { next;} $3 ~ /^fe80::/ { next;} /inet6/ {print $3}' 命令
	//cmd := exec.Command("bash", "-c", "ifconfig | grep inet6 | awk -F '[ \\t]+|/' '$3 == \"::1\" { next;} $3 ~ /^fe80::/ { next;} /inet6/ {print $3}'")
	//output, err := cmd.Output()
	// 执行 ifconfig | grep inet6 | awk -F '[ \t]+|/' '$3 == "::1" { next;} /inet6/ {print $3}' 命令
	cmd := exec.Command("bash", "-c", "ifconfig | grep inet6 | awk -F '[ \\t]+|/' '$3 == \"::1\" { next;} /inet6/ {print $3}'")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// 解析输出
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		ipv6 := scanner.Text()
		if IsIPV6(ipv6) && IsIPv6Public(net.ParseIP(ipv6)) {
			return ipv6
		}
	}

	return ""
}
