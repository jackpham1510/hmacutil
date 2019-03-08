## Hmacutil

Small HMAC utilities for Golang

## Download

```bash
go get github.com/tiendung1510/hmacutil
```

## Features

- Computes HMAC using secret key
- Hex string encoded
- Base64 encoded
- Algorithm supports:
  - MD5
  - SHA1
  - SHA256
  - SHA512

## Examples

```go
// Computes HMAC
mac := hmacutil.Encode(hmacutil.SHA256, "secret_key", "data")

// Computes HMAC in hex string encoded
mac := hmacutil.HexStringEncode(hmacutil.SHA256, "secret_key", "data")

// Computes HMAC in base64 encoded
mac := hmacutil.Base64Encode(hmacutil.SHA256, "secret_key", "data")
```

## License

[MIT](LICENSE)