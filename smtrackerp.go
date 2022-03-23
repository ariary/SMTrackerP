package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/ariary/go-utils/pkg/check"
	"github.com/ariary/go-utils/pkg/color"
	ua "github.com/mileusna/useragent"
)

type Config struct {
	Url     string
	Target  string
	Verbose bool
}

func main() {
	var verbose bool
	url := flag.String("e", "", "Server external reachable ip")
	flag.BoolVar(&verbose, "v", false, "verbose mode (see HTTP request headers)")

	flag.Parse()
	target := flag.Arg(0)

	cfg := &Config{Url: *url, Target: target, Verbose: verbose}
	generatePayload(cfg)
	fmt.Println(color.Bold("✉️ Generate HTML payload:"), color.Italic("payload.html"))

	http.Handle("/", TrackHandler(cfg))

	fmt.Println(color.Bold("🚀 Serve at:"), color.Italic(cfg.Url))
	fmt.Println()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TrackHandler(cfg *Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "image/png")

		//Generate fake image (an 1 x 1 image) to serve it
		img := image.NewRGBA(image.Rect(0, 0, 1, 1))
		png.Encode(w, img) //less bytes than jpeg.Encode(w, img, nil)

		// track info
		if cfg.Verbose {
			headers := r.Header
			for k, v := range headers {
				fmt.Println(k, v)
			}
		}

		target := r.FormValue("target")
		if target != "" {
			fmt.Println(color.Evil("👁️ ", strings.ToUpper(target), " IS READING.."))
		} else {
			fmt.Println(color.Evil("👁️ SOMEONE IS READING.."))
		}

		now := time.Now()
		fmt.Println(color.Cyan("🕚 at ", now))
		//parse user-agent
		userAgent := r.UserAgent()
		ua := ua.Parse(userAgent)
		//browser
		fmt.Print(color.Cyan("🌐 browser: "))
		if ua.Name != "" {
			fmt.Println(color.Cyan(ua.Name))
		} else {
			fmt.Println(color.Cyan(color.Italic("not detected")))
		}
		//OS
		fmt.Print(color.Cyan("🤖 OS: "))
		if ua.OS != "" {
			fmt.Print(color.Cyan(ua.OS))
			if ua.OSVersion != "" {
				fmt.Print(color.Cyan("(", ua.OSVersion, ")"))
			}
			fmt.Println()
		} else {
			fmt.Println(color.Cyan(color.Italic("not detected")))
		}
		//hardware type
		fmt.Print(color.Cyan("🖥️ device type: "))
		if ua.Device != "" {
			fmt.Println(color.Cyan(ua.Device))
		} else {
			fmt.Println(color.Cyan(color.Italic("not detected")))
		}

		//More things to test..
		switch {
		case strings.Contains(userAgent, "Thunderbird"):
			fmt.Println(color.Cyan("📨 SMTP Client: Thunderbird"))
		}

		referer := r.Referer()
		if referer != "" {
			fmt.Println(color.Cyan("📨 Referer:", referer))
		}
		fmt.Println()
	})

}

// generatePayload: generate the HTML transparent image and save it in a file
func generatePayload(cfg *Config) {
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
