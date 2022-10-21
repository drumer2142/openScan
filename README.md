# openScan

### Introduction

openScan is a terminal based tool for network scanning released under the MIT License. 

### Usage

#### Flags:
- `-ip` At the moment always needs the CIDR prefix ex. `/24`
- `-p`
- `-a`
- `-sc`

# Examples

**Port scan**
```bash
./openScan -ip 192.168.1.1/24 -p
```

**Host Alive**
```bash
./openScan -ip 192.168.1.1/24 -a
```

**Scan the subnet for hosts alive**
```bash
./openScan -ip 192.168.1.1/24 -sc
```

<!-- ## License
The MIT License (MIT) - see [`LICENSE.md`](https://github.com/drumer2142/openScan/LICENSE.md) for more details -->
