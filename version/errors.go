package version

import "errors"

var (
	ParserNotFoundErr = errors.New("parser not found")
	NilVersionErr     = errors.New("version is nil")
	NilSHASUMSErr     = errors.New("SHASUM256.txt file not found")
)
