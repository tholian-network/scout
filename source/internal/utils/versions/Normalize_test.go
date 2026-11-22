package versions

import "testing"

func TestNormalize(t *testing.T) {

	tests := []struct {
		value string
		want  string
		ok    bool
	}{
		{"1.2.3", "1.2.3", true},
		{"v1.2.3", "1.2.3", true},
		{"1.2", "1.2", true},
		{"1.2.x", "1.2.0", true},
		{"1.x", "1.0", true},
		{"1!2.0", "1:2.0", true},
		{"1.2.3-beta.1", "1.2.3-beta.1", true},
		{"1.0.post1", "1.0.post1", true},
		{"0.0.0", "0.0.0", true},
		{"*", "", false},
		{"latest", "", false},
		{"dev-main", "", false},
		{"", "", false},
	}

	for _, test := range tests {

		got, ok := Normalize(test.value)

		if got != test.want || ok != test.ok {
			t.Fatalf("Normalize(%q) = %q, %v; want %q, %v", test.value, got, ok, test.want, test.ok)
		}

	}

}
