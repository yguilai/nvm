package version

// SourceType source type of nodejs download url
type SourceType string

const (
	Unknown SourceType = ""
	// Standard e.g. https://nodejs.org/dist/
	Standard SourceType = "standard"
	// Taobao e.g. https://registry.npmmirror.com/-/binary/node/
	Taobao SourceType = "taobao"
)
