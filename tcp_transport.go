package p2p
	
import ("log"
	"net"
	"fmt")
//represent the remote node
type TCPPeer struct{
	conn	net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn,outbound bool) *TCPPeer{
	return &TCPPeer{
		conn:conn,
		outbound:outbound,
	}
}
type server struct{
	TcpTransportOpts
	
	listener	net.Listener
	mu	sync.Mutex
	Peer	map[net.Addr]Peer
}
//Tcp Transport options

type TcpTransportOpts struct{
	listenAddr string
	
	shakehands	HandshakeFunc 
	decoder		Decode
	
}
func (s *server) NewServer(addr string,opts TcpTransportOpts) *server{
	return &server{
		TcpTransportOpts: opts,	
//		shakehands: func(any) error {return nil},
		listenAddr: addr,	
}
}

func ListenAndAccept(){
	s.Listener,err:=net.Listener("tcp",s.ListenAddr)
	if err!=nil{
		log.Fatal(err)
	}
	go s.acceptLoop(s.Listener)

}
func (s *server) acceptLoop(conn net.Conn) {
	for{
	accept,err := conn.Accept()
	if err!=nil{
		fmt.Println("error while accepting new connection")
	}
	go func(accept net.Conn){s.Handle(accept)}(accept)
	}
}
func (s *server) Handler(accept net.Conn){
	peer:=NewTCPPeer(accept,false)
	if err:=s.shakehands(peer)!=nil{
		conn.Close()
		fmt.Pritnf("TCP Handshake error : %s\n",err)
	}
	var lencodeError int = 0
	//Read Loop
	msg:=&Temp()
	for{
	if err:=s.decoder.Decode(conn,msg);err!=nil{
	lencodeError++
	fmt.Printf("TCP error : %s\n",err)
	if(lencodeError==5){
	fmt.Println("Too much errors")}
	
	}
	}
	fmt.Println("Hello world!")
}