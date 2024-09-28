package go_time_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeOrder(t *testing.T) {
	timeBefore := time.Now()

	timeAfter := createRandomFile(t).ModTime()

	assert.True(
		t,
		timeBefore.Before(timeAfter),
		"directory modified time should be after time: (%s) > (%s)",
		timeBefore,
		timeAfter,
	)
}

func createRandomFile(t *testing.T) os.FileInfo {
	t.Helper()

	file, err := os.CreateTemp(os.TempDir(), "file")
	require.NoError(t, err)

	fileInfo, err := file.Stat()
	require.NoError(t, err)

	return fileInfo
}
