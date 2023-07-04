package starch

type Param []string

func (t Param) Render(c Context) error {
	param := c.GetParam(t[0])
	return c.Next(Text{param})
}
