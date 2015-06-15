package icb

func (packet *IcbPacket) packParameters (parameters []string) {
    for _, parameter := range parameters {        
		packet.Buffer.WriteString(parameter)
		packet.Buffer.WriteByte(1)
	}
	packet.Buffer.Truncate(packet.Buffer.Len() - 1) // remove the excess seperator
}

func CreatePacket(kind string, parameters ...string)(IcbPacket) { 
	var packet IcbPacket
	
	switch kind {
	    case "login":   {packet.Init([]byte{'a'}, parameters)}
	    
	    case "beep":    {packet.Init([]byte{'h'}, []string{"beep", parameters[0]})}
	    
	    case "public":  {packet.Init([]byte{'b'}, parameters)}
	    
	    case "private": {packet.Init([]byte{'h'}, []string{"m", parameters[0] + " " + parameters[1]})}
	    
	    case "join": {packet.Init([]byte{'h'}, []string{"g", parameters[0]})}
	    
	    case "global_who": {packet.Init([]byte{'h'}, []string{"w\001"})}
	    
	    case "local_who": {packet.Init([]byte{'h'}, append([]string{"w"}, parameters...))}
	    
	    case "nop":   {packet.Init([]byte{'n'}, parameters)}
	}
	
	return packet
}