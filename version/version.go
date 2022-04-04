package version

import (
	"regexp"
	"strconv"
	"strings"
)

type (
	Version struct {
		Name     string
		Url      string
		Sort     int
		Packages []*Package
	}

	Package struct {
		Filename    string
		ShaSums     string
		DownloadUrl string
	}
)

func GetSortByVersion(v string) (sort int) {
	verNums := strings.Split(v[1:], ".")
	multiplier := 10000
	for _, ver := range verNums {
		num, _ := strconv.Atoi(ver)
		sort += num * multiplier
		multiplier = multiplier / 100
	}
	return sort
}

// IsVersionDir verify name of a dir entry
// the version dir name should be vX.X.X or vX.X
// such as v1.17.1 or v1.17
func IsVersionDir(dir string) (bool, error) {
	return regexp.MatchString(`v[1-9]+\.[1-9]+(\.[1-9]+)?`, dir)
}
