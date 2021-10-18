package fst

import (
	"log"
	"os"
	fp "path/filepath"
)

/*
Path's job is to stay nice and neat instead of hoping string
is formatted properly. Abstracting away from the base Go path stuff.
**/
type Path struct {
	path string
}

// Set() must always store an absolute path compatibile with the standard path/filepath pkg
func (p *Path) Set(s string) {
	cleaned, err := fp.Abs(s) // Abs() also calls Clean()
	if err != nil {
		log.Fatal(err)
	}
	p.path = cleaned
}

func (p *Path) ToParent() {
	p.Set(p.Parent())
}

// String() must always keep compatibility with path/filepath pkg
func (p *Path) String() string {
	return p.path
}

// Should this return a Path? Is a string ever helpful?
func (p *Path) Parent() string {
	return fp.Dir(p.path)
}

func (p1 *Path) Copy() *Path {
	p2 := new(Path)
	p2.path = p1.path
	return p2
}

func NewPath(s string) *Path {
	p := new(Path)
	if len(s) > 0 {
		p.path = fp.Clean(s)
	} else {
		cur, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		p.path = fp.Clean(cur)
	}
	return p
}
