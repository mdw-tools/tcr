package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/stopwatch/reset", func(http.ResponseWriter, *http.Request) {
		started = time.Now()
	})
	router.HandleFunc("/stopwatch", func(response http.ResponseWriter, _ *http.Request) {
		io.WriteString(response, elapsed(time.Since(started)))
	})
	router.HandleFunc("/", func(response http.ResponseWriter, _ *http.Request) {
		io.WriteString(response, uiHTML)
	})

	address := ":7890"
	log.Printf("[INFO] Listening for web traffic on %s.", address)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}

var started = time.Now()

func elapsed(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
const uiHTML = `<html>
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
<body style="font-family: monospace;">
</body>
</html>
`
