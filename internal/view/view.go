package view

import (
	"zenith/pkg/jsonutil"

	"github.com/leonelquinteros/gotext"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type View struct {
	Type    string
	RawJson *gjson.Result
	Mo      *gotext.Mo
}

func (v *View) Render() string {
	tp, _ := jsonutil.GetString("type", v.RawJson, "")
	var obj Type
	switch tp {
	case "MONSTER":
		obj = &Monster{}
	default:
		log.Warnf("type: %s is not supported to render", tp)
		return ""
	}
	obj.Bind(v.RawJson, v.Mo)
	switch v.Type {
	case "cli":
		return obj.CliView()
	case "json":
		return obj.JsonView()
	default:
		return ""
	}
}