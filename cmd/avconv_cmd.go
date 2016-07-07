package main

import (
    "os"
    "log"
    "flag"
    "errors"
    "github.com/alexjch/avconv_wrapper"
)

func parseArgs()(int, *string, *string, *string, error){

    var rate = flag.Int("rate", 4800, "Sample rate")
    var inFile = flag.String("inFile", "", "File source to transcode")
    var outFile = flag.String("outFile", "", "File destination to transcode")
    var outFormat = flag.String("outFormat", "flac", "Output file extension")
    flag.Parse()

    log.Println(*inFile, *outFile, *outFormat)

    if *inFile == "" || *outFile == "" {
        return 0, nil, nil, nil, errors.New("input and output arguments should have a value")
    }

    return *rate, inFile, outFile, outFormat, nil
}

func main() {

    rate, inFile, outFile, outFormat, err := parseArgs()

    if err != nil {
        log.Println(err.Error())
        os.Exit(1)
    }

    avconv := avconv_wrapper.AVConv{}
    if err := avconv.Run(rate, *inFile, *outFile, *outFormat); err != nil {
        log.Println(err.Error())
    }

}
