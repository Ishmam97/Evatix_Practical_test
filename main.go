package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
)

// "Title,Message 1,Message 2,Stream Delay,Run Times\nCLI Invoker Name,First Message,Second Msg,2,10"
type CliStreamerRecord struct {
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

type CliRunnerRecord struct {
	// How many streamer will run.
	Run         string `csv:"Run"`
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

func (cliRunnerRecord CliRunnerRecord) CliStreamerRecord() CliStreamerRecord {
	return CliStreamerRecord{
		Title:       cliRunnerRecord.Title,
		Message1:    cliRunnerRecord.Message1,
		Message2:    cliRunnerRecord.Message2,
		StreamDelay: cliRunnerRecord.StreamDelay,
		RunTimes:    cliRunnerRecord.RunTimes,
	}
}

func (cliRunnerRecord CliRunnerRecord) CliStreamerRecordCsv() string {
	cliStreamerRecords := []CliStreamerRecord{cliRunnerRecord.CliStreamerRecord()}

	out, err := gocsv.MarshalString(cliStreamerRecords)

	if err != nil {
		panic(err)
	}

	return out
}

func Csv(cliRunners *[]CliRunnerRecord) string {
	out, err := gocsv.MarshalString(cliRunners)

	if err != nil {
		panic(err)
	}

	return out
}

//stream function
func (cliStreamerRecord CliStreamerRecord) Stream() {
	fmt.Printf("%s -> %s\n", cliStreamerRecord.Title, cliStreamerRecord.Message1)
	//delay
	time.Sleep(time.Duration(cliStreamerRecord.StreamDelay) * time.Second)
	fmt.Printf("%s -> %s\n", cliStreamerRecord.Title, cliStreamerRecord.Message2)

}

func main() {
	args := "Run,Title,Message 1,Message 2,Stream Delay,Run Times\n2,CLI Invoke1,First helloo,Second Ishmam,3,2\n2,CLI Invoke2,yoooo,TSM,5,2\n3,CLI Invoke3,First Msg 1,Second Msg 2,2,150"
	var cliRunners []CliRunnerRecord
	// c := make(chan []string)
	var wg sync.WaitGroup
	gocsv.UnmarshalString(
		args,
		&cliRunners)

	// fmt.Print(Csv(&cliRunners))
	// fmt.Println("---------------------------------")
	for _, runner := range cliRunners {
		wg.Add(1)
		go func(runner CliRunnerRecord) {
			for i := 0; i < 10; i++ {
				runner.CliStreamerRecord().Stream()
			}
			wg.Done()
		}(runner)
	}
	wg.Wait()

}
