## Weird `time.Now()` behaviour in Containers

Consider the test code in [time_test.go](./time_test.go) that checks the timestamps before and after a file is created:

```go
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
```

The test passes when I run it on my Computer (macOS, M2) using:
```bash
go test ./...
```

But it fails, If I try to run the same test command in a `go` Docker container.

You can test it out by running the following ocmmand in this directory:
```bash
docker build -t go_time_test . && docker run --rm --name time_test go_time_test
```

I get an error message similar to:
```
--- FAIL: TestTimeOrder (0.00s)
    time_test.go:24: 
        	Error Trace:	/app/time_test.go:24
        	Error:      	Should be true
        	Test:       	TestTimeOrder
        	Messages:   	directory modified time should be after time: (2024-10-07 19:29:57.07753727 +0000 UTC m=+0.000286669) > (2024-10-07 19:29:57.076465559 +0000 UTC)
```

> You'll also notice the test fails on CircleCI too, which I assume is using a similar Container setup to run the tests.

### Why does this happen?

I'm investigating why this happens, because I assumed time.Now() should be using the system time also used by the filesystem.

Please open an issue or reach out to me to let me know if you have an understanding of why this happens.
