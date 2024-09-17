package check

// GoSimple is the linter for Go source code that specializes in simplifying code
type GoSimple struct {
	Dir       string
	Filenames []string
}

// Name returns the name of the display name of the command
func (g GoSimple) Name() string {
	return "gosimple"
}

// Weight returns the weight this check has in the overall average
func (g GoSimple) Weight() float64 {
	return .10
}

// Percentage returns the percentage of .go files that pass gosimple
func (g GoSimple) Percentage() (float64, []FileSummary, error) {
	return GoTool(g.Dir, g.Filenames, []string{"golangci-lint", "run", "--disable-all", "--enable=gosimple"})
}

// Description returns the description of go vet
func (g GoSimple) Description() string {
	return `GoSimple is the linter for Go source code that specializes in simplifying code.`
}
