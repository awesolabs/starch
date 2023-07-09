package starch

func (t CtxParam) Render(c Context) error {
	param := c.GetParam(t[0])
	return Text{param}.Render(c)
}
