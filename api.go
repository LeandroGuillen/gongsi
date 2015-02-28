package gongsi

func (g *Gongsi) Version() (string, error) {
	res, err := g.get("/version")
	if err != nil {
		return "", err
	}

	out, err := res.Body.ToString()
	if err != nil {
		return "", err
	}

	return out, nil
}

func (g *Gongsi) UpdateContext(elem ContextElement) (string, error) {

	ucr := &UpdateContextRequest{
		ContextElements: []ContextElement{elem},
		UpdateAction:    "APPEND",
	}

	res, err := g.post("/v1/updateContext", ucr)
	if err != nil {
		return "", err
	}

	out, err := res.Body.ToString()
	if err != nil {
		return "", err
	}

	return out, nil
}

func (g *Gongsi) QueryContext(elem ContextElement) (string, error) {

	qcr := &QueryContextRequest{
		Entities: []ContextElement{elem},
		// Attributes: []string{
		// 	"temperature",
		// },
	}

	res, err := g.post("/v1/queryContext", qcr)
	if err != nil {
		return "", err
	}

	out, err := res.Body.ToString()
	if err != nil {
		return "", err
	}

	return out, nil
}

func (g *Gongsi) ConvQueryContext(eid string) (string, error) {
	res, err := g.get("/v1/contextEntities/" + eid)
	if err != nil {
		return "", err
	}

	out, err := res.Body.ToString()
	if err != nil {
		return "", err
	}

	return out, nil
}
