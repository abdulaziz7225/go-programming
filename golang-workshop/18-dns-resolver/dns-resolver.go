package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func resolveDNS(domain string) (string, error) {
	// 1. Header
	header := []byte{
		0xAA, 0xBB, // ID
		0x01, 0x00, // Flags (standard query)
		0x00, 0x01, // Questions
		0x00, 0x00, // Answer RRs
		0x00, 0x00, // Authority RRs
		0x00, 0x00, // Additional RRs
	}

	// 2. DNS Question - encode domain as labels
	var question []byte
	parts := strings.Split(domain, ".")
	for _, part := range parts {
		question = append(question, byte(len(part)))
		question = append(question, []byte(part)...)
	}
	question = append(question, 0)    // Root label
	question = append(question, 0, 1) // Type: A record
	question = append(question, 0, 1) // Class: IN

	packet := append(header, question...)

	// 3. Connect to DNS server
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// 4. Send the DNS packet
	_, err = conn.Write(packet)
	if err != nil {
		return "", err
	}

	// 5. Receive the answer (max 512 bytes in UDP)
	response := make([]byte, 512)
	n, err := conn.Read(response)
	if err != nil {
		return "", err
	}

	// 6. Parse response
	// Skip header (12 bytes) and question section
	offset := 12

	// Skip question section by reading labels until we hit 0
	for offset < n {
		length := response[offset]
		offset++
		if length == 0 {
			break
		}
		offset += int(length)
	}
	offset += 4 // Skip type (2 bytes) and class (2 bytes)

	// Parse answer section
	// Skip name (compressed pointer: 2 bytes starting with 0xC0)
	offset += 2
	offset += 2 // Skip type
	offset += 2 // Skip class
	offset += 4 // Skip TTL

	rdLen := binary.BigEndian.Uint16(response[offset : offset+2])
	offset += 2

	if rdLen == 4 {
		ip := net.IPv4(response[offset], response[offset+1], response[offset+2], response[offset+3])
		return ip.String(), nil
	}

	return "", fmt.Errorf("invalid response")
}
