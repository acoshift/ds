package ds

import (
	"testing"
)

func TestStampModel(t *testing.T) {
	x := &StampModel{}
	if !x.CreatedAt.IsZero() {
		t.Errorf("expected initial CreatedAt to be zero; got %v", x.CreatedAt)
	}
	if !x.UpdatedAt.IsZero() {
		t.Errorf("expected initial UpdatedAt to be zero; got %v", x.UpdatedAt)
	}
	x.Stamp()
	if x.CreatedAt.IsZero() {
		t.Errorf("expected CreatedAt not to be zero")
	}
	if x.UpdatedAt.IsZero() {
		t.Errorf("expected UpdatedAt not to be zero")
	}
	if !x.CreatedAt.Equal(x.UpdatedAt) {
		t.Errorf("expected CreatedAt and UpdatedAt to be equals")
	}
	x.Stamp()
	if x.CreatedAt.Equal(x.UpdatedAt) {
		t.Errorf("expected CreatedAt and UpdatedAt not to be equals")
	}
}
