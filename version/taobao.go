package version

import (
	"bufio"
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

func (p *TaobaoParser) GerVersions(source string) ([]*Version, error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	tbVersions, err := p.getTbVersions(resp)
	if err != nil {
		return nil, err
	}
	versions := sl.MapperStream[*TbVersion, *Version](
		sl.Filter(
			sl.Stream(tbVersions),
			func(v *TbVersion) bool {
				// skip v0.x.x
				if strings.Index(v.Name, "v") == 0 && v.Type == "dir" && !strings.Contains(v.Name, skipVersionPrefix) {
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

func (p *TaobaoParser) GetPackages(v *Version, os, arch string) ([]*Package, error) {
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

	var sha *TbVersion
	for _, tb := range tbVersions {
		if tb.Name == shaSums256Filename {
			sha = tb
		}
	}
	if sha == nil {
		return nil, NilSHASUMSErr
	}
	sumsMap, err := p.GetShaSumsMap(sha.Url)
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
				ShaSums:     sumsMap[tb.Name],
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

func (p *TaobaoParser) GetShaSumsMap(sumsUrl string) (map[string]string, error) {
	resp, err := http.Get(sumsUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	sumsMap := make(map[string]string)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		text := scanner.Text()
		splits := strings.Split(text, "  ")
		sumsMap[splits[1]] = splits[0]
	}
	return sumsMap, nil
}

var _ Parser = (*TaobaoParser)(nil)
