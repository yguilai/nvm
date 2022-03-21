package cmd

import "strings"

const v = "1.0.0"

var (
	Branch string
	Commit string
)

func Version() string {
	var b strings.Builder
	b.WriteString(v)

	if Branch != "" {
		b.WriteByte('\n')
		b.WriteString("branch: ")
		b.WriteString(Branch)
	}

	if Commit != "" {
		b.WriteByte('\n')
		b.WriteString("commit: ")
		b.WriteString(Commit)
	}
	return b.String()
}
