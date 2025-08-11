# Connecting Peers in libp2p

## Understanding Nodes
- **libp2p** is modular: Use it for any P2P app, even without IPFS.  
- A **node** = a peer (any device sending/receiving messages in the network).  

### Example: Creating a Node (Go)
```go
import "github.com/libp2p/go-libp2p"

func main() {
    node, err := libp2p.New() // Creates a default node
    if err != nil { panic(err) }
}
```

---

## Peer Identity
- Each node has a unique **peer ID**, derived from its **public key**.  
- Peer IDs are **multihashes** (hashes with algorithm prefixes).  

### How Peer IDs Are Generated:
1. **If public key > 42 bytes**: Hash it.  
2. **If ≤ 42 bytes**: Use the key directly (no hashing).  

---

## Multiaddresses
- **Locate peers** by specifying:  
  - **Transport** (TCP/UDP/QUIC).  
  - **Location** (IP/DNS).  

### Example:
`/ip4/127.0.0.1/tcp/8080`  
*(Connects to IPv4 `127.0.0.1` via TCP port `8080`)*  

---

## Establishing a Connection
1. **Handshake**:  
   - Peers verify they understand `multistream-select` (protocol negotiator).  
   ```plaintext
   > /multistream/1.0.0  # "Do you understand this?"
   + /multistream/1.0.0  # "Yes!"
   ```

2. **Security Negotiation**:  
   - Peers agree on encryption (e.g., TLS or Noise).  
   ```plaintext
   > /tls/1.0.0    # "TLS supported?"
   + na            # "No."
   > /noise/1.0.0  # "Noise supported?"
   + /noise/1.0.0  # "Yes!"
   ```

3. **Multiplexing**:  
   - Split one connection into **streams** (like lanes on a highway).  
   - Each stream handles a different protocol (e.g., DHT, ping).  

---

## Peer Discovery
- **Advertiser**: Shares protocols it supports.  
- **Discoverer**: Finds peers.  
- **Methods**:  
  - **mDNS**: For local networks.  
  - **Kademlia DHT**: Used in IPFS (global discovery).  
