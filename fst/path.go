package fst

import (
	"log"
	"os"
	fp "path/filepath"
)

/*
Path's job is to stay nice and neat instead of hoping string
is formatted properly.
**/
type Path struct {
	cur string
}

func (p *Path) Set(s string) {
	p.cur = fp.Clean(s)
}

func (p *Path) String() string {
	if len(p.cur) > 1 {
		return p.cur + string(fp.Separator)
	}
	return p.cur
}

// should this return a Path??
func (p *Path) Parent() string {
	return fp.Dir(p.cur)
}

// right now only works for starting at current directory
func NewPath() *Path {
	p := new(Path)
	cur, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	p.Set(cur)
	return p
}
