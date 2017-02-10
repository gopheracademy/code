package grifts

import (
	"errors"
	"os"
	"path/filepath"

	. "github.com/markbates/grift/grift"
)

var _ = Desc("module:new", "Creates a new module with associated markdown file and folder structure.  Parameter: modulename")
var _ = Add("module:new", func(c *Context) error {
	if len(c.Args) != 1 {
		return errors.New("Please specify the module name as a filesystem safe string, eg: mymodule")
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	module := c.Args[0]
	path := filepath.Join(dir, module)
	err = os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(path, "module.md"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write([]byte(slide))
	for _, sub := range []string{"demos", "exercises", "solutions"} {
		subpath := filepath.Join(path, sub)
		err = os.Mkdir(subpath, 0755)
		if err != nil {
			return err
		}

		gf, err := os.Create(filepath.Join(subpath, ".gitsave"))
		if err != nil {
			return err
		}
		gf.Close()
	}

	return err
})

const slide = `class: center, middle
name: Slide Name
video: 
description: 
level: {Beginner, Intermediate, Advanced}
topic: {Go, Kubernetes, Docker, Web, Distributed Systems}

# Module Title Here (this is the Title Page of Presentation)

name: Slide Name
video: 
description: 
level: {Beginner, Intermediate, Advanced}
topic: {Go, Kubernetes, Docker, Web, Distributed Systems}

# Slide 

???
These are presenter notes, not shown in remark.js

---
`
