package starch

func (t CtxParam) Render(c Context) error {
	param := c.Param(t[0])
	return Text{param}.Render(c)
}
