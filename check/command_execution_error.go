package check

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// CommandExecutionError encapsulates command, stdout, stderr, exit code, and diagnostic information.
type CommandExecutionError struct {
	Command  []string
	Stdout   *bytes.Buffer
	Stderr   *bytes.Buffer
	ExitCode int
	Err      error
}

// Error implements the error interface, returns the verbose command execution diagnostic information.
func (e *CommandExecutionError) Error() string {
	return fmt.Sprintf("Command failed: %v\n%s", e.Err, generateDiagnostic(e.Command, e.Stdout, e.Stderr, e.ExitCode))
}

// Generates a shell command for reproducing a failed subprocess execution
func generateDiagnostic(command []string, stdout, stderr *bytes.Buffer, exitCode int) string {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Sprintf("Error retrieving working directory: %v", err)
	}

	var result strings.Builder
	result.WriteString(fmt.Sprintf("(cd %s; exec 4>&1 5>&2 6>&1 ;", cwd))
	for _, arg := range command {
		if strings.ContainsAny(arg, " \t\n\"'|&;<>()$*?~#\\") {
			result.WriteString(fmt.Sprintf(" '%s'", strings.ReplaceAll(strings.ReplaceAll(arg, "\\", "'\\\\'"), "'", "'\\''")))
		} else {
			result.WriteString(fmt.Sprintf(" %s", arg))
		}
	}
	result.WriteString(" > /dev/fd/4 2> /dev/fd/5; exit_status=$? ; ")
	result.WriteString(fmt.Sprintf("diff -u <(cat <<'EXPECTED_STDOUT'\n%s\nEXPECTED_STDOUT\n) /dev/fd/4 >&6 ; ", stdout.String()))
	result.WriteString(fmt.Sprintf("diff -u <(cat <<'EXPECTED_STDERR'\n%s\nEXPECTED_STDERR\n) /dev/fd/5 ; ", stderr.String()))
	result.WriteString("cat <&6 ;")
	result.WriteString(fmt.Sprintf("[ $exit_status -ne %d ] && echo \"Exit status: $exit_status instead of %d\" >&2 )\n", exitCode, exitCode))

	return result.String()
}

// readCloserWrapper wraps a TeeReader and the original io.ReadCloser.
type readCloserWrapper struct {
    io.Reader      // This will hold the TeeReader
    closer io.Closer // This will hold the original io.ReadCloser
}

// Close calls the original ReadCloser's Close method.
func (r *readCloserWrapper) Close() error {
    return r.closer.Close()
}

// newTeeReadCloser creates a new io.ReadCloser that duplicates
// the data into a buffer while passing the reader along.
func newTeeReadCloser(rc io.ReadCloser) (io.ReadCloser, *bytes.Buffer) {
    buf := new(bytes.Buffer)
    tee := io.TeeReader(rc, buf) // Reads from rc and writes to buf
    return &readCloserWrapper{Reader: tee, closer: rc}, buf
}


