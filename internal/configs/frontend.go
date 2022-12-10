package configs

type FrontendBase int

const (
	Next FrontendBase = iota
	CRA
)

func (base FrontendBase) String() string {
	return []string{"Next.js", "CRA"}[base]
}

type FrontendConfig struct {
	Base                FrontendBase
	DesiredImageName    string
	UseTypescript       bool
	UseSCSS             bool
	UseTailwind         bool
	UseStyledComponents bool
	HuskyOptions        *struct {
		CommitLint bool
		Prettier   bool
		ESLint     bool
	}
}
