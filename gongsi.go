package ngsi

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

type Gongsi struct {
	Transport string
	Host      string
	Port      int
	Encoding  string
	client    *http.Client
	baseUrl   string
}

func (g *Gongsi) Init() {
	g.Encoding = "application/json"
	g.SetSSL(false)
	g.client = &http.Client{}
	g.baseUrl = g.Transport + "://" + g.Host + ":" + strconv.Itoa(g.Port)
}

func (g *Gongsi) SetSSL(useSSL bool) {
	if useSSL {
		g.Transport = "https"
	} else {
		g.Transport = "http"
	}
}

func (g *Gongsi) get(resource string) ([]byte, error) {
	req, err := http.NewRequest("GET", g.baseUrl+resource, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := g.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
