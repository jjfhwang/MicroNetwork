```markdown
# MicroNetwork

## Description

MicroNetwork is a lightweight Go library designed to facilitate the creation and management of simple, in-memory network simulations. It provides the necessary tools to define network topologies, simulate packet transmission, and observe network behavior without the overhead of real-world network infrastructure. This library is particularly useful for testing distributed algorithms, prototyping network protocols, and educational purposes where a controlled and reproducible network environment is required.

## Features

*   **Node and Link Abstraction:** Defines `Node` and `Link` interfaces, allowing for flexible network topology creation. Implementations provided for basic node and link functionalities.
*   **Packet Simulation:** Provides a `Packet` struct to represent network packets and simulates their transmission across links. Packet loss and latency can be configured.
*   **Event-Driven Architecture:** Utilizes Go channels and goroutines to simulate asynchronous packet transmission and event handling within the network.
*   **Topology Management:** Includes functions to add nodes and links to the network, creating complex network topologies programmatically.
*   **Metrics Collection:** Offers basic functionality to collect network statistics, such as packet loss rate and average latency.

## Installation

To install MicroNetwork, you will need Go version 1.18 or later. Follow these steps:

1.  **Install Go:** If you don't have Go installed, download and install it from the official Go website: [https://go.dev/dl/](https://go.dev/dl/)

2.  **Set up your Go workspace:** If you haven't already, set up your Go workspace by defining the `GOPATH` environment variable. A common practice is to set `GOPATH` to `$HOME/go`.

3.  **Clone the repository:** Clone the MicroNetwork repository to your Go workspace:

    ```bash
    git clone https://github.com/jjfhwang/MicroNetwork.git $GOPATH/src/github.com/jjfhwang/MicroNetwork
    ```

4.  **Install dependencies:** Navigate to the MicroNetwork directory and use `go mod` to download and install dependencies:

    ```bash
    cd $GOPATH/src/github.com/jjfhwang/MicroNetwork
    go mod download
    go mod verify
    ```

5.  **Build the project:** You can build the project using the `go build` command:

    ```bash
    go build
    ```

## Usage

Here are some example code snippets demonstrating how to use MicroNetwork:

```go
package main

import (
	"fmt"
	"github.com/jjfhwang/MicroNetwork/network"
	"time"
)

func main() {
	// Create a new network
	net := network.NewNetwork()

	// Create two nodes
	node1 := network.NewBaseNode("Node1")
	node2 := network.NewBaseNode("Node2")

	// Add nodes to the network
	net.AddNode(node1)
	net.AddNode(node2)

	// Create a link between the nodes with a latency of 10ms
	link := network.NewBaseLink(node1, node2, 10*time.Millisecond)
	net.AddLink(link)

	// Send a packet from node1 to node2
	packet := &network.Packet{
		Source:      node1,
		Destination: node2,
		Data:        []byte("Hello, MicroNetwork!"),
	}

	// Simulate sending the packet
	net.SendPacket(packet)

	// Wait for the packet to be delivered (simulated)
	time.Sleep(20 * time.Millisecond)

	// Check if the packet was received
	if len(node2.ReceivedPackets) > 0 {
		receivedPacket := node2.ReceivedPackets[0]
		fmt.Printf("Node2 received packet: %s\n", string(receivedPacket.Data))
	} else {
		fmt.Println("Node2 did not receive the packet.")
	}
}

```

```go
package main

import (
	"fmt"
	"github.com/jjfhwang/MicroNetwork/network"
	"time"
)

func main() {
	// Create a network with packet loss
	net := network.NewNetwork()
	net.PacketLossRate = 0.5 // 50% packet loss

	node1 := network.NewBaseNode("Node1")
	node2 := network.NewBaseNode("Node2")

	net.AddNode(node1)
	net.AddNode(node2)

	link := network.NewBaseLink(node1, node2, 5*time.Millisecond)
	net.AddLink(link)

	packet := &network.Packet{
		Source:      node1,
		Destination: node2,
		Data:        []byte("Test Packet"),
	}

	net.SendPacket(packet)
	time.Sleep(10 * time.Millisecond)

	if len(node2.ReceivedPackets) > 0 {
		fmt.Println("Packet received successfully (despite potential loss).")
	} else {
		fmt.Println("Packet lost during transmission.")
	}
}
```

## Contributing

We welcome contributions to MicroNetwork! To contribute, please follow these guidelines:

1.  **Fork the repository:** Fork the MicroNetwork repository to your GitHub account.

2.  **Create a branch:** Create a new branch for your feature or bug fix.

3.  **Make your changes:** Implement your changes, ensuring that the code is well-documented and follows the project's coding style.

4.  **Write tests:** Write unit tests to ensure that your changes are working correctly.

5.  **Submit a pull request:** Submit a pull request to the main branch of the MicroNetwork repository.

Please ensure that your pull request includes:

*   A clear description of the changes you have made.
*   Unit tests for your changes.
*   Documentation for any new features.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/jjfhwang/MicroNetwork/blob/main/LICENSE) file for details.
```