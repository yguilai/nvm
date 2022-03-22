package version

import (
	"net/http"
)

type Parser interface {
	GerVersions(resp *http.Response) ([]*Version, error)
	GetPackages(v *Version) ([]*Package, error)
}

var parserMap = make(map[SourceType]Parser)

func RegisterParser(st SourceType, parser Parser) {
	parserMap[st] = parser
}

func LoadParser(st SourceType) Parser {
	return parserMap[st]
}
