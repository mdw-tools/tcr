package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	started = time.Now()
	times   []time.Duration
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/stopwatch/reset", func(http.ResponseWriter, *http.Request) {
		mutex.Lock()
		times = append(times, since(started))
		started = time.Now()
		mutex.Unlock()
	})
	router.HandleFunc("/stopwatch", func(response http.ResponseWriter, _ *http.Request) {
		mutex.Lock()
		io.WriteString(response, renderDurations(times))
		mutex.Unlock()
	})
	router.HandleFunc("/", func(response http.ResponseWriter, _ *http.Request) {
		io.WriteString(response, uiHTML)
	})

	address := "localhost:7890"
	log.Printf("[INFO] Listening for web traffic on %s.", address)
	go func() { _ = exec.Command("open", "http://"+address).Run() }()
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}

func since(then time.Time) time.Duration {
	return time.Since(then).Round(time.Second)
}

func renderDuration(duration time.Duration) string {
	return fmt.Sprintf("<li>%s</li>\n", duration.String())
}

func renderDurations(durations []time.Duration) string {
	var builder strings.Builder
	builder.WriteString("<ol>\n")
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
  <body style="font-family: monospace; color: #ffffff; background-color: #000000;">
  </body>
</html>
`
