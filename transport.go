package p2p

type transport interface{
	ListenAndAccept() error
}


type Peer interface{
	Close() error
	Consume() <- chan RPC
}
