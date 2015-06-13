package icb

import (
    "bytes"
    "net"
    "fmt"
)

type IcbPacket struct {
    Buffer bytes.Buffer
}

func (packet *IcbPacket) Init (kind []byte, parameters []string) { 	
    packet.Buffer.Write(kind)
    packet.packParameters(parameters)
}

func (packet *IcbPacket) Write (p []byte) (n int, err error) {
    return packet.Buffer.Write(p)
}

func (packet *IcbPacket) SendTo (connection net.Conn) {
    connection.Write([]byte{byte(packet.Buffer.Len())})
    connection.Write(packet.Buffer.Bytes())
}

func Connection(host, port string)(net.Conn, error) {
    address := fmt.Sprintf("%s:%s", host, port)
    return net.Dial("tcp", address)
}

