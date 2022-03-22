package parser

import (
    "encoding/json"
    "github.com/yguilai/nvm/version"
    "github.com/yguilai/sl"
    "io/ioutil"
    "net/http"
    "strings"
)

func init() {
    RegisterParser(Taobao, &StandardParser{})
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

func (p *TaobaoParser) GerVersions(resp *http.Response) ([]*version.Version, error) {
    tbVersions, err := p.getTbVersions(resp)
    if err != nil {
        return nil, err
    }
    versions := sl.MapperStream[*TbVersion, *version.Version](
        sl.Filter(
            sl.Stream(tbVersions),
            func(v *TbVersion) bool {
                if strings.Index(v.Name, "v") == 0 && v.Type == "dir" {
                    return true
                }
                return false
            },
        ),
        func(tb *TbVersion) *version.Version {
            name := strings.TrimSpace(tb.Name)
            return &version.Version{
                Name:     name,
                Url:      tb.Url,
                Packages: nil,
                Sort:     version.GetSortByVersion(name),
            }
        },
    ).CollectSlice()
    return versions, nil
}

func (p *TaobaoParser) GetPackages(v *version.Version) ([]*version.Package, error) {
    if v == nil {
        return nil, version.NilVersionErr
    }

    resp, err := http.Get(v.Url)
    if err != nil {
        return nil, err
    }
    tbVersions, err := p.getTbVersions(resp)
    if err != nil {
        return nil, err
    }

    packages := sl.MapperStream[*TbVersion, *version.Package](
        sl.Filter(
            sl.Stream(tbVersions),
            func(tb *TbVersion) bool {
                if tb.Type == "file" && strings.Contains(tb.Name, v.Name) {
                    return true
                }
                return false
            },
        ),
        func(tb *TbVersion) *version.Package {
            return &version.Package{
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
