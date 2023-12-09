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
	rpcch 		chan RPC	
	mu	sync.Mutex
	Peer	map[net.Addr]Peer
}
//Consume implements the Transport interface, which will return the channel 
//To read RPCs received from other peer
func (s *server) Consume() <- chan RPC{
	return rpcch
}
//Tcp Transport options

type TcpTransportOpts struct{
	listenAddr string
	
	shakehands	HandshakeFunc 
	decoder		Decode
	OnPeer	func(Peer) error
	
}
func (s *server) NewServer(addr string,opts TcpTransportOpts) *server{
	return &server{
		TcpTransportOpts: opts,	
//		shakehands: func(any) error {return nil},
		listenAddr: addr,	
		rpcch: make(chan RPC),
}
}
//Close implements the Peer interface
func (p *TCPPeer) Close() error{
	return p.conn.Close
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
	rpc:=RPC{}
	defer func(){
		var err error
		fmt.Println("Dropping peer Connection %s",err)
		accept.Close()
	}()
	if s.OnPeer != nil{
		if err:=s.OnPeer(peer);err!=nil {
			return
		}
	}

	for{
	if err:=s.decoder.Decode(conn,&rpc);err!=nil{
	lencodeError++
	fmt.Printf("TCP error : %s\n",err)
	if(lencodeError==5){
	fmt.Println("Too much errors")}
	
	}
	}
	rpc.From=accept.RemoteAddr()
	s.rpcch<- rpc
	fmt.Println("RPC:",rpc)
}