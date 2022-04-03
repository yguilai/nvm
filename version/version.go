package version

import (
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
