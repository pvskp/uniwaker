package main

import (
	// "uniwaker/discovery"
	"uniwaker/ui"
)

func main() {
	ui.Start()
}

// func main() {
// 	devices := discovery.GetDevices()
// 	for _, device := range devices {
// 		fmt.Printf("New device found!\n")
// 		fmt.Printf("IP: %s\n", device.Ip())
// 		fmt.Printf("MacAddress: %s\n", device.MacAddress())
// 		fmt.Printf("------------------------------------\n")
// 		fmt.Printf("\n")
// 	}
// }

// func main () {
//     // destIP := "192.168.0.115"
//     macAdress := "d8:e0:e1:00:1e:d5"
//     parsedMac := strings.Replace(macAdress, ":", "", -1)
//     magicPacket, err := sender.CreatePacket(parsedMac)

//     if err != nil {
//         panic (err)
//     }

//     socket, err := net.Dial("udp", "192.168.0.255:9")

//     if err != nil {
//         panic (err)
//     }

//     var n int

//     for i := 0; i < 5; i++ {
//         n, err = socket.Write(magicPacket)
//         time.Sleep(1 * time.Second)
//     }

//     if err != nil {
//         panic (err)
//     }

//     fmt.Println(n)
// }
