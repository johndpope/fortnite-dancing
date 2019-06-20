package main

import (
    "fmt"
    "flag"
    "os"


    isatty "github.com/mattn/go-isatty" // For all the terminal stuff
)

func main() {
    // This is where the dancing will happen...possibly
    // Delay the frame for a bit
    delay := flag.Int("delay for 75ms", 75, "frame delay in ms")
    orientation := flag.String("orientation", "regular")
    flag.Parse()

    if !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()){
        fmt.Sprintf(os.Stderr, "Error: %s must be ran in a terminal\n", filepath.Base(os.Args[0]))
        os.Exit(1)  // Exit out after Erroring
    }

    err := termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()

    eventQ := make(chan termbox.Event)
    go func() {
        for {
            eventQ <- termbox.PollEvent()
        }
    }()

    termbox.SetOutputMode(termbox.Output256)
    loopIdx := 0
    draw(*orientation)

}
