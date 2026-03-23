package tencentcloud

import (
	"errors"
	"testing"
)

func TestResolveTestCamMfaUin_UsesExplicitOverride(t *testing.T) {
	got, err := resolveTestCamMfaUin("110000000343", func() (string, error) {
		t.Fatal("lookup should not be called when explicit uin is provided")
		return "", nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "110000000343" {
		t.Fatalf("expected explicit uin, got %q", got)
	}
}

func TestResolveTestCamMfaUin_FallsBackToLookup(t *testing.T) {
	got, err := resolveTestCamMfaUin("", func() (string, error) {
		return "110000000230", nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "110000000230" {
		t.Fatalf("expected looked up uin, got %q", got)
	}
}

func TestResolveTestCamMfaUin_ReturnsLookupError(t *testing.T) {
	expected := errors.New("lookup failed")

	_, err := resolveTestCamMfaUin("", func() (string, error) {
		return "", expected
	})
	if !errors.Is(err, expected) {
		t.Fatalf("expected %v, got %v", expected, err)
	}
}
