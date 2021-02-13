# openup
Utility that opens port forwards via upnp.  Written in golang.

### When is this useful?
If you ever want to expose an IP address on a LAN to the internet, you can set a port forward with your router via UPnP that will port forward from the router (public IP) to the machine on the local network (192.168.x.x address).

Then you can hit the external IP address at the given port and it will get forwarded to the given machine on the local network.

## Usage
```
$ ./openup -h
Usage:
  -close
        Close (as opposed to open) the given port.
  -ip
        Display external IP address and exit.
  -port int
        Port to open/close (default -1)
  -udp
        Use UDP (instead of TCP) when opening/closing port forward.
```

### To discover your external IP:
```
$ ./openup -ip
Your external IP is: 123.123.123.123
```

### To open up a port forward:
```
$ ./openup -port 9978
Opening port forward for TCP port 9978

$ ./openup -port 9979 -udp
Opening port forward for UDP port 9979
```

### To close a port forward:
```
$ ./openup -close -port 5233 -udp
Closing port forward for UDP port 5233
``` 
