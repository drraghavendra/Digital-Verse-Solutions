package blockchain

type Network interface {
	Mint(metadataURI string) (*Resp, error)
}
