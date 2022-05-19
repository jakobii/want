# want
Go test utility 



`foo_test.go`
```go
import "github.com/jakobii/want"

func TestFoo(t *testing.T) {
    want.Eq(t, "foo", "bar")
}
```

```bash
go test ./...

# output
--- FAIL: TestFoo (0.00s)
    /foo_test.go:4: 
        want: foo
        got: bar
```
