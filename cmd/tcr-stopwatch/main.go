package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"time"
)

var Version = "dev"

var (
	mutex   sync.Mutex
	started = time.Now()
	times   []time.Duration
)

func main() {
	signals := make(chan os.Signal, 10)
	signal.Notify(signals, os.Interrupt, os.Kill)
	go func(){
		for s := range signals {
			log.Printf("[INFO] Shutdown initiated... (received: %s)", s.String())
			os.Exit(0)
		}
	}()
	router := http.NewServeMux()
	router.HandleFunc("/stopwatch/reset", func(http.ResponseWriter, *http.Request) {
		mutex.Lock()
		times = append(times, since(started))
		started = time.Now()
		mutex.Unlock()
	})
	router.HandleFunc("/stopwatch", func(response http.ResponseWriter, _ *http.Request) {
		mutex.Lock()
		_, _ = io.WriteString(response, renderDurations(times))
		mutex.Unlock()
	})
	router.HandleFunc("/", func(response http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(response, uiHTML)
	})

	address := "localhost:7890"
	log.Printf("[INFO] TCR Stopwatch [%s] listening for web traffic on %s.", Version, address)
	go func() { _ = exec.Command("open", "http://"+address).Run() }()
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}

func since(then time.Time) time.Duration {
	return time.Since(then).Round(time.Second)
}

func renderDuration(duration time.Duration) string {
	return fmt.Sprintf(
		`<li><span style="%s">`+"%s</span></li>\n",
		stylize(duration),
		duration.String(),
	)
}

func stylize(duration time.Duration) string {
	seconds := int(duration.Seconds()) * 3
	var red byte = 31
	var green byte = 221
	const blue byte = 31
	for red <= 221 && seconds > 0 {
		red++
		seconds--
	}
	for green >= 31 && seconds > 0 {
		green--
		seconds--
	}
	return fmt.Sprintf("color: #%s", hexRGB(red, green, blue))
}

func hexRGB(red, green, blue byte) string {
	return hex.EncodeToString([]byte{red, green, blue})
}

func renderDurations(durations []time.Duration) string {
	var builder strings.Builder
	builder.WriteString(`<ol reversed">` + "\n")
	builder.WriteString(renderDuration(since(started)))
	for x := len(durations) - 1; x >= 0; x-- {
		builder.WriteString(renderDuration(durations[x]))
	}
	builder.WriteString("</ol>")
	return builder.String()
}

const uiHTML = `
<html>
  <head>
	<title>tcr-stopwatch</title>
    <script type="text/javascript">
      setInterval(function() {
        var opts = {method: 'GET', headers: {}};
        fetch('/stopwatch', opts).then(function (body) {
          body.text().then(function (data) {
            document.body.innerHTML = data; 
          });
        });
      }, 1000);
    </script>
  </head>
  <body style="font-family: Courier Prime Code, monospace; color: #dddddd; background-color: #000000;">
  </body>
</html>
`
