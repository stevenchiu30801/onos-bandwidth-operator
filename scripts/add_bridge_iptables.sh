#!/bin/bash
# Add iptables rules for bridges to forward packets
# Since sysctl variable net.bridge.bridge-nf-call-iptables is set to 1 by default

declare -a bridgelist=("bansaccess" "banscore")

for bridge in "${bridgelist[@]}"
do
    if ! sudo iptables -t filter -L FORWARD -v | grep $bridge >/dev/null; then
        sudo iptables -t filter -I FORWARD -i $bridge -j ACCEPT
    fi
done
