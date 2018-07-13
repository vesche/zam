# zam

This is a work in-progress, modern network analysis framework.

I probably won't be accepting major pull requests as I'm still actively laying the foundations of this project.

## So, what does this thing do?

Not much at the moment, here's a sneak peak...

```
$ git clone https://github.com/vesche/zam
$ ./zam
zam is a work in-progress, modern network analysis framework.

Usage:
  zam [command]

Available Commands:
  help        Help about any command
  interface   Interface to listen on.
  read        Read in a PCAP file.
  version     Display the current version of zam.

Flags:
      --config string   config file (default is $HOME/.zam.yaml)
  -h, --help            help for zam
  -t, --toggle          Help message for toggle

Use "zam [command] --help" for more information about a command.
$ ./zam interface en0
Listening on en0...
{"level":"info","ts":1531492220.9889936,"caller":"capture/proto.go:58","msg":"packet","src.ip":"192.168.1.69","src.port":"443","src.mac":"aa:aa:aa:bb:bb:bb","dst.ip":"1.2.3.4","dst.port":"51337","dst.mac":"cc:cc:cc:dd:dd:dd","bytes":97}
...
```