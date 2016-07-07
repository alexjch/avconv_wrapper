package avconv_wrapper

import (
    "fmt"
    "log"
    "time"
    "errors"
    "os/exec"
)

type AVConv struct {}

const (
    EXEC_TIMEOUT = 10
)

var TIMEOUT_MSG = fmt.Sprintf("Transcoding timed out after %d seconds", EXEC_TIMEOUT)

// Run the wrapper to transcode inFile to format provided in outFrmt parameter
func (avc *AVConv) Run(rate int, inFile, outFile, outFrmt string) error {
    avconvCmd := "/usr/bin/avconv"
    avconvArgs := []string{"-i", inFile, fmt.Sprintf("%s.%s", outFile, outFrmt)}
    log.Println("Executing:", avconvCmd, "with arguments:", avconvArgs)
    avconv := exec.Command(avconvCmd, avconvArgs...)
    if err := avconv.Start(); err != nil {
        return err
    }

    go func(){
        time.Sleep(EXEC_TIMEOUT * time.Second)
        if avconv.ProcessState.Success() != true {
            log.Println(TIMEOUT_MSG)
            avconv.Process.Kill()
        }
    }()

    if err := avconv.Wait(); err != nil {
        if avconv.ProcessState.Exited() {
            return errors.New("avconv invocation failed")
        } else {
            return errors.New(TIMEOUT_MSG)
        }
    }

    return nil
}
