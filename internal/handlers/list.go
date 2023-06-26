package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/rusalexch/metal/internal/app"
	"github.com/rusalexch/metal/internal/utils"
)

// Структура метрики для шаблона html
type metric struct {
	// Name - наименование метрики
	Name string
	// Value - значение метрики
	Value string
}

// data - структура данных для шаблона html
type data struct {
	// Title - заголовок страницы
	Title string
	// Metrics - список метрик
	Metrics []metric
}

// list - хэндлер вывода списка метрик
func (h *Handlers) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	metrics, err := h.storage.List(ctx)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := data{
		Title: "Метрики",
		Metrics: make([]metric, 0, len(metrics)),
	}

	for _, v := range metrics {
		switch v.Type {
		case app.Counter:
			res.Metrics = append(res.Metrics, metric{
				Name:  v.ID,
				Value: utils.Int64ToStr(*v.Delta),
			})
		case app.Gauge:
			res.Metrics = append(res.Metrics, metric{
				Name:  v.ID,
				Value: utils.Float64ToStr(*v.Value),
			})
		}
	}

	t, err := template.New("metrics").Parse(tmpl)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add(contentType, text)
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
			{{range .Metrics}}<li><b>{{ .Name }}:</b> {{ .Value }}</li>{{else}}<div><strong>no metrics</strong></div>{{end}}
		</ul>
	</body>
</html>`
