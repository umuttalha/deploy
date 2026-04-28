package tagging

import (
	"reflect"
	"testing"
)

func TestStackTags(t *testing.T) {
	got := StackTags("demo")
	want := map[string]string{
		"managed-by": "deploy",
		"stack":      "demo",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("StackTags(demo) = %v, want %v", got, want)
	}
}

func TestParseStack_Found(t *testing.T) {
	tags := map[string]string{"managed-by": "deploy", "stack": "demo", "extra": "x"}
	name, ok := ParseStack(tags)
	if !ok || name != "demo" {
		t.Errorf("ParseStack = (%q, %v), want (demo, true)", name, ok)
	}
}

func TestParseStack_NotManaged(t *testing.T) {
	tags := map[string]string{"stack": "demo"}
	_, ok := ParseStack(tags)
	if ok {
		t.Errorf("ParseStack(no managed-by) ok=true, want false")
	}
}

func TestParseStack_NoStack(t *testing.T) {
	tags := map[string]string{"managed-by": "deploy"}
	_, ok := ParseStack(tags)
	if ok {
		t.Errorf("ParseStack(no stack) ok=true, want false")
	}
}
