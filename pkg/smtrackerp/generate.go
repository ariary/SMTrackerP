package smtrackerp

import (
	"bytes"
	"os"
	"text/template"

	"github.com/ariary/go-utils/pkg/check"
)

// GeneratePayload: generate the HTML transparent image and save it in a file
func GeneratePayload(cfg *Config) {
	// generate template
	payloadTpl := `
<html>
<body>
	<img src="{{ .Url}}/?target={{ .Target}}" opacity=0>
</body>
</html>
`
	t, err := template.New("payload").Parse(payloadTpl)
	check.CheckAndExit(err, "❌ Failed loading payload html template")
	data := struct {
		Url    string
		Target string
	}{
		Url:    cfg.Url,
		Target: cfg.Target,
	}

	//save file
	f, err := os.Create("payload.html")
	check.CheckAndExit(err, "❌ Failed creating file")

	defer f.Close()

	check.CheckAndExit(t.Execute(f, data), "❌ Failed writing html in file")
}

// GetTracker: Return the img tag used for tracking purpose
func GetTracker(url string, target string) (tracker string, err error) {
	// generate template
	trackerTpl := `
<img src="{{ .Url}}/?target={{ .Target}}" opacity=0>
`
	t, err := template.New("payload").Parse(trackerTpl)
	check.CheckAndExit(err, "❌ Failed loading payload html template")
	data := struct {
		Url    string
		Target string
	}{
		Url:    url,
		Target: target,
	}
	var w bytes.Buffer
	err = t.Execute(&w, data)
	tracker = w.String()

	return tracker, err
}
