package main

import (
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
		io.WriteString(response, elapsed())
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

func elapsed() string {
	return time.Since(started).Round(time.Second).String()
}

const uiHTML = `<html>
<head>
  <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
  <script type="text/javascript">
    jQuery(document).ready(function() {
      setInterval(function() {
        $.ajax("/stopwatch", {
          success: function(data) {
            $('body').html(data);
          }
        });
      }, 1000)
    });
  </script>
</head>
<body style="font-family: monospace;">
</body>
</html>
`
