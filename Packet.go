package icb

import "bytes"

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

// Send writes the packet buffer to the supplied connection
func (packet *Packet) Send() {
	Connection.Write([]byte{byte(packet.Buffer.Len())})
	Connection.Write(packet.Buffer.Bytes())
}
