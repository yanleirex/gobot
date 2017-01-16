package firmata

import (
	"io"

	"gobot.io/x/gobot/platforms/ble"
)

const (
	ReceiveID  = "6e400003b5a3f393e0a9e50e24dcca9e"
	TransmitID = "6e400002b5a3f393e0a9e50e24dcca9e"
)

// BLEAdaptor represents a Bluetooth LE based connection to a microcontroller running
// FirmataBLE
type BLEAdaptor struct {
	SerialPort *ble.SerialPort
	*Adaptor
}

// NewBLEAdaptor opens and uses a BLE connection to a microcontroller running
// FirmataBLE
func NewBLEAdaptor(args ...interface{}) *BLEAdaptor {
	address := args[0].(string)
	rid := ReceiveID
	wid := TransmitID

	if len(args) >= 3 {
		rid = args[1].(string)
		wid = args[2].(string)
	}

	sp := ble.NewSerialPort(address, rid, wid)
	a := NewAdaptor(address)
	a.SetName("BLEFirmata")
	a.PortOpener = func(port string) (io.ReadWriteCloser, error) {
		sp.Open()
		return sp, nil
	}

	return &BLEAdaptor{
		Adaptor:    a,
		SerialPort: sp,
	}
}
