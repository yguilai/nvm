package version

// SourceType source type of nodejs download url
type SourceType int

const (
    UNKNOWN SourceType = iota
    // Standard e.g. https://nodejs.org/dist/
    Standard
    // Taobao e.g. https://registry.npmmirror.com/-/binary/node/
    Taobao
)

const (
    defaultSource     = "https://nodejs.org/dist/"
    defaultSourceType = Standard
)
