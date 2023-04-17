package discovery

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	// "sync"
)

// func main() {
// 	// Get the local IP address
// 	ip := getLocalNetworkIP()

// 	// Get the IP subnet by removing the last octet of the IP address
//   fmt.Println(ip)
// 	subnet := getSubnet(ip)

// 	// Create a wait group to synchronize the goroutines
// 	var wg sync.WaitGroup

// 	// Iterate through all possible IP addresses in the subnet
// 	for i := 1; i < 255; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()

// 			// Construct the IP address
// 			ip := fmt.Sprintf("%s.%d", subnet, i)

// 			// Ping the IP address to check if it's online
//       if ip != "nil" {
//         if ping(ip) {
//           fmt.Println("Device found:", ip)
//           // fmt.Print(getVendor(ip))
//         }
//       }
// 		}(i)
// 	}

// 	// Wait for all goroutines to finish
// 	wg.Wait()
// }


func getLocalNetworkIP() string {
	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Failed to get network interfaces:", err)
		os.Exit(1)
	}

	// Iterate through the network interfaces
	for _, iface := range interfaces {
		// Skip loopback and down interfaces
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		// Get the addresses of the current interface
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Failed to get addresses for interface", iface.Name, ":", err)
			continue
		}

		// Iterate through the addresses
		for _, addr := range addrs {
			// Convert the address to a net.IPNet
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			// Check if the address is an IPv4 address
			if ipnet.IP.To4() != nil {
				// Return the IPv4 address as a string
				return ipnet.IP.String()
			}
		}
	}
  return ""
}

func getSubnet(ip string) string {
	// Split the IP address by dot
	parts := strings.Split(ip, ".")

	// Remove the last part (the last octet)
	parts = parts[:len(parts)-1]

	// Join the remaining parts by dot
	return strings.Join(parts, ".")
}

func ping(ip string) bool {
	// Run the "ping" command with a timeout of 1 second
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}