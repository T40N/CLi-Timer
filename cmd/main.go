package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	hours := flag.Int("h", 0, "Liczba godzin")
	minutes := flag.Int("m", 0, "Liczba minut")
	seconds := flag.Int("s", 0, "Liczba sekund")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: timer -h=1 -m=30 -s=10")
		return
	}

	s := spinner.New(spinner.CharSets[7], 1* time.Second)
	s.Start()

	totalSeconds := (*hours * 3600) + (*minutes * 60) + *seconds
	for totalSeconds > 0  {
		hs := totalSeconds / 3600
		mins := (totalSeconds % 3600) / 60
		sec := totalSeconds % 60

		s.Suffix = fmt.Sprintf(" %02d:%02d:%02d", hs, mins, sec)

		time.Sleep(1 * time.Second)
		totalSeconds--
	}
	s.Stop()
	playAlarm()
}

func playAlarm() {
	cmd := exec.Command("afplay", "/System/Library/Sounds/Glass.aiff")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Błąd podczas odtwarzania alarmu:", err)
	}
}