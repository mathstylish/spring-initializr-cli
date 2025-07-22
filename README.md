# 📦 initzr

An interactive CLI for quickly start Spring Boot projects — right from the terminal!

---

## ✨ Features

- Choose project type: Maven or Gradle (Groovy or Kotlin)
- Language selection: Java, Kotlin, Groovy
- Spring Boot version selection
- Interactive dependency filtering
- Define project metadata
- Project generation based on [Spring Initializr](https://start.spring.io)

---

## 🚀 Installation

### Install from package

Pre-built packages for Windows, macOS, and Linux are found on the <u>[Releases](https://github.com/mathstylish/initzr/releases)</u> page. After that, place the executable in your operating system's PATH.

### Using go

```bash
go install github.com/mathstylish/initzr@latest
```

### Usage

```bash
initzr
```

You will be guided by prompts to set up your project step by step.

![CLI usage demo](./assets/usage_example.gif)

The project will be downloaded and extracted to the directory where you ran the `initzr` command. The directory name is the same as the one defined in the project **artifact**.
