package sender

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strings"
)

// createPacket creates a magic packet byte slice from a given MAC address.
// The magic packet is used for waking up a remote device from sleep or power-off state.
// It consists of a payload of 6 bytes (0xFF repeated 6 times) followed by 16 repetitions
// of the MAC address bytes.
// The resulting magic packet byte slice can be sent over the network to wake up the remote device.
func createPacket(macAddress string) ([]byte, error) {
	repeatedMac := strings.Repeat(macAddress, 16)
	macAddressBytes, err := hex.DecodeString(repeatedMac)

	if err != nil {
		return nil, err
	}

	magicPacket := make([]byte, 0, 102)
	payload, err := hex.DecodeString(strings.Repeat("FF", 6))

	magicPacket = append(magicPacket, payload...)
	magicPacket = append(magicPacket, macAddressBytes...)

	return magicPacket, nil
}

// createSocket creates a UDP socket connection to the specified host and port.
// The returned net.Conn interface can be used to send data over the network.
// The host parameter should be a hostname or an IP address, and the port parameter
// should be a string representing the port number.
// Example usage: connection, err := CreateSocket("192.168.1.100", "9")
func createSocket(host string, port string) (net.Conn, error) {
	address := host + ":" + port
	connection, err := net.Dial("udp", address)

	fmt.Println("Socket created at", &connection)

	if err != nil {
		return nil, err
	}

	return connection, nil
}

func createMagicSocket(host string) (net.Conn, error) {
	return createSocket(host, "9")
}

// SendMessage sends the magic packet byte slice over the network using the provided
// net.Conn connection. It returns true if the message was sent successfully, otherwise false.
func SendMessage(macAddress, host string) (err error) {
	connection, err := createMagicSocket(host)
	magicPacket, err := createPacket(macAddress)
	if err != nil {
		err = errors.New("Could not create socket to host")
	}
	_, err = connection.Write(magicPacket)
	if err != nil {
		err = errors.New("Could not send Magic Packet")
	}
	return
}
