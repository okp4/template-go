package version

import (
	"strings"
	"testing"
)

func TestNewInfo(t *testing.T) {
	expectedGoVersion := "go version"
	result := NewInfo()
	if !strings.HasPrefix(result.GoVersion, expectedGoVersion) {
		t.Errorf("Expected %s, starts with %s", result.GoVersion, expectedGoVersion)
	}
}

func TestInfoString(t *testing.T) {
	expected := `go-template: 0.0
git commit: 13245
go fake version 42`

	info := new(Info)
	info.Name = "go-template"
	info.Version = "0.0"
	info.GitCommit = "13245"
	info.GoVersion = "go fake version 42"

	result := info.String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
