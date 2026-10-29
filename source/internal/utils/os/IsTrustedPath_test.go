package os

import "testing"

func TestIsTrustedPath(t *testing.T) {

	tests := []struct {
		value string
		want  bool
	}{
		{"/usr/bin/foo", true},
		{"/usr/local/bin/foo", true},
		{"/sbin/init", true},
		{"/usr/lib/libfoo.so", true},
		{"/tmp/evil", false},
		{"/var/tmp/evil", false},
		{"/dev/shm/evil", false},
		{"/home/user/evil", false},
		{"/overlay/upper/usr/bin/foo", true},
		{"/overlay/upper/tmp/evil", false},
		{"/rom/usr/sbin/foo", true},
		{"memfd:evil", false},
		{"/tmp/evil (deleted)", false},
	}

	for _, test := range tests {

		if got := IsTrustedPath(test.value); got != test.want {
			t.Fatalf("IsUntrustedPath(%q) = %v, want %v", test.value, got, test.want)
		}

	}

}

