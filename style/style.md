name: style
description: Style Guide
level: Beginner
topic: Training
class: center, middle
module-time: 5
# Style Guide

---
name: First Slide
# First Slide

The first slide in the deck is thrown away and used only as the title of the
module in the online training system.

However, it is used as the opening slide (remark.js) for the in person training
that is done.

---
name: module meta data
# Module Meta Data

The meta data for the module is the very first thing in the `module.md` file,
and appears before the first slide. While all meta data is saved to the
database, it is not always used in the online training system or IPT rendered
files.

```
name: style
description: Style Guide
level: Beginner
topic: Training
class: center, middle
module-time: 5
```

`name`         - Name of the module.
`description`  - Description of the module.
`level`        - Level of the module.  Should be `Beginner`, `Intermmediate`, or `Advanced`.
`topic`        - Topics are used to group content and primarily used to help put courses together quickly.
`class`        - This is the css class that `remark.js` will use. [class reference](https://github.com/gnab/remark/wiki/Markdown#class)
`module- time` - This is an estimate in `minutes` of how long this module will take to get through.

---
name: slide meta data
# Slides Meta Data

Slides have the following meta data

`name` - The name of the slide.  This will become an anchor tag in `remark.js` rendered files.

---
name: images
# Images

To include an image in your slide, use the following syntax:

```
![alternate text](/training-assets/style/demo.png)
```

`/training-assets` is the route prefix that maps to the root of the training repo.
`/style/demo.png` is the relative route in the training repo where the image exists.

![alternate text](/training-assets/style/demo.png)

---
name: code
# Code

Use the following syntax to have code injected into your presentation:

```
<code src="style/code-sample.go" />
```

<code src="style/code-sample.go" />
