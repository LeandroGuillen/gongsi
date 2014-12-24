package ngsi

func (g *Gongsi) Version() (string, error) {
	resp, err := g.get("/version")
	return string(resp[:]), err
}
