package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// ConnectionEntity 定义连接实体
type ConnectionEntity struct {
	Port uint32 `json:"port"`
	Pid  int32  `json:"pid"`
	Name string `json:"name"`
}

// KillProcessByPort 根据端口号关闭进程
func KillProcessByPort(port uint32) {
	entity := GetConnectionByPort(port)
	if entity == nil {
		fmt.Println("No process found on port", port)
		return
	}
	pr, err := process.NewProcess(entity.Pid)
	if err != nil {
		log.Fatal(err)
	}
	err = pr.Terminate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Process terminated:", entity.Name, "PID:", entity.Pid)
}

// GetConnectionByPort 根据端口号获取连接
func GetConnectionByPort(port uint32) *ConnectionEntity {
	connections := GetConnections()
	for _, con := range connections {
		if con.Port == port {
			return &con
		}
	}
	return nil
}

// GetConnections 获取处于 LISTEN 状态的 TCP 连接
func GetConnections() []ConnectionEntity {
	connections, _ := net.Connections("tcp")
	entities := make([]ConnectionEntity, 0)
	for _, con := range connections {
		if con.Status == "LISTEN" {
			entities = append(entities, ConnectionEntity{
				Port: con.Laddr.Port,
				Pid:  con.Pid,
				Name: GetPidName(con.Pid),
			})
		}
	}
	return entities
}

// List 列出端口信息
func List(port uint32) {
	connections := GetConnections()
	if port == 0 {
		fmt.Printf("%-10s%-10s%s\n", "port", "pid", "name")
		for _, con := range connections {
			fmt.Printf("%-10d%-10d%s\n", con.Port, con.Pid, con.Name)
		}
	} else {
		for _, con := range connections {
			if con.Port == port {
				fmt.Printf("%-10s%-10s%s\n", "port", "pid", "name")
				fmt.Printf("%-10d%-10d%s\n", con.Port, con.Pid, con.Name)
				break
			}
		}
	}
}

// GetPidName 根据 PID 获取进程名称
func GetPidName(pid int32) string {
	pro, err := process.NewProcess(pid)
	if err != nil {
		return "Unknown"
	}
	name, err := pro.Name()
	if err != nil {
		return "Unknown"
	}
	return name
}

func main() {
	args := os.Args
	if len(args) == 1 {
		List(0)
	} else {
		if isNumber(args[1]) {
			port, _ := strconv.Atoi(args[1])
			List(uint32(port))
		} else if args[1] == "-c" || args[1] == "-C" {
			if len(args) < 3 || !isNumber(args[2]) {
				errorCommand()
			}
			port, _ := strconv.Atoi(args[2])
			KillProcessByPort(uint32(port))
		} else {
			errorCommand()
		}
	}
}

func isNumber(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func errorCommand() {
	fmt.Println("unrecognizable command formatter")
}
