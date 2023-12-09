package p2p
import "net"
//Message any data transfered between transport nodes
type RPC struct{
	Payload []byte
	from net.Addr
}