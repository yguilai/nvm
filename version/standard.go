package version

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func init() {
	RegisterParser(Standard, &StandardParser{})
}

type StandardParser struct {
}

func (p *StandardParser) GerVersions(source string) ([]*Version, error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var versions []*Version
	doc.Find("a").Each(func(_ int, sel *goquery.Selection) {
		// also skip v0.x.x
		if val, exists := sel.Attr("href"); exists && strings.Index(val, "v") == 0 && !strings.Contains(val, skipVersionPrefix) {
			name := strings.TrimRight(val, "/")
			versions = append(versions, &Version{
				Name: name,
				Url:  source + val,
				Sort: GetSortByVersion(name),
			})
		}
	})
	return versions, nil
}

func (p *StandardParser) GetPackages(v *Version, os, arch string) ([]*Package, error) {
	if v == nil {
		return nil, NilVersionErr
	}
	resp, err := http.Get(v.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(v.Url, "/") {
		v.Url += "/"
	}

	var packages []*Package
	var sumsFileUrl string
	doc.Find("a").Each(func(_ int, sel *goquery.Selection) {
		val, exists := sel.Attr("href")
		if !exists {
			return
		}
		if val == shaSums256Filename {
			sumsFileUrl = v.Url + shaSums256Filename
		}
		if strings.HasPrefix(v.Url, nodeFilePrefix) {
			packages = append(packages, &Package{
				Filename:    val,
				DownloadUrl: v.Url + val,
			})
		}
	})

	sumsMap, err := p.GetShaSumsMap(sumsFileUrl)
	if err != nil {
		return nil, err
	}

	for _, pac := range packages {
		if s, ok := sumsMap[pac.Filename]; ok {
			pac.ShaSums = s
		}
	}
	return packages, nil
}

func (p *StandardParser) GetShaSumsMap(url string) (map[string]string, error) {
	if url == "" {
		return nil, NilSHASUMSErr
	}
	resp, err := http.Get(url)
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

var _ Parser = (*StandardParser)(nil)
