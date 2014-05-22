package ostential
import (
	"ostential/view"

	"log"
	"runtime"
	"net/http"
	"html/template"
)

type recovery struct{
	production bool
	logger     *log.Logger
	handler    http.Handler
}

func Recovering(production bool, logger *log.Logger, handler http.Handler) http.Handler {
	return &recovery{
		production: production,
		logger:     logger,
		handler:    handler,
	}
}

func RecoveringFunc(production bool, logger *log.Logger) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return Recovering(production, logger, handler)
	}
}

func (rc *recovery) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if E := recover(); E != nil {
			rc.recov(w, E)
		}
	}()
	rc.handler.ServeHTTP(w, r)
}

const panicstatuscode = http.StatusInternalServerError
var   panicstatustext = view.StatusLine(panicstatuscode)

func (rc *recovery) recov(w http.ResponseWriter, E interface{}) {
	w.WriteHeader(panicstatuscode) // NB

	var description string
	if err, ok := E.(error); ok {
		description = err.Error()
	}
	var stack string
	if !rc.production {
		sbuf := make([]byte, 4096 - len(panicstatustext) - len(description))
		size := runtime.Stack(sbuf, false)
		stack = string(sbuf[:size])
	}
	rctemplate.Execute(w, struct {
		Title, Description, Stack string
	}{
		Title:       panicstatustext,
		Description: description,
		Stack:       stack,
	})
}

var rctemplate = template.Must(template.New("recovery.html").Parse(`
<html>
<head><title>{{.Title}}</title></head>
<body bgcolor="white">
<center><h1>{{.Description}}</h1></center>
<hr><pre>{{.Stack}}</pre>
</body>
</html>
`))
