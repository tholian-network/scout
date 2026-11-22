package versions

import "testing"

func TestIsPartial(t *testing.T) {

	tests := []struct {
		value string
		want  bool
	}{
		{"1", true},
		{"1.2", true},
		{"1.2.3", false},
		{"1.2.3.4", false},
		{"1.x", true},
		{"1.2.*", true},
		{"1.2.3-beta.1", false},
	}

	for _, test := range tests {

		if got := IsPartial(test.value); got != test.want {
			t.Fatalf("IsPartial(%q) = %v, want %v", test.value, got, test.want)
		}

	}

}
