package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSortByVersion(t *testing.T) {
	sort := GetSortByVersion("v9.7.1")
	assert.Equal(t, sort, 90701)
}

func TestLoadParser(t *testing.T) {
	assert.Condition(t, func() bool {
		parser := LoadParser(Standard)
		_, ok := parser.(*StandardParser)
		return ok
	})
}
