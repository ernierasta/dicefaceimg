# About diceimg

This is golang library designed to produce dice face images.
At the first stage it will generate D6 dice faces. Maybe I will other dices.

## Why?

It is used in tobytables project. I wanted to be able to have some textual 
representation of a dice in source and graphical dice with this value in the
output.

## How to use

```golang
width := 10
height := 10
strokeWidth := 0.2
radius := 0.5

d := NewBlank(width+strokeWidth, height+strokeWidth, strokeWidth, radius, canvas.Gray, canvas.Salmon, canvas.Lightgray)
d.D6(3, width, height, false)
d.WriteOutput("myfile.svg")
```
You can generate jpg, png, pdf giving propriate extension to WriteOutput filename.
