package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

type NginxBlock struct {
	StartLine   string
	EndLine     string
	AllContents string
	// split lines by \n on AllContents,
	// use make to create *[],
	// first create make([]*Type..)
	// then use &var to make it *
	AllLines          *[]*string
	NestedBlocks      []*NginxBlock
	TotalBlocksInside int
}

func (ngBlock *NginxBlock) IsBlock(line string) bool {
	// TODO Solve it using regex
	return false
}

func (ngBlock *NginxBlock) IsLine(line string) bool {
	// TODO Solve it using regex
	var validLine = regexp.MustCompile(`([^\n]*\n+)+`)
	return validLine.MatchString(line)
}

func (ngBlock *NginxBlock) HasComment(line string) bool {
	// TODO Solve it using regex
	return false
}

func GetNginxBlock(
	lines *[]*string,
	startIndex,
	endIndex,
	recursionMax int,
) *NginxBlock {
	var b *NginxBlock
	return b
}

type NginxBlocks struct {
	blocks      *[]*NginxBlock
	AllContents string
	// split lines by \n on AllContents
	AllLines *[]*string
}

func GetNginxBlocks(configContent string) *NginxBlocks {
	var b NginxBlocks
	var al = make([]*string, 0)
	b.AllContents = configContent
	var lineRegex string = `.*\n`
	sent := regexp.MustCompile(lineRegex)
	matches := sent.FindAllStringSubmatchIndex(configContent, -1)

	for _, match := range matches {
		for i := 0; i < len(match)-1; i += 2 {
			al = append(al, &configContent[match[i]+match[i+1]])
			fmt.Println("------")
		}

	}
	return &b
}
func main() {
	content, err := ioutil.ReadFile("nginx.conf")
	if err != nil {
		log.Fatal(err)
	}
	nbs := GetNginxBlocks(string(content))
	fmt.Println(nbs)
}
