# Redigo

Redigo is a Redis-like key-value store written in Go. It follows a single-threaded model, utilizing a single goroutine, and employs epoll on Linux and kqueue on macOS for I/O multiplexing.

## Features

- **Single-threaded:** Redigo operates using a single goroutine for efficient execution.

- **I/O Multiplexing:** Utilizes epoll on Linux and kqueue on macOS for effective I/O multiplexing, enhancing performance.

- **Key Expiry Mechanism:** Implements a Redis-like key expiry mechanism inspired by [Redis Expiry](https://redis.io/commands/expire/#:~:text=How%20Redis%20expires%20keys).

  - **Passive Expiry:** Keys are passively expired when a client attempts to access them, and the key is found to be timed out.

  - **Active Expiry:** Redis actively expires keys by periodically testing a few keys at random among those with an associated expire set. The process involves the following:

    1. Test 20 random keys from the set of keys with an associated expire.
    2. Delete all keys found to be expired.
    3. If more than 25% of keys were expired, repeat the process from step 1.

  - **Probabilistic Algorithm:** Redigo employs a probabilistic algorithm assuming that the sample of tested keys is representative of the entire key space. The expiration continues until the percentage of likely expired keys is below 25%.

  - **Memory Efficiency:** This approach ensures that, at any given moment, the maximum amount of keys already expired that are using memory is at max equal to the maximum amount of write operations per second divided by 4.

## Getting Started

Follow these steps to get started with Redigo:

1. Clone the repository: `git clone https://github.com/KavetiRohith/redigo.git`
2. Build the project: `go build -o redigo`
3. Run Redigo: `./redigo`

Feel free to explore and contribute to the project. For more details, refer to the [documentation](docs/README.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.