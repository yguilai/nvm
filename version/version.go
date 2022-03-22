package version

import (
	"net/http"
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

func FindAllValidVersions(url string, sourceType SourceType) ([]*Version, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	p := LoadParser(sourceType)
	if p == nil {
		return nil, ParserNotFoundErr
	}
	return p.GerVersions(resp)
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
	multiplier := 1000
	for _, ver := range verNums {
		num, _ := strconv.Atoi(ver)
		sort += num * multiplier
		multiplier = multiplier / 10
	}
	return sort
}
