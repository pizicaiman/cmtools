package cmtools

import (
	"log"
	"net"
	"os"
	"strings"
	"time"
)

/**
 * @Classname funcs
 * @Description TODO
 * @author cjf
 * @Date 2021/7/15 13:42
 * @Version V1.0
 */

func GetNowTimeStr() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func GetHostName() string {
	name, _ := os.Hostname()
	return name +"v1.0.2"
}

func GetLocalIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Printf("get local addr err:%v", err)
		return ""
	} else {
		localIp := strings.Split(conn.LocalAddr().String(), ":")[0]
		conn.Close()
		return localIp
	}
}
