package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var (
		mutex   sync.Mutex
		started = time.Now()
	)

	router := http.NewServeMux()
	router.HandleFunc("/stopwatch/reset", func(http.ResponseWriter, *http.Request) {
		mutex.Lock()
		started = time.Now()
		mutex.Unlock()
	})
	router.HandleFunc("/stopwatch", func(response http.ResponseWriter, _ *http.Request) {
		mutex.Lock()
		io.WriteString(response, time.Since(started).Round(time.Second).String())
		mutex.Unlock()
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
