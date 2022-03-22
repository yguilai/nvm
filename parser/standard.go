package parser

import (
    "github.com/yguilai/nvm/version"
    "net/http"
)

func init() {
    RegisterParser(Standard, &StandardParser{})
}

type StandardParser struct {
}

func (s *StandardParser) GerVersions(resp *http.Response) ([]*version.Version, error) {
    //TODO implement me
    panic("implement me")
}

func (p *StandardParser) GetPackages(v *version.Version) ([]*version.Package, error) {
    //TODO implement me
    panic("implement me")
}

var _ Parser = (*StandardParser)(nil)
