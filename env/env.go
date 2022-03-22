package env

import (
    "github.com/yguilai/nvm/parser"
    "os"
    "path/filepath"
)

const (
    NvmHomeVariable       = "NVM_HOME"
    NvmSourceVariable     = "NVM_SOURCE"
    NvmSourceTypeVariable = "NVM_SOURCE_TYPE"
    NodeHomeVariable      = "NODE_HOME"
)

func NvmHome() string {
    if home := os.Getenv(NvmHomeVariable); home != "" {
        return home
    }
    defaultNvmHomeVariable, _ := os.UserHomeDir()
    return filepath.Join(defaultNvmHomeVariable, ".nvm")
}

func NodeHome() string {
    return getEnvDefault(NodeHomeVariable, filepath.Join(NvmHome(), "node"))
}

func NvmSource() (s string, st parser.SourceType) {
    s = getEnvDefault(NvmSourceVariable, parser.DefaultSource)
    st = parser.SourceType(getEnvDefault(NvmSourceTypeVariable, string(parser.DefaultSourceType)))
    return
}

func getEnvDefault(v string, d string) string {
    if val := os.Getenv(v); val != "" {
        return val
    }
    return d
}
