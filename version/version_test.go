package version

import (
	"fmt"
	"testing"
)

func TestGetSort(t *testing.T) {
	fmt.Println(getSortByVersion("v9.7.1"))
}
