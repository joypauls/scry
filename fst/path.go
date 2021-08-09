package fst

import (
	"log"
	"os"
	fp "path/filepath"
)

type Path struct {
	cur    string
	parent string
}

func (p *Path) Set(s string) {
	p.cur = fp.Clean(s)
	p.parent = fp.Dir(s)
}

func (p *Path) Cur() string {
	return p.cur
}

func (p *Path) Parent() string {
	return p.parent
}

// right now only works for starting at current directory
func InitPath() *Path {
	p := new(Path)
	cur, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	p.Set(cur)
	return p
}
