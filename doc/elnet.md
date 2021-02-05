## elnet
- 网络方面的工具包
- utilities for network
- ネットワーク関連のユーティリティ
```go
import "github.com/el-ideal-ideas/ellib/elnet"

// 通过免费API确认自己的公网IP(ipv4)
// get the public ip address use free API(ipv4)
// 無料APIを使用してパブリックIPを確認(ipv4)
elnet.GetMyPublicIPv4() string

// 通过免费API确认自己的公网IP(ipv6)
// get the public ip address use free API(ipv6)
// 無料APIを使用してパブリックIPを確認(ipv6)
elnet.GetMyPublicIPv4() string

// 通过免费API确认自己的公网IP
// get the public ip address use free API
// 無料APIを使用してパブリックIPを確認
elnet.GetMyPublicIP() string

// 获取内网IP
// the first NIC's local IP.
// プライベートIPを取得
elnet.GetLocalIP() (string, error)

// 获取 outbound IP
// get outbound IP
// outbound IP を取得
elnet.GetOutboundIP() (string, error)

// 检查是否为公网IP
// If ip is a public ip address, return true
// パブリックIPかどうかを確認する
elnet.IsPublicIP(ip net.IP) bool

// 检查是否为区域网IP
// If ip is a private ip address, return true
// プライベートIPかどうかを確認する
elnet.IsPrivateIP(ip net.IP) bool

// 获取被分配到的全部IP地址
// get all ip address for this machine
// 割り当てられた全てのIPを取得する
elnet.GetIPs() (ips []string)

// 获取Mac地址
// get mac address for this machine.
// 割り当てられたMacアドレスを取得
elnet.GetMacAddrs() (macAddrs []string)

// 获取主机名
// get the hostname of this machine
// ホスト名を取得
elnet.Hostname() (string, error)

// 尝试通过主机名获取对象IP
// get IP address by given hostname
// 与えられたホスト名に対応するIPを取得
elnet.GetIpByHostname(hostname string) (string, error)

// 尝试通过域名获取对象IP
// get IP address by given domain
// 与えられたドメインに対応するIPを取得
elnet.GetIpsByDomain(domain string) ([]string, error)

// 通过IP获取主机名
// get hostname by given IP address
// 与えられたIPに対応するホスト名を取得
elnet.GetHostnameByIp(ip string)

// 检测端口是否合法
// check the port number is valid
// ポート番号が正しい範囲かどうかを確認
elnet.IsValidPort(port int) bool

// 检测IP地址是否符合规范
// check ip address format is valid or not
// IPアドレスの形式が正しいかどうかを確認
elnet.IsValidIP(ip string) bool

// 检测DNS名是否符合规范
// check dns name format is valid or not
// DNS名の形式が正しいかどうかを確認
elnet.IsValidDNSName(name string)

// 检测主机名是否符合规范
// check host name format is valid or not
// ホスト名の形式が正しいかどうかを確認
elnet.IsValidHost(host string) bool

// 检测端口是否打开
// check if the port is open
// ポートが開いているかどうかを確認
elnet.IsPortOpen(host string, port int, protocols ...string) bool

// IP地址转换整数值
// convert ip address to uint32
// IPアドレスをuint32型に変換する
elnet.Ip2Long(ipAddress string) uint32

// 整数值转换到IP地址
// convert uint32 to ip address
// uint32をIPアドレスに変換する
elnet.Long2Ip(properAddress uint32) string
```