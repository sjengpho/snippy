package snippy

import (
	"reflect"
	"testing"
)

func TestConfigDefaultConfig(t *testing.T) {
	want := reflect.TypeOf(Config{})
	got := reflect.TypeOf(DefaultConfig())
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
