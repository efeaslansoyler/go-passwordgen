# Go Password Generator

[![en](https://img.shields.io/badge/lang-en-red.svg)](README.md)
[![tr](https://img.shields.io/badge/lang-tr-blue.svg)](README_TR.md)

A flexible and secure command-line password generator written in Go.

## Features

- Generate passwords with customizable length
- Include/exclude special characters
- Include/exclude numbers
- Include/exclude uppercase letters
- Include/exclude lowercase letters
- Generate multiple passwords at once
- Password strength analysis (Excellent, Strong, Moderate, Weak)
- Entropy calculation and display
- Generation time display
- Version information
- Easy-to-use command-line interface

## Installation

```bash
go install github.com/efeaslansoyler/go-passwordgen@latest
```

## Usage

Basic usage:
```bash
go-passwordgen
```

This will generate a 12-character password with all character types included.

### Options

- `-l, --length`: Set password length (default: 12)
- `-s, --special`: Include special characters (default: true)
- `-n, --numbers`: Include numbers (default: true)
- `-u, --upper`: Include uppercase letters (default: true)
- `-o, --lower`: Include lowercase letters (default: true)
- `-c, --count`: Number of passwords to generate (default: 1)
- `-q, --quiet`: Suppress output (print only password(s)) (default: false)
- `-v, --version`: Display version information

### Examples

Generate a 16-character password:
```bash
go-passwordgen -l 16
```

Generate 5 passwords:
```bash
go-passwordgen -c 5
```

Generate a password without special characters:
```bash
go-passwordgen --special=false
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

Efe Aslan Söyler (efeaslan1703@gmail.com)
