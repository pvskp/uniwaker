package sender

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

func CreatePacket (macAdress string) ([]byte, error) {
    repeatedMac := strings.Repeat(macAdress, 16)
    macAdressBytes, err := hex.DecodeString(repeatedMac)

    if err != nil {
        return nil, err
    }
    
    magicPacket := make ([]byte, 0, 102)
    payload, err := hex.DecodeString(strings.Repeat("FF", 6))

    magicPacket = append(magicPacket, payload...)
    magicPacket = append(magicPacket, macAdressBytes...)

    return magicPacket, nil
}

func CreateSocket (host string, port string) (net.Conn, error) {
    address := host+":"+port
    connection, err := net.Dial("upd", address) 

    fmt.Println("Socket created at", &connection)

    if err != nil {
        return nil, nil
    }

    return connection, nil
}

func SendMessage (connection net.Conn, magicPacket []byte) (string, error) {
    _, err := connection.Write(magicPacket)
    if err != nil {
        return "", err
    }

    return "Ok", nil
}
