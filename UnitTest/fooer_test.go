package unittest

import (
	"testing"
)

func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s", result, "Foo")
	}
}

func TestFooer2(t *testing.T) {
	input := 3
	result := Fooer(3)
	t.Logf("The input was %d", input)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, was: %s", result, "Foo")
	}
	t.Fatalf("Stop the test now, we have seen enough")
	t.Error("This won't be executed")
}

func TestFooerTableDriven(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  string
	}{
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 should bee Foo", 0, "Foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Fooer(tt.input)
			if ans != tt.want {
				t.Errorf("got: %s, want %s", ans, tt.want)
			}
		})
	}
}
