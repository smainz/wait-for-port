package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// testConnection tries to establish a TCP connection to the specified host and port
// within the given timeout duration.
func testConnection(host, port string, timeout time.Duration) error {
	// Formulate the target address by combining the IP address and port
	target := fmt.Sprintf("%s:%s", host, port)

	// Record the start time for measuring timeout
	startTime := time.Now()

	// Attempt to establish a TCP connection until the timeout is reached
	for time.Since(startTime) < timeout {
		conn, err := net.DialTimeout("tcp", target, 100*time.Millisecond)
		if err == nil {
			defer conn.Close()
			fmt.Println("Connected to", target)
			return nil
		}
		fmt.Println("Connection attempt failed:", err)
		time.Sleep(time.Second) // Wait before the next attempt
	}

	// If timeout is reached without a successful connection, return an error
	return fmt.Errorf("Connection attempts to %s timed out after %ds", target, timeout/time.Second)
}

func main() {
	// Create a new CLI app instance
	app := &cli.App{
		Name:  "TCP Connection Tester",
		Usage: "Test TCP connection with timeout",
		Flags: []cli.Flag{
			// Define command-line flags for host, port, and timeout
			&cli.StringFlag{
				Name:    "host",
				Value:   "127.0.0.1",
				Usage:   "IP address or DNS name of host to connect to",
				EnvVars: []string{"PLUGIN_HOST"}, // Use environment variable if provided
			},
			&cli.StringFlag{
				Name:    "port",
				Value:   "8080",
				Usage:   "Port to connect to",
				EnvVars: []string{"PLUGIN_PORT"}, // Use environment variable if provided
			},
			&cli.DurationFlag{
				Name:    "timeout",
				Value:   30 * time.Second,
				Usage:   "Timeout duration for connection attempts",
				EnvVars: []string{"PLUGIN_TIMEOUT"}, // Use environment variable if provided
			},
		},
		Action: func(c *cli.Context) error {
			// Retrieve values from command-line flags
			host := c.String("host")
			port := c.String("port")
			timeout := c.Duration("timeout")

			// Call the testConnection function to attempt the connection
			err := testConnection(host, port, timeout)
			return err
		},
	}

	// Run the CLI app
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // Exit with failure code
	}
}
