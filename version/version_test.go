package version

import (
	"fmt"
	"testing"
)

func TestGetSortByVersion(t *testing.T) {
	fmt.Println(GetSortByVersion("v9.7.1"))
}
