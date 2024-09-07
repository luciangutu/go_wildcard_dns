package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

var ipAddr string

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	clientIP := w.RemoteAddr().String()
	log.Printf("Received DNS query from %s for domain %s", clientIP, r.Question[0].Name)

	msg := dns.Msg{}
	msg.SetReply(r)
	msg.Authoritative = true

	for _, q := range r.Question {
		switch q.Qtype {
		case dns.TypeA:
			rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ipAddr))
			if err != nil {
				log.Printf("Failed to generate an A record for the DNS response (%s): %v)", q.Name, err)
				continue
			}
			msg.Answer = append(msg.Answer, rr)
		}
	}

	if err := w.WriteMsg(&msg); err != nil {
		log.Printf("Failed to send DNS response to %s: %v", clientIP, err)
	}
}

func main() {
	flag.StringVar(&ipAddr, "ip", "", "IP address to return for A records")
	flag.Parse()

	if ipAddr == "" {
		log.Fatal("Please provide an IP address using the -ip flag.")
	}

	if net.ParseIP(ipAddr) == nil {
		log.Fatalf("Invalid IP address: %s", ipAddr)
	}

	listen_addr := "0.0.0.0:53"

	dns.HandleFunc(".", handleDNSRequest)
	server := &dns.Server{Addr: listen_addr, Net: "udp"}
	log.Printf("Wildcard DNS server on %s will respond with %s", listen_addr, ipAddr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start DNS server: %s", err.Error())
	}
}
