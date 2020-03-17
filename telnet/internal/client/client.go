package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// Client TCP client
type Client struct {
	conn       net.Conn
	ctx        context.Context
	cancelFunc context.CancelFunc
	options    *ConnectOptions
}

type ConnectOptions struct {
	Host    string
	Port    string
	Timeout time.Duration
}

func NewClient(options *ConnectOptions) (*Client, error) {
	dialer := &net.Dialer{}
	dialer.Timeout = options.Timeout

	ctx, cancelFunc := context.WithCancel(context.Background())
	conn, err := dialer.DialContext(ctx, "tcp", net.JoinHostPort(options.Host, options.Port))
	if err != nil {
		cancelFunc()
		return nil, err
	}

	fmt.Printf("Connected to %s:%s...\n", options.Host, options.Port)

	return &Client{conn, ctx, cancelFunc, options}, nil
}

func (client *Client) Run() {
	inputData := make(chan string)
	serverData := make(chan string)
	fmt.Println("Press Ctrl+D for exit")

	go client.readInputData(inputData)
	go client.readServerData(serverData)

	for {
		select {
		case data := <-inputData:
			if _, err := client.conn.Write([]byte(data)); err != nil {
				fmt.Printf("Failed to write: %s\n", err)
				return
			}
		case data := <-serverData:
			fmt.Println(data)
		case <-client.ctx.Done():
			return
		}
	}
}

func (client *Client) readServerData(out chan<- string) {
	scanner := bufio.NewScanner(client.conn)
	for scanner.Scan() {
		out <- scanner.Text()
	}

	client.cancelFunc()
}

func (client *Client) readInputData(out chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			// EOF (Ctrl+D)
			if err != io.EOF {
				log.Fatalln("Failed to scan stdin")
			}
			break
		}
		out <- text
	}

	client.cancelFunc()
}

func (client *Client) Close() {
	_ = client.conn.Close()
}
