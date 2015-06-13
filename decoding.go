package icb

import (
    "fmt"
    "os"
)

type packetMethod func()(string)

func (packet *IcbPacket) Decode ()(string) {    
    kind, _ := packet.Buffer.ReadByte()

    parser := map[string]packetMethod {
        "1": packet.serverID,
        "a": packet.loginResult,
        "b": packet.publicMessage,
        "c": packet.privateMessage,
        "d": packet.statusMessage,
        "e": packet.errorMessage,
        "f": packet.importantMessage,
        "g": packet.serverExit,
        "i": packet.commandOutput,
        "k": packet.beep,
    }
    
    message := parser[string(kind)]()
    if (message != "") {
        return message
    }
    return ""
}

func (packet *IcbPacket) serverID ()(string) {
    return fmt.Sprintf("[=Login=] %s", packet.Buffer.String())
}

func (packet *IcbPacket) loginResult ()(string) {
    return fmt.Sprintf("[=Login=] %s", packet.Buffer.String())
}

func (packet *IcbPacket) publicMessage ()(string) {
    from, _ := packet.Buffer.ReadString(1)
    text    := packet.Buffer.String()
    
    return fmt.Sprintf("<%s> %s", from, text)
}

func (packet *IcbPacket) privateMessage ()(string) {
    from, _ := packet.Buffer.ReadString(1)
    text    := packet.Buffer.String()
    
    return fmt.Sprintf("<*%s*> %s", from, text)
}

func (packet *IcbPacket) statusMessage ()(string) {
    category, _ := packet.Buffer.ReadString(1)
    text        := packet.Buffer.String()
    
    return fmt.Sprintf("[=%s=] %s", category, text)
}

func (packet *IcbPacket) errorMessage ()(string) {
    return fmt.Sprintf("[=Error=] %s", packet.Buffer.String())
}

func (packet *IcbPacket) importantMessage ()(string) {
    category, _ := packet.Buffer.ReadString(1)
    text        := packet.Buffer.String()
    
    return fmt.Sprintf("[!%s!] %s", category, text)
}

func (packet *IcbPacket) serverExit ()(string) {
    os.Exit(1) // TODO: need to refactor to utilize clean exit somehow most likely send signal to self and catch from sig handler
    return""
}

func (packet *IcbPacket) commandOutput ()(string) {
    return fmt.Sprintf("[=result=] %s", packet.Buffer.String())
}

func (packet *IcbPacket) beep ()(string) {
    return fmt.Sprintf("[=Beep!=] %s sent you a beep!", packet.Buffer.String())
}

