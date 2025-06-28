# WattSense 2.0 Backend

Welcome to the backend repository for **WattSense 2.0**, an AI-powered Smart Home Energy Management System. This service is responsible for handling data ingestion, AI-driven analytics, and providing a robust API for the frontend and other clients.

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [API Reference](#api-reference)
- [Getting Started](#getting-started)
- [Environment Variables](#environment-variables)
- [Project Structure](#project-structure)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

---

## Overview

WattSense 2.0 backend is a Golang-based web server that enables users to upload energy consumption data, receive AI-driven analysis and recommendations, and interact with a conversational AI assistant for smart energy management.

- **Data Input**: CSV files containing timestamped appliance-level energy usage.
- **AI Integration**: Leverages HuggingFace models (`tapas-base-finetuned-wtq` for tabular data, and optionally chat models such as `microsoft/Phi-3.5-mini-instruct`).
- **RESTful API**: Exposes endpoints for file upload, data analysis, and chatbot interaction.

---

## Features

- **CSV Data Upload**: Accepts household energy usage data in a standardized format.
- **AI Analytics**: Predicts future energy consumption and provides personalized saving recommendations.
- **Chatbot**: Answers natural language questions about your data and energy usage.
- **Modular Design**: Easily extend or swap AI models and services.

---

## Architecture

- **main.go**: API server and route definitions.
- **model/model.go**: Data structures for CSV records, chat messages, and AI responses.
- **repository/fileRepository.go**: File handling utilities (upload, read, write).
- **service/file_service.go**: File processing and CSV parsing.
- **service/ai_service.go**: AI model interaction (table QA, chatbot).

---

## API Reference

### `POST /upload`
Upload a CSV file containing energy usage data. The backend parses and analyzes the file; returns summary statistics and a preview.

- **Request**: `multipart/form-data` with a file field.
- **Response**:  
  ```json
  {
    "status": "success",
    "summary": {
      "totalEnergy": 123.4,
      "topAppliances": [...],
      ...
    }
  }
  ```

### `POST /chat`
Ask questions about your household energy data or request recommendations.

- **Request**:  
  ```json
  {
    "context": "optional previous chat context",
    "query": "How much energy did I use last week?"
  }
  ```
- **Response**:  
  ```json
  {
    "status": "success",
    "answer": "You used 45 kWh last week."
  }
  ```

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/) 1.18+
- (Recommended) [Node.js](https://nodejs.org/) for frontend integration/testing
- HuggingFace API Token for AI model access

### Installation

1. **Clone the repository**  
    ```bash
    git clone https://github.com/DumbiFadhil/WattSense2.0.git
    cd WattSense2.0
    ```

2. **Install Go dependencies**  
    ```bash
    go mod tidy
    ```

3. **Configure environment variables**  
    Copy `.env.example` to `.env` and set your HuggingFace API token and other configs.

4. **Run the server**  
    ```bash
    go run main.go
    ```
    The API will be available at `http://localhost:8080`.

---

## Environment Variables

Create a `.env` file in the root directory:

```
HUGGINGFACE_TOKEN=your_huggingface_token_here
PORT=8080
```

---

## Project Structure

```
WattSense2.0/
├── main.go
├── model/
│   └── model.go
├── repository/
│   └── fileRepository.go
├── service/
│   ├── ai_service.go
│   └── file_service.go
├── test/
│   └── (unit and integration tests)
├── data-series.csv
├── .env.example
├── go.mod
└── README.md
```

---

## Testing

- **Unit tests** are located in the `test/` directory and alongside service files.
- Run all tests with:
    ```bash
    go test ./...
    ```

---

## Contributing

1. Fork this repository.
2. Create your feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a Pull Request.

Follow Go best practices and ensure all tests pass before submitting.

---

## License

Distributed under the MIT License. See `LICENSE` for details.

---

## Contact

For issues, suggestions, or support, please open an issue or contact [DumbiFadhil](https://github.com/DumbiFadhil).

---
