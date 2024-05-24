package gotest

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/tcr/should"
)

func TestFormat(t *testing.T) {
	actual := strings.TrimSpace(Format(rawOutput))
	expected := strings.TrimSpace(formattedOutput)
	should.So(t, actual, should.Equal, expected)
}

const formattedOutput = `go mod tidy && go fmt ./...
go test -timeout=1s -race -covermode=atomic -short -count=1 ./...
?    root/app/wsx                        [no test files]
?    root/app/tgb                        [no test files]
-    root/app                       [missing test files]
?    root/rfv                            [no test files]
-    root/http                      [missing test files]
-    root/http/internal/yui         [missing test files]
-    root/main/api                  [missing test files]
-    root/app/hjk                   [missing test files]
?    root/oiuy                           [no test files]
-    root/tools/foo                 [missing test files]
ok   root/app/zxcv            85.5%               0.221s
--- FAIL: Test (0.01s)
    --- FAIL: Test/TestBlah (0.00s)
        handler_test.go:20: Test definition:
            /path/to/file_test.go:62
        file_test.go:20: 
            Expected: hi
            Actual:   bye
FAIL
coverage: 94.6% of statements
FAIL root/app/internal/bar                        0.141s
ok   root/app/internal/baz   100.0%               0.281s
ok   root/app/internal/boink  68.8%               0.391s
ok   root/storage/qwer         0.0%               0.515s
FAIL
`

const rawOutput = `go mod tidy && go fmt ./...
go test -timeout=1s -race -covermode=atomic -short -count=1 ./...
?   	root/app/wsx	[no test files]
?   	root/app/tgb	[no test files]
	root/app		coverage: 0.0% of statements
?   	root/rfv	[no test files]
	root/http		coverage: 0.0% of statements
	root/http/internal/yui		coverage: 0.0% of statements
	root/main/api		coverage: 0.0% of statements
	root/app/hjk		coverage: 0.0% of statements
?   	root/oiuy	[no test files]
	root/tools/foo		coverage: 0.0% of statements
ok  	root/app/zxcv	0.221s	coverage: 85.5% of statements
--- FAIL: Test (0.01s)
    --- FAIL: Test/TestBlah (0.00s)
        handler_test.go:20: Test definition:
            /path/to/file_test.go:62
        file_test.go:20: 
            Expected: hi
            Actual:   bye
FAIL
coverage: 94.6% of statements
FAIL	root/app/internal/bar	0.141s
ok  	root/app/internal/baz	0.281s	coverage: 100.0% of statements
ok  	root/app/internal/boink	0.391s	coverage: 68.8% of statements
ok  	root/storage/qwer	0.515s	coverage: 0.0% of statements
FAIL`
