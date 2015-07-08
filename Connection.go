package icb

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// Connection is the singleton icb connection
var Connection net.Conn

// Connect connects to the icb host and port provided, returns a net.Conn and error
func Connect(host, port string) {
	address := fmt.Sprintf("%s:%s", host, port)

	var err error
	Connection, err = net.Dial("tcp", address)
	if err == nil {
		Connection.SetDeadline(time.Time{}) //   do not time out on i/o operations
	} else {
		log.Fatal("Failed:", err)
		os.Exit(-1)
	}
}
