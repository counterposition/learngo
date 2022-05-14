package bar

import (
	"fmt"

	"github.com/counterposition/learngo/foo"
)

func Bar() string {
	result := foo.Foo()
	return fmt.Sprintf("%s, Bar", result)
}
