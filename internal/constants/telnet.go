package constants

// Telnet commands
var (
	EchoOff = []byte{255, 251, 1} //"\xFF\xFB\x01\xFF\xFB\x03"
	EchoOn  = []byte{255, 252, 1}
	EOL     = "\r\n"
)
