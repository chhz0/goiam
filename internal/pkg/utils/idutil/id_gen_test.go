package idutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IDGEN(t *testing.T) {
	set := map[string]struct{}{}
	for i := 0; i < 10000000; i++ {
		testId := GenerateInstanceID("test-table1", uint64(i), "test-")
		t.Log(testId)
		assert.NotEmpty(t, testId)
		_, ok := set[testId]
		assert.False(t, ok)
	}
}
