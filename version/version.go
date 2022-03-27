package version

import (
	"runtime"
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

func FindAllValidVersions(source string, sourceType SourceType) ([]*Version, error) {
	p := LoadParser(sourceType)
	if p == nil {
		return nil, ParserNotFoundErr
	}
	return p.GerVersions(source)
}

func FindAllValidPackages(v *Version, sourceType SourceType) ([]*Package, error) {
	if v == nil {
		return nil, NilVersionErr
	}
	p := LoadParser(sourceType)
	if p == nil {
		return nil, ParserNotFoundErr
	}
	return p.GetPackages(v, runtime.GOOS, runtime.GOARCH)
}

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
