package p2p
import ("encoding/gob"
	"io")
type GobDecoder struct{}
type Decoder interface{
	decode(io.Reader,*RPC) error
}
func (dec GobDecoder) Decoder(r io.Reader,msg *RPC) error{
	return gob.NewDecoder(r).Decode(msg)
}
type NopDecode struct{}
func (dec NopDecode) Decode(r io.Reader,msg *RPC) error{
	buf := make([]byte,1024)
	n,err:=r.Read(buf)
	if err!=nil{
	return err}	
	msg.Payload = buf[:n]
	return nil
}
