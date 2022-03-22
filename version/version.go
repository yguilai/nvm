package version

import (
    "github.com/yguilai/nvm/parser"
    "net/http"
    "strconv"
    "strings"
)

type (
    Version struct {
        Name     string
        Url      string
        Packages []*Package
        Sort     int
    }

    Package struct {
        Filename    string
        ShaSums     string
        DownloadUrl string
    }
)

func FindAllValidVersions(url string, sourceType parser.SourceType) ([]*Version, error) {
    if url == "" {
        url = parser.DefaultSource
    }
    if sourceType == parser.UNKNOWN {
        sourceType = parser.DefaultSourceType
    }

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    p := parser.LoadParser(sourceType)
    if p == nil {
        return nil, ParserNotFoundErr
    }
    return p.GerVersions(resp)
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
