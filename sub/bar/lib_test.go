package bar

import (
	"log"
	"testing"
)

func TestBar(t *testing.T) {
	result := Bar()
	expected := "Foo, Bar"
	if result != expected {
		log.Fatal("test failed")
	}
}
