package gotest

import "testing"

func TestFormat(t *testing.T) {
	reformatted := Format(rawOutput)
	t.Log("\n\n"+reformatted+"\n\n")
}

const formattedOutput = `
go version go1.15.6 darwin/amd64
go mod tidy
go fmt ./...
go test -timeout=1s -short -cover ./...
ok   project/code/app/awefads                       100.0%                    (cached)
?    project/code/app/asdf                                            [no tests files]
ok   project/code/app/internal/asdfad               85.5%                     (cached)
ok   project/code/app/internal/aefasdfcxa           100.0%                    (cached)
ok   project/code/app/internal/asdfefawe            100.0%                    (cached)
ok   project/code/app/internal/asdfefaxfd           100.0%                    (cached)
ok   project/code/app/internal/asdfdd               100.0%                    (cached)
ok   project/code/app/internal/asdfdfd              98.9%                     (cached)
ok   project/code/app/internal/asdfdf               100.0%                    (cached)
?    project/code/app/asdfa                                           [no tests files]
?    project/code/eiei/internal/adfd                                  [no tests files]
?    project/code/eiei/internal/asdfd/werdfd                          [no tests files]
ok   project/code/eiei/internal/efazsd              86.8%                     (cached)
ok   project/code/eiei/internal/asdfdfd/moawerrdels 97.8%                     (cached)
--- FAIL: TestBLAH (0.00s)
	--- FAIL: TestBLAH (0.00s)
		test_case.go:71: Test definition:
			/blah_test.go:38
		fixture.go:43: 
			Expected: 201
			Actual:   200
FAIL
coverage: 87.4% of statements
FAIL project/code/eiei/internal/feefzsdffe                                      0.355s
ok   project/code/eiei/internal/ef3raw3r3           66.7%                     (cached)
ok   project/code/eiei/internal/cvzxrfg             82.2%                     (cached)
ok   project/code/eiei/internal/radsfd              75.0%                     (cached)
?    project/code/eiei/awer3adf                                       [no tests files]
?    project/code/adsfads                                             [no tests files]
ok   project/code/asdfasd/awef                      [no tests to run]         (cached)
?    project/code/asdfa/asdfe4/aw3r3                                  [no tests files]
?    project/main/zxcvzx                                              [no tests files]
?    project/main/asdf                                                [no tests files]
ok   project/main/wefa/w23fa                        61.0%                     (cached)
?    project/main/zxcv/asdfasd                                        [no tests files]
FAIL
make: *** [test] Error 1
`

const rawOutput = `go version go1.15.6 darwin/amd64
go mod tidy
go fmt ./...
go test -timeout=1s -short -cover ./...
ok  	project/code/app/awefads	(cached)	coverage: 100.0% of statements
?   	project/code/app/asdf	[no test files]
ok  	project/code/app/internal/asdfad	(cached)	coverage: 85.5% of statements
ok  	project/code/app/internal/aefasdfcxa	(cached)	coverage: 100.0% of statements
ok  	project/code/app/internal/asdfefawe	(cached)	coverage: 100.0% of statements
ok  	project/code/app/internal/asdfefaxfd	(cached)	coverage: 100.0% of statements
ok  	project/code/app/internal/asdfdd	(cached)	coverage: 100.0% of statements
ok  	project/code/app/internal/asdfdfd	(cached)	coverage: 98.9% of statements
ok  	project/code/app/internal/asdfdf	(cached)	coverage: 100.0% of statements
?   	project/code/app/asdfa	[no test files]
?   	project/code/eiei/internal/adfd	[no test files]
?   	project/code/eiei/internal/asdfd/werdfd	[no test files]
ok  	project/code/eiei/internal/efazsd	(cached)	coverage: 86.8% of statements
ok  	project/code/eiei/internal/asdfdfd/moawerrdels	(cached)	coverage: 97.8% of statements
--- FAIL: TestBLAH (0.00s)
    --- FAIL: TestBLAH (0.00s)
        test_case.go:71: Test definition:
            /blah_test.go:38
        fixture.go:43: 
            Expected: 201
            Actual:   200
FAIL
coverage: 87.4% of statements
FAIL	project/code/eiei/internal/feefzsdffe	0.355s
ok  	project/code/eiei/internal/ef3raw3r3	(cached)	coverage: 66.7% of statements
ok  	project/code/eiei/internal/cvzxrfg	(cached)	coverage: 82.2% of statements
ok  	project/code/eiei/internal/radsfd	(cached)	coverage: 75.0% of statements
?   	project/code/eiei/awer3adf	[no test files]
?   	project/code/adsfads	[no test files]
ok  	project/code/asdfasd/awef	(cached)	[no tests to run]
?   	project/code/asdfa/asdfe4/aw3r3	[no test files]
?   	project/main/zxcvzx	[no test files]
?   	project/main/asdf	[no test files]
ok  	project/main/wefa/w23fa	(cached)	coverage: 61.0% of statements
?   	project/main/zxcv/asdfasd	[no test files]
FAIL
make: *** [test] Error 1`
