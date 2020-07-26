# go-k8s-b64

A very simple wrapper to base64 encode strings.

## Why

I use `envsubst` for injecting Kubernetes secrets when `kubectl apply`ing.

`echo` adds a newline by default and I always forget the `-n` argument, which ends up breaking secrets. This tool does not add a newline.

## Usage

Build the project:

```bash
go build -o b64 .
```

Place the binary in your PATH:

```bash
cp b64 ~/bin/b64
```


Invoke it with one argument, the string to encode:


```bash
b64 'foobarSECRE&A'
```

You can verify the string by piping to `base64 -d`
