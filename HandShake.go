package p2p
import "errors"
/*type Handshaker interface{
	Handshake() error
}
type DefaultHandshaker struct{
	
}*/
//handshake error 
var HandshakeError error = errors.New("Error while doing handshake")
type HandshakeFunc func(Peer) error
func NOPHandshakeFunc(Peer) error {}