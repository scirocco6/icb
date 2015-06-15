package icb

import (
    "fmt"
    "time"
    "strconv"
)

type packetMethod func()(string)

func (packet *IcbPacket) Decode ()(string) {    
    kind, _ := packet.Buffer.ReadByte()
    
    switch string(kind) {
        case "1": {return packet.serverID()}
        case "a": {return packet.loginResult()}
        case "b": {return packet.publicMessage()}
        case "c": {return packet.privateMessage()}
        case "d": {return packet.statusMessage()}
        case "e": {return packet.statusMessage()} 
        case "f": {return packet.importantMessage()}
        case "g": {return packet.serverExit()}
        case "i": {return packet.commandOutput()}
        case "k": {return packet.beep()}
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
    return "[!Exit!] Warning the server is exiting"
}

func (packet *IcbPacket) commandOutput ()(string) {
    var kind []byte = make([]byte, 2)
    packet.Buffer.Read(kind)
    
    switch(string(kind)) { // generic command output
        case "co": {
            return packet.Buffer.String()
        }
           
        case "ec": {  // end of output terminator in theory there should never be any actual message from this
            return packet.Buffer.String()
        }

        case "wh": { // output the who header
          return "   Nickname          Idle       Sign-On        Account"
        }
        
        case "wl": { // item in a who listing
            return packet.whoItem()
        }        
    }
    return ""
}

func (packet *IcbPacket) whoItem ()(string) {
        // kind, _  := packet.nextParameter()
        _      = packet.nextParameter()
        flag  := packet.nextParameter()
        nick  := packet.nextParameter()
        idle  := packet.nextParameter()
        // respc, _ := packet.nextParameter()
        _      = packet.nextParameter()
        _      = packet.nextParameter()
        user  := packet.nextParameter()
        host  := packet.nextParameter()
        
        account := fmt.Sprintf("%s@%s", user, host)
        if flag == "m" {
            flag = "*"
        }
        
        i, _ := strconv.Atoi(idle)        
        idleTime := time.Duration(i)*time.Second

        return fmt.Sprintf("%s %s %s %s", flag, nick, idleTime.String(), account)
}

func (packet *IcbPacket) nextParameter ()(string) {
    param, _  := packet.Buffer.ReadString(1)
    return param[:len(param) - 1]
}

func (packet *IcbPacket) beep ()(string) {
    return fmt.Sprintf("[=Beep!=] %s sent you a beep!", packet.Buffer.String())
}

