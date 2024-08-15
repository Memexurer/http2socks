package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"

	socks5 "github.com/armon/go-socks5"
	"golang.org/x/net/proxy"
)

func main() {
	// Define command-line flags
	httpProxyAddr := os.Getenv("UPSTREAM_PROXY")
	socksListenAddr := os.Getenv("LISTEN_ADDR")

	// Parse the flags
	flag.Parse()

	// Check if HTTP proxy address is provided
	if httpProxyAddr == "" {
		fmt.Println("Error: HTTP proxy address is required")
		flag.Usage()
		os.Exit(1)
	}

	// Configure your HTTP proxy
	proxyURL, err := url.Parse(httpProxyAddr)
	if err != nil {
		fmt.Printf("Error parsing HTTP proxy address: %v\n", err)
		os.Exit(1)
	}

	httpDialer, err := proxy.FromURL(proxyURL, Direct)
	if err != nil {
		panic(err)
	}

	// Create a custom dialer that uses the HTTP proxy
	dialer := func(ctx context.Context, network, addr string) (net.Conn, error) {
		return httpDialer.Dial(network, addr)
	}

	// Configure the SOCKS5 server
	conf := &socks5.Config{
		Dial: dialer,
	}

	server, err := socks5.New(conf)
	if err != nil {
		fmt.Printf("Error creating SOCKS5 server: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Starting SOCKS5 server on %s\n", socksListenAddr)
	fmt.Printf("Using HTTP proxy: %s\n", httpProxyAddr)

	// Start the SOCKS5 server
	if err := server.ListenAndServe("tcp", socksListenAddr); err != nil {
		fmt.Printf("Error starting SOCKS5 server: %v\n", err)
		os.Exit(1)
	}
}
