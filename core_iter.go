package starch

type Each[T any] struct {
	Items *[]T
	Then  func(T) Component
}

func (t Each[T]) Render(c Context) error {
	for _, item := range *t.Items {
		if err := t.Then(item).Render(c); err != nil {
			return err
		}
	}
	return nil
}
