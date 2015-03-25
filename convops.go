package gongsi

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

func (g *Gongsi) ConvQueryContextAttribute(entity, attribute string) (string, error) {
	res, err := g.get("/v1/contextEntities/" + entity + "/attributes/" + attribute)
	if err != nil {
		return "", err
	}

	out, err := res.Body.ToString()
	if err != nil {
		return "", err
	}

	return out, nil
}