package main

import (
	"log/slog"
	"time"
)

func main() {

	timeNow := time.Now()

	slog.Info("NOW", "time", timeNow.String())

}
