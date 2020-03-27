#!/bin/bash
# Remove iptables rules for bridges

declare -a bridgelist=("bansaccess" "banscore")

for bridge in "${bridgelist[@]}"
do
    sudo iptables -t filter -L FORWARD -v --line-numbers | grep $bridge | cut -d " " -f 1 | sort -r | xargs -rL1 sudo iptables -t filter -D FORWARD
done
