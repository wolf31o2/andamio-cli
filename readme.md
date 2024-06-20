# Andamio CLI

## Roadmap:

### Q1 2024:
- [x] Initialize project. Create project structure and share at Live Coding
- [x] Outline possible project features

### Q2 2024:
- [x] Local node configuration in .env: network and
- [x] Simple metadata writer
- [x] Querying Andamio Network (see Blog Post)
- [ ] Setting up own Andamio index (see Blog Post)
- [ ] Wallet configuration in .env: network and keys

### Q3 2024:
- [ ] Reference script deployment
- [ ] Custom data serialization for Andamio transactions
- [ ] R&D: Querying a side-chain

### Q 2024:
- [ ] Andamio Public API MVP

## Using Andamio CLI

### Config
- [ ] Add config as command line argument
Create a .env file like this:
```bash
CARDANO_NODE_MAGIC="1"
CARDANO_NODE_SOCKET_PATH="<PATH TO>/node.socket"
```


## Credits

Thank you to Blink Labs for sharing examples.
- https://github.com/blinklabs-io/adder-library-starter-kit/tree/main is used in `/sync`