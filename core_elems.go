package starch

type Elems []Component

func (t Elems) Render(c Context) error {
	for _, item := range t {
		if err := item.Render(c); err != nil {
			return err
		}
	}
	return nil
}
