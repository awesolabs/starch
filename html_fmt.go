package starch

import (
	"io"

	"github.com/yosssi/gohtml"
)

type HTMLFormatter struct {
	ghf *gohtml.Writer
}

func NewHTMLFormatter(sink io.Writer) *HTMLFormatter {
	ghf := gohtml.NewWriter(sink)
	return &HTMLFormatter{
		ghf: ghf,
	}
}

func (t *HTMLFormatter) Format(data []byte) (n int, err error) {
	return t.ghf.Write(data)
}
