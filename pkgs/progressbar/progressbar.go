package progressbar

import (
	"bytes"
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"os"
	"time"
)

var CurrentBar pb.ProgressBar

func DisplayProgressBarForOwlFile(filepath string) {
	fmt.Println("Analysing OWL file..")
	f, _ := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	numberOfLines := lineCount(f)
	numberOfSeconds := numberOfLines / 600
	go displayProgressBarFor(numberOfLines, numberOfSeconds)
}

func FinishCurrentBar() {
	CurrentBar.Finish()
}

func displayProgressBarFor(numberOfLines, numberOfSeconds int) {
	CurrentBar := pb.StartNew(numberOfSeconds)
	CurrentBar.ShowCounters = false
	for i := 0; i < numberOfSeconds; i++ {
		CurrentBar.Increment()
		time.Sleep(1 * time.Second)
	}
}

func lineCount(r io.Reader) int {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count

		case err != nil:
			return 0
		}
	}
}
