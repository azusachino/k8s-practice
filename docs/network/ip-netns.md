# IP NetNS

```sh
man ip-netns

# network namespace is logically another copy of the network stack, with its own routes, firewall rules, and network devices.
```

## Enter new NetNS

```sh
sudo ./inspect_net_stack.sh
> Network devices
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00

> Route table

> Iptables rules
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT
```

![.](https://iximiuz.com/container-networking-is-simple/network-namespace-4000-opt.png)

## The magic weapon - `veth`

```sh
man veth
# veth devices are virtual Ethernet devices. They can act as tunnels between network namespaces to create a bridge to a physical network device in another namespace, but can also be used as standalone network devices.
```

## Setup one network namespace

```sh
# From root namespace
$ sudo ip netns add netns1
$ sudo ip link add veth1 type veth peer name ceth1
$ sudo ip link set ceth1 netns netns1
$ sudo ip link set veth1 up
$ sudo ip addr add 172.18.0.21/16 dev veth1

$ sudo nsenter --net=/var/run/netns/netns1
$ ip link set lo up
$ ip link set ceth1 up
$ ip addr add 172.18.0.20/16 dev ceth1
```

## References

- [How Container Networking Works: Practical Explanation](https://iximiuz.com/en/posts/container-networking-is-simple/)
