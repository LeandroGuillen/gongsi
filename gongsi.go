package ngsi

type Gongsi struct {
	Host     string
	Port     int
	Encoding string
}

func (g *Gongsi) Init(host string, port int, encoding string) {
	g.Host = host
	g.Port = port
	g.Encoding = encoding
}
