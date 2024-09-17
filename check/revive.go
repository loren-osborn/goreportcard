package check

// Revive is the check for the revive command
type Revive struct {
	Dir       string
	Filenames []string
}

// Name returns the name of the display name of the command
func (g Revive) Name() string {
	return "revive"
}

// Weight returns the weight this check has in the overall average
func (g Revive) Weight() float64 {
	return .10
}

// Percentage returns the percentage of .go files that pass revive
func (g Revive) Percentage() (float64, []FileSummary, error) {
	return GoTool(g.Dir, g.Filenames, []string{"golangci-lint", "run", "--disable-all", "--enable=revive", "--vendor"})
}

// Description returns the description of revive
func (g Revive) Description() string {
	return `Golint is a linter for Go source code.`
}
