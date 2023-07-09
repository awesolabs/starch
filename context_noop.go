package starch

func NewNoopContext() *NoopContext {
	return &NoopContext{}
}

func (t *NoopContext) GetVar(key string) any {
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

func (t *NoopContext) WriteHeader(key string, value string) {
}

func (t *NoopContext) WriteStatus(code int) {
}

func (t *NoopContext) WriteString(s string, args ...any) error {
	return nil
}

func (t *NoopContext) Next(c Component) error {
	return c.Render(t)
}

func (t *NoopContext) Redirect(string) {}

func (t *NoopContext) GetParam(string) string {
	return ""
}
