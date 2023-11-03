# PasswordGen

**PasswordGen** is a simple command-line tool written in Go that generates random passwords.

## Table of Contents

- [Overview](#overview)
- [Usage](#usage)
- [Installation](#installation)
- [Contributing](#contributing)
- [License](#license)

## Overview

This Go application utilizes the crypto/rand package to generate cryptographically secure random passwords. The generated passwords can be customized by specifying the length and character sets used.

## Usage

To generate a random password, you can use the following command:

```bash
PasswordGen
```

### Options

- `-l`: Specifies the length of the generated password. Default is 12.
- `-d`: Includes digits to password generated.
- `-s`: Include special characters in the character set.

### Examples

Generate a password with a length of 16 characters using only lowercase letters:

```bash
PasswordGen -l 16 -s
```

Generate a password with a length of 20 characters using alphanumeric and special characters:

```bash
passwordgen -l 20 -d -s
```

## Installation

To install PasswordGen, you can use `go get`:

```bash
go get github.com/Unintellectual/PasswordGen
```

## Contributing

If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your branch to your fork.
5. Open a pull request.
