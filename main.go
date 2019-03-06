package main

import (
	"github.com/Unleash/unleash-client-go/v3/context"
	"html/template"
	"log"
	"net/http"
	"github.com/Unleash/unleash-client-go/v3"
	"os"
	"time"
)

var (
	tpl *template.Template
	unleashUrl = os.Getenv("UNLEASH_URL")
)

func init() {
	// Initialize the unleash client
	err := unleash.Initialize(
		unleash.WithUrl(unleashUrl),
		unleash.WithAppName("demo"),
		unleash.WithRefreshInterval(5*time.Second),
		unleash.WithMetricsInterval(1*time.Second),
		unleash.WithListener(unleash.DebugListener{}),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Create some helper functions for use when rendering
	var helpers = template.FuncMap{
		"isEnabled": func(feature string, ctx context.Context) bool {
			enabled := unleash.IsEnabled(feature, unleash.WithContext(ctx))
			println(feature, ctx.RemoteAddress, enabled)
			return enabled
		},
	}

	// Compile the HTML template
	tpl = template.Must(
		template.
			New("index.html").
			Funcs(helpers).
			ParseFiles("features.html", "index.html"),
	)
}

// Index is the handler that writes the rendered HTML to the response or an error.
func Index(w http.ResponseWriter, r *http.Request) {
	ctx := context.Context{
		RemoteAddress: r.RemoteAddr,
	}


	data := struct{
		Ctx context.Context
	} {
		ctx,
	}

	if err := tpl.Execute(w, data); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}

func main() {
	http.HandleFunc("/", Index)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}