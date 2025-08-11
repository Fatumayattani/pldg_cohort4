package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/libp2p/go-libp2p"
)

func main() {
	// Create a basic libp2p node with Windows-friendly defaults
	node, err := libp2p.New(
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/9000",    // TCP port 9000
			"/ip6/::/tcp/9000",         // IPv6 support
		),
	)
	if err != nil {
		log.Fatalf("❌ Failed to create node: %v", err)
	}
	defer node.Close()

	// Print node details
	fmt.Printf("✅ Node ID: %s\n", node.ID())
	fmt.Println("📡 Listening on:")
	for _, addr := range node.Addrs() {
		fmt.Printf("  %s/p2p/%s\n", addr, node.ID())
	}

	// Wait for CTRL+C
	fmt.Println("\n🛑 Press CTRL+C to shutdown...")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}