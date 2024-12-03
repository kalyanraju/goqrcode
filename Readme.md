
# QR Code Generator API 📦🚀

Generate stunning QR codes effortlessly with this API built using **Go** and **Gin**. Customize content and size, and receive a high-quality PNG QR code in response. Perfect for integration into websites, mobile apps, or any project requiring QR code functionality.

---

## 🌟 Features
- **Customizable Content**: Generate QR codes with any text or URL.
- **Dynamic Sizing**: Specify the QR code dimensions to fit your needs.
- **Blazing-Fast**: Powered by Go's high-performance capabilities.
- **Lightweight**: Minimal dependencies for quick deployment.
- **Dockerized**: Ready-to-deploy with a simple Docker setup.

---

## 🚀 Getting Started

### Prerequisites
- **Go** (>=1.20)
- **Docker** (optional, for containerized deployment)

---

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/kalyanraju/goqrcode
   cd goqrcode
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

---

### Generate a QR Code

- **Endpoint**: `POST /generate`
- **Payload**:
  ```json
  {
      "content": "Hello, QR Code!",
      "size": 256
  }
  ```
- **Response**: PNG QR code image.

Example using `curl`:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"content":"Hello, QR Code!","size":256}' http://localhost:8080/generate --output qrcode.png
```

---

## 📦 Docker Deployment

1. Build the Docker image:
   ```bash
   docker build -t qr-code-generator .
   ```

2. Run the container:
   ```bash
   docker run -d -p 8080:8080 qr-code-generator
   ```

3. Access the API at `http://localhost:8080`.

---

## 🛠 Tech Stack
- **Language**: Go
- **Framework**: Gin
- **QR Code Library**: [go-qrcode](https://github.com/skip2/go-qrcode)
- **Containerization**: Docker

---

## 🤝 Contributing
Contributions are welcome! Here's how you can help:
1. Fork the repository.
2. Create a feature branch: `git checkout -b feature-name`.
3. Commit your changes: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature-name`.
5. Open a pull request.

---

## ⭐️ Star This Repository
If you find this project helpful, please give it a star ⭐️ to show your support!

---

## 📜 License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## ✨ Support
Have questions or need help? Open an issue or reach out to the [maintainer](mailto:kalyan2642@gmail.com).

---

## 🎉 Acknowledgements
- Inspired by the simplicity of **go-qrcode**.
- Built with ❤️ and Go!
