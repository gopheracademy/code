# material
Course materials for Caladan Services / GopherTrain classes

## Style Guide

All modules will have the following organization:

```
modulename/
	demos/
	exercises/
	solutions/
	module.md
```

The slides are in Markdown format, with a few extras to enable metadata being pulled from remark.js and other tools (like the `learn` application).

Slides start with metadata about the module:

```
class: center, middle
name: Slide Name
video:
description:
level: Beginner
topic: Go

# Your First PR
## Contributing to a Go OSS Project
### @bketelsen @gopheracademy

```

Each slide is delineated by the content between three dashes `---`.  Key/Value pairs come before the first H1 (#) tag in the markdown content.

Speaker notes come after three question marks: `???`.  Content that appears after three question marks, but before the next set of three dashes is for speaker view only, and won't be displayed in the slideshow or on the `learn` app display page.

You can use remark.js styling helpers to control layout of the slide:

[formatting](https://github.com/gnab/remark/wiki/Formatting)
[markdown](https://github.com/gnab/remark/wiki/Markdown)


### Injecting Code

<code src="path/relative/to/the/training/package/code.go" />
