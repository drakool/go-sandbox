package server

import "testing"
import "time"
import "fmt"

func TestHello(t *testing.T) {
    want := "Hello, world."

    svr := New(Config{"localhost", 1234, time.Second, 10})
    fmt.Println(svr.Hello())
    if got := svr.Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
