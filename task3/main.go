package main

import "regexp"

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

}

func (ngBlock *NginxBlock) IsLine(line string) bool {
	// TODO Solve it using regex
	var validLine = regexp.MustCompile(`([^\n]*\n+)+`)
	return validLine.MatchString(line)
}

func (ngBlock *NginxBlock) HasComment(line string) bool {
	// TODO Solve it using regex
}

type NginxBlocks struct {
	blocks      *[]*NginxBlock
	AllContents string
	// split lines by \n on AllContents
	AllLines *[]*string
}

func GetNginxBlock(
	lines *[]*string,
	startIndex,
	endIndex,
	recursionMax int,
) *NginxBlock {
}

func GetNginxBlocks(configContent string) *NginxBlocks {}
func main() {

}
