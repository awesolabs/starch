package starch

func ConvertSlice[TIn Component](in []TIn) (out []Component) {
	for _, i := range in {
		out = append(out, i)
	}
	return
}
