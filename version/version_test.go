package version

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestGetSortByVersion(t *testing.T) {
	fmt.Println(GetSortByVersion("v9.7.1"))
}

func TestFindAllValidVersions(t *testing.T) {
	versions, err := FindAllValidVersions("https://registry.npmmirror.com/-/binary/node/", Taobao)
	if err != nil {
		panic(err)
	}
	if versions == nil {
		fmt.Println("empty version")
	}
	dump.P(versions)
}

func TestFindAllValidPackages(t *testing.T) {
	versions, err := FindAllValidVersions("https://registry.npmmirror.com/-/binary/node/", Taobao)
	if err != nil {
		panic(err)
	}
	if versions == nil {
		fmt.Println("empty version")
	}
	lastVersion := versions[len(versions)-1]
	dump.P(lastVersion)
	packages, err := FindAllValidPackages(lastVersion, Taobao)
	if err != nil {
		panic(err)
	}
	dump.P(packages)
}
