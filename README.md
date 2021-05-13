# WebHasher

Tool which makes http requests and prints the address of the request along with the
MD5 hash of the response.

## Requiremets

- The tool must be able to perform the requests in parallel so that the tool can complete sooner.
 The order in which addresses are printed is not important.
- The tool must be able to limit the number of parallel requests, to prevent exhausting local resources.
 The tool must accept a flag to indicate this limit, and it should default to 10 if the flag is not provided.
- The tool must have unit tests
- README.md must be included describing the usage of this tool.

## Build

```bash
go build -o myhttp main.go
```

## Test
```

```
## Usage
```

```
### Example
```
./myhttp -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com 

google.com 9df508f589d39855eb2152749a66b308
adjust.com 787d1722dc1a1bd33029b8b9e13a0e1e
facebook.com f96be6908ef0f29cd5694252de404fa2
yandex.com bfb468c0c97ea4f24297d11572bccd8d
twitter.com 09a1ac1c907bf7507ce10ebd4c8e7835
reddit.com/r/notfunny eb735ff74a741cbaf7e423581236a591
reddit.com/r/funny bb6ee2c311150f8782fbf7fea9deb58e
yahoo.com 54dca58e17797bf925996ba939da9791
baroquemusiclibrary.com e833349338ec1d47119fea0df2cb7558
```