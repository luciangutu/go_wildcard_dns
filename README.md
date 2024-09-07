# Go Wildcard DNS Server

A simple wildcard DNS server written in Go. This server responds with a configurable A record for any DNS query.

### Features

- Wildcard DNS server
- Configurable A record IP address


### Prerequisites

- Docker
- Go (for building from source)

### Build the container

```shell
docker build -t wildcard-dns-server .
```

### Run the server

Run the container with the desired IP address for A records
Replace `192.168.100.1` with the IP address you want to use for A records.
The DNS server will listen on port 53 for UDP queries and respond with the specified IP address for any query.

```shell
docker run --rm --name dns-server wildcard-dns-server 192.168.100.1
```

### Example usage

Query the DNS server using dig or another DNS client (replace localhost with the container IP):

```shell
dig @localhost example.com
```

### Loadtesting

`queries.txt` content
```shell
example.com A
test.example.com A
```

Run the loadtest for 60 seconds.
```shell
dnsperf -s <server_ip_address> -d queries.txt -l 60
```
