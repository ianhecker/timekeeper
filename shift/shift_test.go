package shift_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ianhecker/timekeeper/shift"
)

func TestShiftUnmarshal(t *testing.T) {
	var s shift.Shift
	var markdown string = "## 2022-11-28:1000-2000"

	err := s.Unmarshal(markdown)
	assert.NoError(t, err)

	fmt.Println(s.Start, s.End)
}
