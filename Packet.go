package icb

import (
	"bytes"
	"fmt"
	"net"
)

// Packet holds raw icb packets
type Packet struct {
	Buffer bytes.Buffer
}

// packParameters copies the parameters into the buffer seperated by \001 charaters
func (packet *Packet) packParameters(parameters []string) {
	for _, parameter := range parameters {
		packet.Buffer.WriteString(parameter)
		packet.Buffer.WriteByte(1)
	}
	packet.Buffer.Truncate(packet.Buffer.Len() - 1) // remove the excess separator
}

// Init initializes a packet of type kind with the parameters specified
func (packet *Packet) Init(kind []byte, parameters []string) {
	packet.Buffer.Write(kind)
	packet.packParameters(parameters)
}

// Write copies a byte array into the packet's buffer
func (packet *Packet) Write(p []byte) (n int, err error) {
	return packet.Buffer.Write(p)
}

// SendTo writes the packet buffer to the supplied connection
func (packet *Packet) SendTo(connection net.Conn) {
	connection.Write([]byte{byte(packet.Buffer.Len())})
	connection.Write(packet.Buffer.Bytes())
}

// Connect connects to the icb host and port provided, returns a net.Conn and error
func Connect(host, port string) (net.Conn, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	return net.Dial("tcp", address)
}
