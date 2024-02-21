package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parseTrojanLine remains the same as previously defined

// 解析 trojan 数据行
func parseTrojanLine(line string) (password, domain, port string, err error) {
	if !strings.HasPrefix(line, "trojan://") {
		return "", "", "", fmt.Errorf("invalid trojan line format")
	}
	line = strings.TrimPrefix(line, "trojan://")

	// 分割 '#' 以去除后半部分
	parts := strings.Split(line, "#")
	if len(parts) == 0 {
		return "", "", "", fmt.Errorf("invalid trojan line format")
	}

	line = parts[0]

	// 分割 '@' 来提取密码
	atSplit := strings.LastIndex(line, "@")
	if atSplit == -1 {
		return "", "", "", fmt.Errorf("invalid trojan line format")
	}
	password = line[:atSplit]

	// 提取域名和端口
	hostPort := line[atSplit+1:]
	colonSplit := strings.LastIndex(hostPort, ":")
	if colonSplit == -1 {
		return "", "", "", fmt.Errorf("invalid trojan line format")
	}
	domain = hostPort[:colonSplit]
	port = hostPort[colonSplit+1:]

	return password, domain, port, nil
}

func main() {
	// 定义命令行参数
	trojanFilePath := flag.String("t", "", "Path to the trojan.txt file")
	configFilePath := flag.String("c", "", "Path to the config.json file")
	//定义初始端口值，随后进行递增
	tagCounterPtr := flag.Int("p", 10000, "Starting value for tag counter")
	userPtr := flag.String("u", "", "Username for inbound settings")
	passPtr := flag.String("pwd", "", "Password for inbound settings")
	flag.Parse()

	tagCounter := *tagCounterPtr
	user := *userPtr
	pass := *passPtr
	// 处理 trojan.txt 文件
	configFileData, err := os.ReadFile(*configFilePath)
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	var config Config
	err = json.Unmarshal(configFileData, &config)
	if err != nil {
		fmt.Printf("Error unmarshalling config file: %s\n", err)
		return
	}

	// 处理 trojan.txt 文件
	trojanFile, err := os.Open(*trojanFilePath)
	if err != nil {
		fmt.Printf("Error opening trojan file: %s\n", err)
		return
	}
	defer trojanFile.Close()

	scanner := bufio.NewScanner(trojanFile)
	for scanner.Scan() {
		line := scanner.Text()
		password, domain, portStr, err := parseTrojanLine(line)
		if err != nil {
			fmt.Printf("Error parsing trojan line: %s\n", err)
			continue
		}

		port, err := strconv.Atoi(portStr)
		if err != nil {
			fmt.Printf("Error converting port to integer: %s\n", err)
			continue
		}
		newInbound := Inbound{
			Protocol: "socks", // 根据实际协议调整
			Port:     tagCounter,
			Listen:   nil,
			Settings: map[string]interface{}{
				"auth": "password",
				"accounts": []map[string]interface{}{
					{
						"user": user,
						"pass": pass,
					},
				},
				"udp": false,
				"ip":  "127.0.0.1",
			},
			Tag:            fmt.Sprintf("inbound-0.0.0.0:%d", tagCounter),
			StreamSettings: nil,
			Sniffing:       nil,
		}
		config.Inbounds = append(config.Inbounds, newInbound)
		// 定义默认的 StreamSettings
		defaultStreamSettings := StreamSettings{
			Network:  "tcp",
			Security: "tls",
			TLSSettings: &TLSSettings{
				ServerName:    "",
				ALPN:          []string{},
				Fingerprint:   "",
				AllowInsecure: false,
			},
			TCPSettings: &TCPSettings{
				Header: Header{
					Type: "none",
				},
			},
		}

		// 创建新的 Outbound 对象并添加到数组中
		newOutbound := Outbound{
			Tag:      fmt.Sprintf("trojan%d", tagCounter),
			Protocol: "trojan", // 根据实际需要可能需要调整
			Settings: OutboundSettings{
				Servers: []Server{
					{
						Address:  domain,
						Port:     port,
						Password: password,
					},
				},
			},
			StreamSettings: &defaultStreamSettings,
		}
		config.Outbounds = append(config.Outbounds, newOutbound)

		newRule := Rule{
			Type:        "field",
			InboundTag:  []string{fmt.Sprintf("inbound-0.0.0.0:%d", tagCounter)},
			OutboundTag: fmt.Sprintf("trojan%d", tagCounter),
		}
		config.Routing.Rules = append(config.Routing.Rules, newRule)
		tagCounter++
	}

	// 将修改后的数据写回 config.json
	updatedConfigData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling updated config: %s\n", err)
		return
	}
	fmt.Println("Updated config.json content:")
	fmt.Println(string(updatedConfigData))
	err = os.WriteFile(*configFilePath, updatedConfigData, 0644)
	if err != nil {
		fmt.Printf("Error writing updated config to file: %s\n", err)
	}
}
