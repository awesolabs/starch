package starch

func GetVar[T comparable](c Context, key string, defaultv ...T) T {
	var nilt T
	v := c.Var(key)
	if v != nil {
		return v.(T)
	}
	for _, dv := range defaultv {
		if dv != nilt {
			return dv
		}
	}
	return nilt
}
