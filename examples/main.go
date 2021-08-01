package main

import (
	"io"
	"os"
	"strconv"
	"time"

	"github.com/syrinsec/termfx"
)

func main() {

	registry := termfx.New()
	registry.RegisterVariable("username", "root")

	registry.RegisterFunction("sleep", func(session io.Writer, args string) (int, error) {

		sleep, err := strconv.Atoi(args)
		if err != nil {
			return 0, err
		}

		time.Sleep(time.Millisecond * time.Duration(sleep))
		return 0, nil
	})

	registry.Execute(`
hello [[$username]]

start
[[sleep(500)]]
that was 500ms
[[sleep(1000)]]
and that last one was 1 second

	`, os.Stdout)
}
