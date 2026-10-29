package pkg

import "strings"
import "testing"

func TestParsePackageFiles(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"e2e-scout-pkg\t/usr/local/bin/e2e-scout-pkg",
		"e2e-scout-pkg\t/usr/local/etc/e2e-scout-pkg.conf",
		"curl\t/usr/local/bin/curl",
		"curl\t/usr/local/lib/libcurl.so.4",
	}, "\r\n"))

	files := parsePackageFiles(string(buffer), "")

	if len(files) != 4 {
		t.Fatalf("Expected 4 files, got %d", len(files))
	}

	if files[0] != "/usr/local/bin/e2e-scout-pkg" {
		t.Fatalf("Unexpected file: %s", files[0])
	}

	if files[1] != "/usr/local/etc/e2e-scout-pkg.conf" {
		t.Fatalf("Unexpected file: %s", files[1])
	}

	if files[2] != "/usr/local/bin/curl" {
		t.Fatalf("Unexpected file: %s", files[2])
	}

	if files[3] != "/usr/local/lib/libcurl.so.4" {
		t.Fatalf("Unexpected file: %s", files[3])
	}

}

func TestParsePackageFilesByParent(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"e2e-scout-pkg\t/usr/local/bin/e2e-scout-pkg",
		"e2e-scout-pkg\t/usr/local/etc/e2e-scout-pkg.conf",
		"curl\t/usr/local/bin/curl",
		"curl\t/usr/local/lib/libcurl.so.4",
	}, "\r\n"))

	files := parsePackageFiles(string(buffer), "e2e-scout-pkg")

	if len(files) != 2 {
		t.Fatalf("Expected 2 files for e2e-scout-pkg, got %d", len(files))
	}

	if files[0] != "/usr/local/bin/e2e-scout-pkg" {
		t.Fatalf("Unexpected file: %s", files[0])
	}

	if files[1] != "/usr/local/etc/e2e-scout-pkg.conf" {
		t.Fatalf("Unexpected file: %s", files[1])
	}

}

func TestParsePackageFilesUnknownParent(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"e2e-scout-pkg\t/usr/local/bin/e2e-scout-pkg",
		"e2e-scout-pkg\t/usr/local/etc/e2e-scout-pkg.conf",
		"curl\t/usr/local/bin/curl",
		"curl\t/usr/local/lib/libcurl.so.4",
	}, "\r\n"))

	files := parsePackageFiles(string(buffer), "missing")

	if len(files) != 0 {
		t.Fatalf("Expected 0 files for missing parent, got %d", len(files))
	}

}
