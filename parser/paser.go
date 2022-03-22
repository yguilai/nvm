package parser

import (
    "github.com/yguilai/nvm/version"
    "net/http"
)

type Parser interface {
    GerVersions(resp *http.Response) ([]*version.Version, error)
    GetPackages(v *version.Version) ([]*version.Package, error)
}

var parserMap = make(map[SourceType]Parser)

func RegisterParser(st SourceType, parser Parser) {
    parserMap[st] = parser
}

func LoadParser(st SourceType) Parser {
    return parserMap[st]
}
