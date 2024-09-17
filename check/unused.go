package check

// Unused checks Go code for unused constants, variables, functions and types
type Unused struct {
	Dir       string
	Filenames []string
}

// Name returns the name of the display name of the command
func (g Unused) Name() string {
	return "ineffassign"
}

// Weight returns the weight this check has in the overall average
func (g Unused) Weight() float64 {
	return 0.10
}

// Percentage returns the percentage of .go files that pass gofmt
func (g Unused) Percentage() (float64, []FileSummary, error) {
	return GoTool(g.Dir, g.Filenames, []string{"golangci-lint", "run", "--disable-all", "--enable=unused"})
}

// Description returns the description of Unused
func (g Unused) Description() string {
	return `<a href="https://github.com/dominikh/go-tools/tree/master/unused">unused</a> checks Go code for unused constants, variables, functions and types.`
}
