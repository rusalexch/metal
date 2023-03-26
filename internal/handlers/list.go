package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/utils"
)

type item struct {
	Name  string
	Value string
}

type res struct {
	Title string
	Items []item
}

func (h *Handlers) list(w http.ResponseWriter, r *http.Request) {
	metrics := h.services.MetricsService.List()
	fmt.Println(metrics)
	res := res{
		Title: "Метрики",
		Items: make([]item, 0, len(metrics)),
	}

	for _, v := range metrics {
		switch v.Type {
		case app.Counter:
			res.Items = append(res.Items, item{
				Name:  v.ID,
				Value: utils.Int64ToStr(*v.Delta),
			})
		case app.Gauge:
			res.Items = append(res.Items, item{
				Name:  v.ID,
				Value: utils.Float64ToStr(*v.Value),
			})
		}
	}

	t, err := template.New("metrics").Parse(tmpl)
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(w, res)
	if err != nil {
		log.Println(err)
	}
}

var tmpl = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<ul>
			{{range .Items}}<li><b>{{ .Name }}:</b> {{ .Value }}</li>{{else}}<div><strong>no metrics</strong></div>{{end}}
		</ul>
	</body>
</html>`
