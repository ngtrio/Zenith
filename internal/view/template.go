package view

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/leonelquinteros/gotext"
	"github.com/tidwall/gjson"
	"html/template"
	"io"
	"zenith/internal/i18n"
)

type Template struct {
	template *template.Template
}

func NewTemplate() *Template {
	t := &Template{
		template: template.Must(template.New("").Funcs(template.FuncMap{
			"parseFgColor": ParseFgColor,
			"parseBgColor": ParseBgColor,
			"genMap":       GenMap,
			"tranUI":       TranUI,
			"html":         Html,
			"getJsonField": GetJsonField,
		}).ParseGlob("web/template/*.html")),
	}

	return t
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func ParseFgColor(color string) string {
	var l Color
	l.Load(color)

	return l.FgColor
}

func ParseBgColor(color string) string {
	var l Color
	l.Load(color)
	return l.BgColor
}

func GenMap(p ...any) (map[string]any, error) {
	if len(p)%2 != 0 {
		return nil, fmt.Errorf("param error")
	}

	m := make(map[string]any)

	for i := 0; i < len(p); i += 2 {
		if k, ok := p[i].(string); !ok {
			return nil, fmt.Errorf("param error")
		} else {
			v := p[i+1]
			m[k] = v
		}
	}

	return m, nil
}

func TranUI(word string, po *gotext.Po) string {
	return i18n.TranUI(word, po)
}

func Html(str string) template.HTML {
	return template.HTML(str)
}

func GetJsonField(json *gjson.Result, field string) string {
	return json.Get(field).String()
}