# Building Blocks of libp2p

## What is libp2p?
- A **modular toolbox** for building **peer-to-peer (P2P)** networks.  
- Lets you pick the best **protocols** for your app.  

---

### 1. **Transport** (How data moves)  
- The **transport layer** (like roads for internet data) uses:  
  - **TCP**: Reliable but slower (used for websites, emails).  
  - **UDP**: Faster but less reliable (used for video calls, games).  
  - **QUIC**: Newer protocol (built on UDP, faster for modern apps).  
- **libp2p lets you choose** the best transport for your needs.  

---

### 2. **Stream Multiplexing** (Sharing one connection)  
- Opening **many connections** wastes resources.  
- libp2p **splits one connection** into smaller **streams** (like lanes on a highway).  
  - Each stream carries **different types of data** (e.g., chat, files).  

---

### 3. **NAT Traversal** (Connecting behind routers)  
- Problem: Devices at home share **one public IP** (like a mailbox for a whole house).  
  - Hard for outsiders to **directly connect** to your device.  
- Solution: libp2p uses tricks like **hole-punching** or **relays** to help peers connect.  

---

### 4. **Protocols** (Rules for communication)  
- Each stream has a **protocol tag** (like a label on a package).  
  - Example: A `random-words` protocol sends fun words between peers.  
- Peers agree on protocols **before sending data**.  

---

### 5. **Peer Identity** (Who’s who?)  
- Every peer has a **unique ID** (like a username).  
- IDs are based on **cryptographic keys** (so they’re hard to fake).  

---

### 6. **Addressing** (Finding peers)  
- Normal addresses (like `127.0.0.1:8080`) are **hard to understand**.  
- libp2p uses **multiaddresses** (self-explaining paths):  
  - Example: `/ip4/127.0.0.1/tcp/8080/http` (clearly shows IP + port + protocol).  

---

### 7. **Security** (Private & safe connections)  
- All connections are **encrypted by default**.  
- Peer IDs **verify identities** (no imposters!).  
- *Note:* libp2p doesn’t handle **permissions** (e.g., "Can this peer delete files?").  

---

### 8. **Publish/Subscribe** (Group messaging)  
- Peers join **topics** (like chat rooms).  
  - Example: A topic named `animals` sends messages **only about animals**.  
- Subscribe to topics you care about!  

---

> ℹ️ **Learn more** in the [libp2p docs](https://docs.libp2p.io/) or later in the Launchpad course!
```

