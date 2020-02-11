package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Unleash/unleash-client-go/v3"
	"github.com/Unleash/unleash-client-go/v3/context"
)

var (
	indexTpl   *template.Template
	newPageTpl *template.Template
	unleashUrl = os.Getenv("UNLEASH_URL")
)

type templateData struct {
	Ctx context.Context
}

func init() {
	// Create some helper functions for use when rendering
	var helpers = template.FuncMap{
		"isEnabled": func(feature string, ctx context.Context) bool {
			enabled := unleash.IsEnabled(feature, unleash.WithContext(ctx))
			println(feature, ctx.RemoteAddress, enabled)
			return enabled
		},
	}

	// Compile the HTML templates
	indexTpl = template.Must(
		template.
			New("index").
			Funcs(helpers).
			ParseFiles("html/index.html", "html/features.html", "html/base.html"),
	)
}

func contextFromRequest(r *http.Request) context.Context {
	return context.Context{
		RemoteAddress: r.RemoteAddr,
		Properties: map[string]string{
			"userAgent": r.UserAgent(),
		},
	}
}

// Index is the handler that writes the rendered HTML to the response or an error.
func Index(w http.ResponseWriter, r *http.Request) {
	ctx := contextFromRequest(r)

	isEnabled := unleash.IsEnabled(
		"RedirectToNewPage",
		unleash.WithContext(ctx),
	)

	if isEnabled {
		http.Redirect(w, r, "https://pkg.go.dev/github.com/Unleash/unleash-client-go/v3", 307)
		return
	}

	data := templateData{ctx}

	var helpers = template.FuncMap{
		"isEnabled": func(feature string, ctx context.Context) bool {
			enabled := unleash.IsEnabled(feature, unleash.WithContext(ctx))
			println(feature, ctx.RemoteAddress, enabled)
			return enabled
		},
	}

	// Compile the HTML templates
	indexTpl = template.Must(
		template.
			New("index").
			Funcs(helpers).
			ParseFiles("html/index.html", "html/features.html", "html/base.html"),
	)

	if err := indexTpl.ExecuteTemplate(w, "base.html", data); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

}

func main() {
	// Load out custom strategy
	userAgent := &UserAgentStrategy{}

	// Initialize the unleash client
	err := unleash.Initialize(
		unleash.WithUrl(unleashUrl),
		unleash.WithAppName("demo"),
		unleash.WithRefreshInterval(5*time.Second),
		unleash.WithMetricsInterval(1*time.Second),
		unleash.WithStrategies(userAgent),
		unleash.WithListener(unleash.DebugListener{}),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Specify the routes and handlers
	http.HandleFunc("/", Index)

	// Start the server
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
