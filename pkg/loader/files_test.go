package loader

import (
	"os"
	"testing"

	"github.com/lyraproj/lyra/pkg/logger"
	"github.com/stretchr/testify/assert"
)

var findfileTests = []struct {
	title    string
	dir      string
	glob     string
	error    bool
	expected []string
}{
	{
		title:    "glob match",
		dir:      "testdata/files",
		glob:     "prefix-*",
		error:    false,
		expected: []string{"testdata/files/prefix-a", "testdata/files/prefix-b"},
	},
	{
		title:    "exact match",
		dir:      "testdata/files",
		glob:     "prefix-a",
		error:    false,
		expected: []string{"testdata/files/prefix-a"},
	},
	{
		title:    "no matches",
		dir:      "testdata/files",
		glob:     "nomatch-*",
		error:    false,
		expected: []string{},
	},
	{
		title: "invalid dir",
		dir:   "testdata/invaliddir",
		glob:  "prefix-*",
		error: true,
	},
}

func Test_FindFile(t *testing.T) {
	for _, test := range findfileTests {
		t.Run(test.title, func(t *testing.T) {
			spec := logger.Spec{
				Name:   "find-files",
				Level:  "debug",
				Output: os.Stderr,
			}
			l := Loader{
				pluginPath: []string{test.dir},
				logger:     logger.Initialise(spec),
			}
			files := l.findFiles(test.glob)
			if !test.error {
				assert.Equal(t, len(test.expected), len(files))
				for _, expected := range test.expected {
					assert.Contains(t, files, expected)
				}
			}
		})
	}
}
