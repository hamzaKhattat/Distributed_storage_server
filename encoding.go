package p2p
import ("encoding/gob"
	"io")
type Decoder interface{
	decode(io.Reader,any) error
}
func GobDecoder(r io.Reader,v any) error{
	return gob.NewDecoder(r).Decode(v)
}