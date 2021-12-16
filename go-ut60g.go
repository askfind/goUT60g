/*
 * This application can be used to read frame UT60G serial port
 */
package main

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

func main() {

	// Open the first serial port detected at 19200bps O71
	mode := &serial.Mode{
		BaudRate: 19200,
		Parity:   serial.OddParity,
		DataBits: 7,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	err1 := port.SetDTR(true)
	if err1 != nil {
		log.Fatal(err)
	}
	fmt.Println("Set DTR ON")

	err2 := port.SetRTS(false)
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println("Set DTR OFF")

	status, err3 := port.GetModemStatusBits()
	if err3 != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status: %+v\n", status)

	port.ResetInputBuffer()

	// Read and print the response
	buff := make([]byte, 20)
	var index int = 0
	for {
		// Reads up to 100 bytes
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			//break
		} else {
			//fmt.Printf("%v", string(buff[:n]))
			//fmt.Printf("[%d] %s", n, hex.Dump(buff[:n]))

			for i := 0; i < n; i++ {
				//  <CR> or <LF>
				if buff[i] == 0x0d {
					fmt.Println("")
					index = 0
				} else if buff[i] == 0x0a {
					fmt.Println("")
					index = 0
				} else {
					//	Print digit3,digit2,digit1,digit0
					//fmt.Println(index, buff[i])
					switch index {
					case 1:
						fmt.Printf("%c", buff[i])
						break
					case 2:
						fmt.Printf("%c", buff[i])
						break
					case 3:
						fmt.Printf("%c", buff[i])
						break
					case 4:
						fmt.Printf("%c", buff[i])
						break
					default:
						break
					}
					index += 1
				}
			}
		}
	}

}
