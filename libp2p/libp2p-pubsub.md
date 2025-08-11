# libp2p PubSub System Guide

## What is PubSub?
- Decentralized messaging system where:
  - **Publishers** send messages to topics
  - **Subscribers** receive only messages from topics they follow
- Example: Chat app where users subscribe to "#music" to get music-related messages

## Key Implementations

### FloodSub
- **How it works**: 
  - Messages flood through all connected peers
  - Simple but bandwidth-heavy (duplicate messages)
- **Diagram**:
  ```mermaid
  graph LR
    P1[Publisher] --> P2[Peer]
    P1 --> P3[Peer]
    P2 --> P4[Subscriber]
    P3 --> P4
  ```

### Gossipsub (Improved Version)
- **Optimizations**:
  - Two connection types:
    - **Full-message**: Receives all data
    - **Metadata-only**: Receives only message IDs
  - Dynamic adjustment:
    - **Grafting**: Upgrade to full-message when interested
    - **Pruning**: Downgrade to metadata-only when not
- **Benefits**: 50-70% less bandwidth than FloodSub

## Core Concepts
| Term | Description |
|------|-------------|
| Topic | Category/channel for messages (e.g., "/news") |
| Network Degree (D) | Target number of full-message connections |

## Use Cases
1. Decentralized chat apps
2. IoT device coordination
3. Blockchain event notifications

## Resources
- [libp2p PubSub Docs](https://pl-launchpad.io/curriculum/libp2p/pubsub/)
```

