package elnet

import (
	"encoding/binary"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const PatternDNSName = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`

var RegDNSName = regexp.MustCompile(PatternDNSName)

// GetLocalIP Returns the first NIC's local IP.
func GetLocalIP() (string, error) {
	res := ""
	addrs, err := net.InterfaceAddrs()
	if err != nil || len(addrs) <= 0 {
		return res, err
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			res = ipNet.IP.String()
			break
		}
	}
	return res, err
}

// GetOutboundIP returns the outbound IP address.
func GetOutboundIP() (string, error) {
	res := ""
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if conn != nil {
		addr := conn.LocalAddr().(*net.UDPAddr)
		res = addr.IP.String()
		_ = conn.Close()
	}
	return res, err
}

// IsPublicIP returns true if ip is a public ip address.
func IsPublicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := ip.To4(); ip4 != nil {
		if ip4[0] == 10 {
			return false
		} else if ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31 {
			return false
		} else if ip4[0] == 192 && ip4[1] == 168 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

// IsPrivateIP returns true if ip is a private ip address.
func IsPrivateIP(ip net.IP) bool {
	return !IsPublicIP(ip)
}

// Get the list of ip address for this machine.
func GetIPs() (ips []string) {
	interfaceAddrs, err := net.InterfaceAddrs()
	if err != nil || len(interfaceAddrs) <= 0 {
		return
	}
	for _, addr := range interfaceAddrs {
		ipNet, isValidIpNet := addr.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}
	return
}

// Get the list of mac address for this machine.
func GetMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil || len(netInterfaces) <= 0 {
		return
	}
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if macAddr == "" {
			continue
		} else {
			macAddrs = append(macAddrs, macAddr)
		}
	}
	return
}

// Get the Hostname of this machine.
func Hostname() (string, error) {
	return os.Hostname()
}

// Get ip address by given hostname.
func GetIpByHostname(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if err == nil && ips != nil {
		for _, ip := range ips {
			if ip.To4() != nil {
				return ip.String(), nil
			}
		}
		return "", nil
	} else {
		return "", err
	}
}

// Get ip address by given domain.
func GetIpsByDomain(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if err == nil && ips != nil {
		var ipStrings []string
		for _, ip := range ips {
			if ip.To4() != nil {
				ipStrings = append(ipStrings, ip.String())
			}
		}
		return ipStrings, nil
	} else {
		return nil, err
	}
}

// Get hostname by given ip address.
func GetHostnameByIp(ip string) (string, error) {
	names, err := net.LookupAddr(ip)
	if err == nil && names != nil {
		return strings.TrimRight(names[0], "."), nil
	} else {
		return "", err
	}
}

// IsValidPort returns true if the port is valid number.
func IsValidPort(port int) bool {
	if port > 0 && port < 65536 {
		return true
	} else {
		return false
	}
}

// IsValidIP returns true if ip is a valid IP address.
func IsValidIP(ip string) bool {
	return ip != "" && net.ParseIP(ip) != nil
}

// IsValidDNSName returns true if name is a valid DNS name.
func IsValidDNSName(name string) bool {
	if name == "" || len(strings.Replace(name, ".", "", -1)) > 255 {
		// constraints already violated
		return false
	}
	return !IsValidIP(name) && RegDNSName.MatchString(name)
}

// IsValidHost returns true if the host is a valid hostname.
func IsValidHost(host string) bool {
	return IsValidIP(host) || IsValidDNSName(host)
}

// IsPortOpen returns true if the port on the host is open.
func IsPortOpen(host string, port int, protocols ...string) bool {
	if IsValidHost(host) && IsValidPort(port) {
		protocol := "tcp"
		if len(protocols) > 0 && len(protocols[0]) > 0 {
			protocol = strings.ToLower(protocols[0])
		}
		conn, _ := net.DialTimeout(protocol, net.JoinHostPort(host, strconv.Itoa(port)), time.Second*5)
		if conn != nil {
			_ = conn.Close()
			return true
		}
	}
	return false
}

// convert ipv4 address to uint
// if `ipAddress` is not a valid ip, return 0
func Ip2Long(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

// convert uint to ipv4 address.
func Long2Ip(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}

// Get text from target url, If some error occurred, return empty string.
func getTextFromUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	} else {
		defer resp.Body.Close()
		if b, err := io.ReadAll(resp.Body); err != nil {
			return ""
		} else {
			return string(b)
		}
	}
}

// Try get the public IPv4 address of this machine.
// If some error occurred, return empty string.
func GetMyPublicIPv4() string {
	if res := getTextFromUrl("https://inet-ip.info/ip"); res != "" && len(res) <= 15 && IsValidIP(res) {
		return res
	} else if res = getTextFromUrl("https://api.ipify.org/"); res != "" && len(res) <= 15 && IsValidIP(res) {
		return res
	} else {
		res = getTextFromUrl("http://checkip.dyndns.com/")
		if res == "" {
			return ""
		}
		info := strings.Split(res, ":")
		if len(info) != 2 {
			return ""
		}
		res = strings.TrimSpace(info[1])
		res = strings.TrimRight(res, "</body></html>\r\n")
		if IsValidIP(res) {
			return res
		} else {
			return ""
		}
	}
}

// Try get the public IPv6 address of this machine.
// If some error occurred, return empty string.
func GetMyPublicIPv6() string {
	if res := getTextFromUrl("https://ident.me"); len(res) >= 16 && IsValidIP(res) {
		return res
	} else if res = getTextFromUrl("https://icanhazip.com"); len(res) >= 16 && IsValidIP(res) {
		return res
	} else if res = getTextFromUrl("https://www.trackip.net/ip"); len(res) >= 16 && IsValidIP(res) {
		return res
	} else {
		return ""
	}
}

// Try get the public IP address of this machine.
// If some error occurred, return empty string.
func GetMyPublicIP() string {
	if res := GetMyPublicIPv4(); res != "" {
		return res
	} else if res = GetMyPublicIPv6(); res != "" {
		return res
	} else {
		return ""
	}
}
