package main

type Config interface {
	Map() map[string]interface{}
}

type CouterConfig struct {
	data []byte
}

func (p *CouterConfig) Map() interface{} {
	return nil
}
