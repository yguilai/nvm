package version

import (
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

    Parser interface {
        GerVersions(resp *http.Response) ([]*Version, error)
        GetPackages(v *Version) ([]*Package, error)
    }
)

var parserMap = make(map[SourceType]Parser)

func FindAllValidVersions(url string, sourceType SourceType) ([]*Version, error) {
    if url == "" {
        url = defaultSource
    }
    if sourceType == UNKNOWN {
        sourceType = defaultSourceType
    }

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    parser := parserMap[sourceType]
    if parser == nil {
        return nil, ParserNotFoundErr
    }
    return parser.GerVersions(resp)
}

func RegisterParser(st SourceType, parser Parser) {
    parserMap[st] = parser
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
