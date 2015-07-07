package icb

import (
	"fmt"
	"strconv"
	"time"
)

type packetMethod func() string

// Decode an icb packet
func (packet *Packet) Decode() string {
	kind, _ := packet.Buffer.ReadByte()

	switch string(kind) {
	case "a": // login sucessful
		{
			return ""
		}
	case "b":
		{
			return packet.publicMessage()
		}
	case "c":
		{
			return packet.privateMessage()
		}
	case "d":
		{
			return packet.statusMessage()
		}
	case "e":
		{
			return packet.errorMessage()
		}
	case "f":
		{
			return packet.importantMessage()
		}
	case "g":
		{
			return packet.serverExit()
		}
	case "i":
		{
			return packet.commandOutput()
		}
	case "k":
		{
			return packet.beep()
		}
	case "n": // nop packet
		{
			return ""
		}
	}

	return ""
}

func (packet *Packet) publicMessage() string {
	from, _ := packet.Buffer.ReadString(1)
	text := packet.Buffer.String()

	return fmt.Sprintf("<%s> %s", from, text)
}

func (packet *Packet) privateMessage() string {
	from, _ := packet.Buffer.ReadString(1)
	text := packet.Buffer.String()

	return fmt.Sprintf("<*%s*> %s", from, text)
}

func (packet *Packet) statusMessage() string {
	category, _ := packet.Buffer.ReadString(1)
	text := packet.Buffer.String()

	return fmt.Sprintf("[=%s=] %s", category, text)
}

func (packet *Packet) errorMessage() string {
	return fmt.Sprintf("[=Error=] %s", packet.Buffer.String())
}

func (packet *Packet) importantMessage() string {
	category, _ := packet.Buffer.ReadString(1)
	text := packet.Buffer.String()

	return fmt.Sprintf("[!%s!] %s", category, text)
}

func (packet *Packet) serverExit() string {
	return "[!Exit!] Warning the server is exiting"
}

func (packet *Packet) commandOutput() string {
	var kind = make([]byte, 2)
	packet.Buffer.Read(kind)

	switch string(kind) { // generic command output
	case "co":
		{
			return packet.Buffer.String()
		}

	case "ec": // end of output terminator in theory there should never be any actual message from this
		{
			return packet.Buffer.String()
		}
	case "wh": // output the who header
		{
			return "Nickname          Idle               Sign-On              Account"
		}

	case "wl": // item in a who listing
		{
			return packet.whoItem()
		}
	}
	return fmt.Sprintf("%%ERROR unknown command output packet of type %q\n", kind)
}

func (packet *Packet) whoItem() string {
	// kind, _  := packet.nextParameter()
	_ = packet.nextParameter()
	flag := packet.nextParameter() // moderator flag
	nick := packet.nextParameter()
	idle := packet.nextParameter()
	_ = packet.nextParameter() // responce time, obsolete
	login := packet.nextParameter()
	user := packet.nextParameter()
	host := packet.nextParameter()

	account := fmt.Sprintf("%s@%s", user, host)
	if flag == "m" {
		flag = "*"
	}

	idleTime := stringFromIdleSeconds(idle)
	loginTime := stringForLoginDate(login)

	return fmt.Sprintf("%s %-12s %-16s\t%-15s    %s", flag, nick, idleTime, loginTime, account)
}

func stringForLoginDate(login string) string {
	i, _ := strconv.ParseInt(login, 10, 64)
	const loginFormat = "Jan 02, 03:04pm"

	return time.Unix(i, 0).Format(loginFormat)
}

func stringFromIdleSeconds(idle string) string {
	seconds, _ := strconv.ParseInt(idle, 10, 64)

	// golang Time doesn't handle days so we have to do it ourselves
	days := seconds / (24 * 60 * 60)
	seconds -= days * 24 * 60 * 60

	idleString := ""
	if days > 0 {
		idleString += strconv.FormatInt(days, 10) + "d"
	}
	idleTime := time.Duration(seconds) * time.Second
	idleString += idleTime.String()

	return idleString
}

func (packet *Packet) nextParameter() string {
	param, _ := packet.Buffer.ReadString(1)
	return param[:len(param)-1]
}

func (packet *Packet) beep() string {
	return fmt.Sprintf("[=Beep!=] %s sent you a beep!", packet.Buffer.String())
}
