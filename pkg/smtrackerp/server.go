package smtrackerp

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ariary/go-utils/pkg/color"
	ua "github.com/mileusna/useragent"
)

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

		ip := strings.Split(r.RemoteAddr, ":")[0]
		if ip == "127.0.0.1" {
			ip = r.Header.Get("X-Forwarded-For")
		}
		fmt.Println(color.Cyan("📍 IP address: ", ip))

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
			fmt.Println(color.Cyan("📬 SMTP Client: Thunderbird"))
		}

		referer := r.Referer()
		if referer != "" {
			fmt.Println(color.Cyan("📨 Referer: ", referer))
		}
		fmt.Println()
	})

}

// Serve: wait for request indicating that user read the mail
func Serve(cfg *Config) {
	http.Handle("/", TrackHandler(cfg))

	fmt.Println(color.Bold("🚀 Serve at:"), color.Italic(cfg.Url))
	fmt.Println()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
