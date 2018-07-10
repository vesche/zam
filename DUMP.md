
Just random notes for sanity.

# afpacket vs libpcap

https://discuss.elastic.co/t/performance-difference-between-af-packet-libpcap/69766

```
pcap, which uses the libpcap library and works on most platforms, but it’s not the fastest option.
af__packet, which uses memory mapped sniffing. This option is faster than libpcap and doesn’t require a kernel module, but it’s Linux-specific
```