package version

import (
	"encoding/json"
	"github.com/yguilai/sl"
	"io/ioutil"
	"net/http"
	"strings"
)

func init() {
	RegisterParser(Taobao, &TaobaoParser{})
}

type (
	TaobaoParser struct {
	}

	TbVersion struct {
		Id       string
		Category string
		Name     string
		Date     string
		Type     string
		Url      string
		Modified string
	}
)

const SHASUMS256Filename = "SHASUMS256.txt"

func (p *TaobaoParser) GerVersions(resp *http.Response) ([]*Version, error) {
	tbVersions, err := p.getTbVersions(resp)
	if err != nil {
		return nil, err
	}
	versions := sl.MapperStream[*TbVersion, *Version](
		sl.Filter(
			sl.Stream(tbVersions),
			func(v *TbVersion) bool {
				if strings.Index(v.Name, "v") == 0 && v.Type == "dir" {
					return true
				}
				return false
			},
		),
		func(tb *TbVersion) *Version {
			name := strings.TrimRight(tb.Name, "/")
			return &Version{
				Name:     name,
				Url:      tb.Url,
				Packages: nil,
				Sort:     GetSortByVersion(name),
			}
		},
	).CollectSlice()
	return versions, nil
}

func (p *TaobaoParser) GetPackages(v *Version) ([]*Package, error) {
	if v == nil {
		return nil, NilVersionErr
	}

	resp, err := http.Get(v.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	tbVersions, err := p.getTbVersions(resp)
	if err != nil {
		return nil, err
	}

	packages := sl.MapperStream[*TbVersion, *Package](
		sl.Filter(
			sl.Stream(tbVersions),
			func(tb *TbVersion) bool {
				if tb.Type == "file" {
					return true
				}
				return false
			},
		),
		func(tb *TbVersion) *Package {
			return &Package{
				Filename:    tb.Name,
				DownloadUrl: tb.Url,
			}
		},
	).CollectSlice()

	return packages, nil
}

func (p *TaobaoParser) getTbVersions(resp *http.Response) ([]*TbVersion, error) {
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var tbVersions []*TbVersion
	err = json.Unmarshal(bytes, &tbVersions)
	if err != nil {
		return nil, err
	}
	return tbVersions, nil
}

var _ Parser = (*TaobaoParser)(nil)
