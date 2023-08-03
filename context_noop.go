package starch

type NoopContext struct{}

func NewNoopContext() *NoopContext {
	return &NoopContext{}
}

func (t *NoopContext) Var(key string) any {
	return nil
}

func (t *NoopContext) SetVar(key string, value any) {
}

func (t *NoopContext) Read(buff []byte) (int, error) {
	return 0, nil
}

func (t *NoopContext) Write(buff []byte) (int, error) {
	return 0, nil
}

func (t *NoopContext) WriteHeader(key string, value string) error {
	return nil
}

func (t *NoopContext) WriteStatus(code int) error {
	return nil
}

func (t *NoopContext) WriteString(s string, args ...any) error {
	return nil
}

func (t *NoopContext) Render(c Component) error {
	return c.Render(t)
}

func (t *NoopContext) Redirect(string) error {
	return nil
}

func (t *NoopContext) Param(string) string {
	return ""
}
