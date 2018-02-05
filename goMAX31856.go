/*
Copyright (c) 2018 Forrest Sibley <My^Name^Without^The^Surname@ieee.org>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package goMAX31856

import (
	//"sync"

	"github.com/the-sibyl/piSPI"
)


type MAX31856 struct {
	dev *spi.Device
	spidevPath string
	spiClockSpeed int64
}

// Add functionality for DRDY pin
func Setup(spidevPath string, spiClockSpeed int64) (MAX31856, error) {

	m := MAX31856{
		spidevPath: spidevPath,
		spiClockSpeed: spiClockSpeed,
	}

	o := spi.Devfs{
		Dev:	m.spidevPath,
		Mode:	spi.Mode1,
		MaxSpeed: m.spiClockSpeed,
	}

	dev, err := spi.Open(&o)

	if err != nil {
		return m, err
	}

	m.dev = dev

	return m, nil
}

// Reset the faults register
func (m *MAX31856) ResetFaults() {

}

// Intended to be placed into a Goroutine
// Singleton behavior
func (m *MAX31856) GetTempAuto() {

}

// Intended to be called once per measurement
func (m *MAX31856) GetTempOnce() {

}

// Internal function to get temperature. Return a float32 containing the temperature in degrees centigrade.
func getTemp(dev *spi.Device) float32 {

	readValue := make([]byte, 4)

	// Read 0xC, 0xD, 0xE. The address auto-increments on the chip.
	dev.Tx([]byte{
	0xC, 0x0, 0x0, 0x0,
	}, readValue)

	// Discard the first byte, save the rest, and shift them to their proper positions. The data are in two's
	// complement, and the math here works out nicely.
	temp := int16(readValue[1]) << 8 | int16(readValue[2])
	linearTempDegC := float32(temp) * 0.0625

	return linearTempDegC
}
