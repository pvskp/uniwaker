package discovery

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/mostlygeek/arp"
)

type device struct {
	ip         string
	macAddress string
}

func (d device) Ip() string {
	return d.ip
}

func (d device) MacAddress() string {
	return d.macAddress
}

func GetDevices() (ipList []device) {
	// Get the local IP address
	ip := getLocalNetworkIP()

	// Get the IP subnet by removing the last octet of the IP address
	subnet := getSubnet(ip)

	// Create a wait group to synchronize the goroutines
	var wg sync.WaitGroup
	defer wg.Wait()

	// Iterate through all possible IP addresses in the subnet
	for i := 1; i < 255; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Construct the IP address
			ip := fmt.Sprintf("%s.%d", subnet, i)

			// Ping the IP address to check if it's online
			if ip != "nil" {
				if ping(ip) {
					ipList = append(ipList, device{ip, arp.Search(ip)})
				}
			}
		}(i)
	}
	return
}

// getLocalNetworkIP retrieves the local IPv4 address of the network interface that is currently up and not a loopback interface.
// It returns the IPv4 address as a string, or an empty string if no valid IPv4 address is found.
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

// getSubnet takes an IPv4 address as input and returns the subnet address by removing the last octet.
// It receives an IPv4 address as a string in the format "x.x.x.x" and returns the subnet address as a string in the format "x.x.x".
func getSubnet(ip string) string {
	// Split the IP address by dot
	parts := strings.Split(ip, ".")

	// Remove the last part (the last octet)
	parts = parts[:len(parts)-1]

	// Join the remaining parts by dot
	return strings.Join(parts, ".")
}

// ping sends an ICMP Echo Request (ping) packet to the specified IPv4 address and waits for a response.
// It takes an IPv4 address as input in the format "x.x.x.x" and returns a boolean value indicating whether
// the address responded to the ping (true) or not (false).
func ping(ip string) bool {
	// Run the "ping" command with a timeout of 1 second
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
