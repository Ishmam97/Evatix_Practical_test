package main

import (
	"fmt"
	"log"
	"os"
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
	o := fmt.Sprintf("%s -> %s\n", cliStreamerRecord.Title, cliStreamerRecord.Message1)
	fmt.Println(o)
	go writeToFile("output", o)
	//delay
	time.Sleep(time.Duration(cliStreamerRecord.StreamDelay) * time.Second)
	o1 := fmt.Sprintf("%s -> %s\n", cliStreamerRecord.Title, cliStreamerRecord.Message2)
	fmt.Println(o1)
	go writeToFile("output", o1)
}

var mutex = sync.Mutex{}

func writeToFile(file string, o string) {
	mutex.Lock()
	// Write to file
	f, err := os.OpenFile("out.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(o); err != nil {
		log.Println(err)
	}
	mutex.Unlock()

}
func main() {
	args := "Run,Title,Message 1,Message 2,Stream Delay,Run Times\n2,CLI Invoke1,First helloo,Second Ishmam,3,2\n2,CLI Invoke2,yoooo,TSM,5,2\n3,CLI Invoke3,First run 10 times,Second Msg 2,2,10"
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
			for i := 0; i < runner.RunTimes; i++ {
				runner.CliStreamerRecord().Stream()
			}
			wg.Done()
		}(runner)
	}
	//wait
	wg.Wait()

}
