# Rumba20 Implementation in Go

## Overview
This project provides an implementation of the **Rumba20** cryptographic algorithm in Go. Rumba20 is a stream cipher derived from ChaCha20, designed for secure and efficient encryption. This implementation focuses on simplicity, portability, and ease of integration into modern Go applications.

## Features
- Implements the full Rumba20 algorithm with 20 rounds of encryption.
- Generates secure 64-byte blocks for use in stream encryption.
- Written entirely in Go for ease of integration and portability.
- Customizable key and nonce sizes for various cryptographic applications.

## Installation
To use this implementation, ensure you have Go installed. Clone the repository and run the program:

```bash
# Clone the repository
git clone https://github.com/<your-username>/rumba20-go.git
cd rumba20-go

# Run the code
go run main.go
