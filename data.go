package main

import "github.com/nsf/termbox-go"

const num_frames = 10

var frames = [num_frames]string{
	``,
}

var colors = [num_frames]termbox.Attribute{
	// approx colors from original gif
	termbox.Attribute(210), // peach
	termbox.Attribute(222), // orange
	termbox.Attribute(120), // green
	termbox.Attribute(123), // cyan
	termbox.Attribute(111), // blue
	termbox.Attribute(134), // purple
	termbox.Attribute(177), // pink
	termbox.Attribute(207), // fuschia
	termbox.Attribute(206), // magenta
	termbox.Attribute(204), // red
}
