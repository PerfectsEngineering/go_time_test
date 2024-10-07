package go_time_test

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeOrder(t *testing.T) {
	timeBefore := time.Now()
	slog.Info("timeBefore", "time", timeBefore.String())

	timeAfter := createRandomFile(t).ModTime()

	timeNowAfter := time.Now()

	duration := timeBefore.Sub(timeAfter)
	slog.Info("difference", "duration", duration.String(), "timeAfter", timeNowAfter.String())

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
	slog.Info("created file", "file", file.Name())

	fileInfo, err := file.Stat()
	require.NoError(t, err)

	return fileInfo
}

//func sysTime(t *testing.T) {
//var timespec unix.Timespec
//err := unix.ClockGettime(unix.CLOCK_DEFAULT, &timespec)

//require.NoError(t, err)

//sysTime := time.Unix(timespec.Sec, timespec.Nsec)

//slog.Info("sysTime", "time", sysTime.String())
//}
