package version

import (
	"encoding/json"
	"github.com/yguilai/sl"
	"io/ioutil"
	"net/http"
	"strings"
)

func init() {
	parserMap[Taobao] = &TaobaoParser{}
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

func (p *TaobaoParser) GerVersions(resp *http.Response) ([]*Version, error) {
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var tbVersions []*TbVersion
	err = json.Unmarshal(bytes, &tbVersions)
	if err != nil {
		return nil, err
	}

	if tbVersions == nil {
		return nil, nil
	}

	versions := sl.MapperStream[*TbVersion, *Version](
		sl.Filter(
			sl.Stream(tbVersions),
			func(v *TbVersion) bool {
				if strings.Index(v.Name, "v") == 0 && v.Type == "dir" {
					return true
				}
				return false
			}),
		func(tb *TbVersion) *Version {
			name := strings.TrimSpace(tb.Name)
			return &Version{
				Name:     name,
				Url:      tb.Url,
				Packages: nil,
				Sort:     getSortByVersion(name),
			}
		},
	).CollectSlice()
	return versions, nil
}

func (p *TaobaoParser) GetPackages(v *Version) ([]*Package, error) {
	//TODO implement me
	panic("implement me")
}

var _ Parser = (*TaobaoParser)(nil)
