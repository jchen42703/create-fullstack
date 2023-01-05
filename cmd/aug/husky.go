package aug

import (
	"fmt"

	"github.com/jchen42703/create-fullstack/cmd/opts"
)

type HuskyOptions struct {
	CommitLint *opts.PackageOptions `yaml:"commitlint"`
	// Might need to change these to include specific formatters/linters and their versions
	Format bool `yaml:"format"`
	Lint   bool `yaml:"lint"`
}

// Adds husky to a GitHub repository.
type HuskyAugmenter struct {
	HuskyOpts *HuskyOptions
}

// Adds husky to a GitHub repository.
// Exits early if any step fails.
func (a *HuskyAugmenter) Augment() error {
	var err error
	if a.HuskyOpts.CommitLint != nil {
		err = addCommitLint(a.HuskyOpts.CommitLint.Version)
		if err != nil {
			return fmt.Errorf("failed to add commitlint: %s", err)
		}
	}

	return nil
}

func addCommitLint(version string) error {
	return nil
}
