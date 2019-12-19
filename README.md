# BNCProject

# Decode Micro

A small microservice for querying blockchain data:

```bash
curl -XPOST localhost:1234/tx -d $txhash
```

### Build and Run

To build and run `decode-micro` run:

```bash
$ go build  -o build/decode-micro .
$ decode-micro


Usage:
  decode-micro [command]

Available Commands:
  help        Help about any command
  serve       Runs the server
  version     Prints version information

Flags:
      --config string   config file (default is $HOME/.amino-micro.yaml)
  -h, --help            help for amino-micro
  -t, --toggle          Help message for toggle

Use "amino-micro [command] --help" for more information about a command.
```