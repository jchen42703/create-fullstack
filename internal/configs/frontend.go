package configs

type FrontendBase int

const (
	Next FrontendBase = iota
	CRA
)

func (base FrontendBase) String() string {
	return []string{"Next.js", "CRA"}[base]
}

type HuskyOptions struct {
	CommitLint bool
	Prettier   bool
	ESLint     bool
}

type FrontendConfig struct {
	Base                FrontendBase
	UseTypescript       bool
	UseSCSS             bool
	UseTailwind         bool
	UseStyledComponents bool
	HuskyOpts           *HuskyOptions
}
