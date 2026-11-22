package versions

import "testing"

func TestConstraint(t *testing.T) {

	tests := []struct {
		operator string
		value    string
		want     string
		ok       bool
	}{
		{"", "1.2.3", "= 1.2.3", true},
		{"=", "1.2.3", "= 1.2.3", true},
		{"==", "1.2.3", "= 1.2.3", true},
		{">=", "1.2.3", ">= 1.2.3", true},
		{">", "1.2.3", "> 1.2.3", true},
		{"<=", "1.2.3", "<= 1.2.3", true},
		{"<", "1.2.3", "< 1.2.3", true},
		{"!=", "1.2.3", "", false},
		{">=", "latest", "", false},
	}

	for _, test := range tests {

		got, ok := Constraint(test.operator, test.value)

		if got != test.want || ok != test.ok {
			t.Fatalf("Constraint(%q, %q) = %q, %v; want %q, %v", test.operator, test.value, got, ok, test.want, test.ok)
		}

	}

}
