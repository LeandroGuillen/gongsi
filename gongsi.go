package ngsi

import (
	"github.com/franela/goreq"
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

func (g *Gongsi) get(resource string) (*goreq.Response, error) {
	res, err := goreq.Request{
		Method: "GET",
		Uri:    g.baseUrl + resource,
		Accept: "application/json",
	}.Do()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *Gongsi) post(resource string, payload interface{}) (*goreq.Response, error) {
	res, err := goreq.Request{
		Method:      "POST",
		Uri:         g.baseUrl + resource,
		Accept:      "application/json",
		ContentType: "application/json",
		Body:        payload,
	}.Do()

	if err != nil {
		return nil, err
	}

	return res, nil
}
