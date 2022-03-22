package version

import (
	"net/http"
)

func init() {
	RegisterParser(Standard, &StandardParser{})
}

type StandardParser struct {
}

func (s *StandardParser) GerVersions(resp *http.Response) ([]*Version, error) {
	//TODO implement me
	panic("implement me")
}

func (p *StandardParser) GetPackages(v *Version, os, arch string) ([]*Package, error) {
	//TODO implement me
	panic("implement me")
}

var _ Parser = (*StandardParser)(nil)
