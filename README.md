
# Reminder

**Reminder** is a lightweight cross-platform application designed to help you manage tasks and improve your productivity. With its minimalist design and efficient features, it is the perfect tool for personal and professional task management.

## Features

- Simple and intuitive interface
- Cross-platform support
- Lightweight and fast
- CLI support for quick task creation

---

## Table of Contents

1. [Getting Started](#getting-started)
2. [Download](#download)
3. [Build](#build)
4. [Contributing](#contributing)
5. [License](#license)

---

## Getting Started

### Prerequisites

To run Reminder, ensure you have the following installed on your system:
- Go 1.21 or newer (for building the app)
- Git (to clone the repository, if building locally)

---

## Download

Prebuilt binaries for Linux, macOS, and Windows are available on the releases page.

1. Go to the [Releases](https://github.com/lrqtech/reminder/releases).
2. Download the binary for your platform:
    - **Linux**: `reminder-linux-amd64` or `reminder-linux-arm64`
    - **macOS**: `reminder-darwin-amd64` or `reminder-darwin-arm64`
    - **Windows**: `reminder-windows-amd64.exe` or `reminder-windows-arm64.exe`

3. Make the binary executable (Linux/macOS):
   ```bash
   chmod +x reminder-<platform>-<arch>
   ```

4. Run the binary:
   ```bash
   ./reminder-<platform>-<arch> --help
   ```

---

## Build

If you'd like to build Reminder from source, follow these steps:

### Clone the Repository

```bash
git clone https://github.com/lrqtech/reminder.git
cd reminder
```

### Build for Your Platform

Run the following commands to build for your platform:

```bash
go build -ldflags "-s -w"
```

---

## Contributing

Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a feature branch: `git checkout -b feature-name`.
3. Commit your changes: `git commit -m "Description of changes"`.
4. Push to your fork: `git push origin feature-name`.
5. Open a Pull Request.

---

## License

This project is licensed under the [GPL 3.0](LICENSE).

---

## Support

If you encounter any issues or have questions, feel free to [open an issue](https://github.com/lrqtech/reminder/issues).
