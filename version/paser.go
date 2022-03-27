package version

type Parser interface {
	GerVersions(source string) ([]*Version, error)
	GetPackages(v *Version, os, arch string) ([]*Package, error)
	GetShaSumsMap(url string) (map[string]string, error)
}

var parserMap = make(map[SourceType]Parser)

func RegisterParser(st SourceType, parser Parser) {
	parserMap[st] = parser
}

func LoadParser(st SourceType) Parser {
	return parserMap[st]
}
