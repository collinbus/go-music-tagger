package flac

import "testing"

func TestCreatePadding(t *testing.T) {
	padding := CreatePadding(10, 15)

	if len(padding) != 9 {
		t.Errorf("Wrong paddding size. It should be %d but was %d", 9, len(padding))
	}
}
