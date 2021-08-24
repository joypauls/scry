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

func NewPath(s string) *Path {
	p := new(Path)
	if len(s) > 0 {
		p.cur = s
	} else {
		cur, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		p.cur = cur
	}
	return p
}
