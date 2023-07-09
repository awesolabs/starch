package starch

type Each[T any] struct {
	Items *[]T
	Thenf func(T) Component
}

func (t Each[T]) Render(c Context) error {
	for _, item := range *t.Items {
		if err := t.Thenf(item).Render(c); err != nil {
			return err
		}
	}
	return nil
}
