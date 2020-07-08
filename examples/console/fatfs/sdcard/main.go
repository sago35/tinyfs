package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/bgould/tinyfs/examples/console"
	"github.com/bgould/tinyfs/fatfs"
	"github.com/tinygo-org/drivers/sdcard"
)

var (
	spi    machine.SPI
	csPin  machine.Pin
	ledPin machine.Pin
)

func main() {

	sd := sdcard.New(spi, csPin)
	err := sd.Configure()
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		for {
			time.Sleep(time.Hour)
		}
	}

	filesystem := fatfs.New(&sd)

	// Configure FATFS with sector size (must match value in ff.h - use 512)
	filesystem.Configure(&fatfs.Config{
		SectorSize: 512,
	})

	console.RunFor(&sd, filesystem)
}
