package main

import (
	"context"
	"os"

	"github.com/bulwarkid/virtual-fido/usbip"
	"github.com/bulwarkid/virtual-fido/util"
)

type DummyDevice struct{}

func (d *DummyDevice) HandleMessage(id uint32, onFinish func(response []byte), endpoint uint32, setupBytes []byte, transferBuffer []byte) {
}
func (d *DummyDevice) RemoveWaitingRequest(id uint32) bool     { return false }
func (d *DummyDevice) BusID() string                           { return "" }
func (d *DummyDevice) DeviceSummary() usbip.USBIPDeviceSummary { return usbip.USBIPDeviceSummary{} }

func main() {
	util.SetLogOutput(os.Stdout)
	util.SetLogLevel(util.LogLevelTrace)
	device := DummyDevice{}
	usbipserver := usbip.NewUSBIPServer([]usbip.USBIPDevice{&device})
	usbipserver.Start(context.Background())
}
