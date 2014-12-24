package ngsi

import (
	"io/ioutil"
	"net/http"
)

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

func (g *Gongsi) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// ...

	return body, err
}
