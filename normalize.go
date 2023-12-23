package strutil

import (
	"strings"
)

type NormalizeOptions struct {
	TrimSpaces bool
	StripHTML  bool
	Lowercase  bool
	Uppercase  bool
}

var DefaultNormalizeOptions NormalizeOptions = NormalizeOptions{
	TrimSpaces: true,
	StripHTML:  true,
}

func NormalizeWithOptions(data string, options NormalizeOptions) string {
	if options.TrimSpaces {
		data = strings.TrimSpace(data)
	}

	if options.Lowercase {
		data = strings.ToLower(data)
	}

	if options.Uppercase {
		data = strings.ToUpper(data)
	}

	return data
}

func Normalize(data string) string {
	return NormalizeWithOptions(data, DefaultNormalizeOptions)
}
