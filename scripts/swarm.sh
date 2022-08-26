#!/bin/bash
local_ip=`ifconfig -a|grep inet|grep -v inet6|awk '{print $2}'|tr -d "addr:"​`

array=(`echo $local_ip | tr '\n' ' '` )

pos=0
echo "\033[032m******************************\033[0m"
echo "\033[036m*    请选择安装的IP地址        \033[0m"
echo "\033[032m******************************\033[0m"
for element in ${array[*]}
do
  let pos++
  echo "\033[0m*      ${pos} : ${element}        \033[0m"
done
echo "\033[032m******************************\033[0m"

read str

a=`uname  -a`
darwin="Darwin"
if [[ $a =~ $darwin ]];then
    IP=${array[str-1]}
else
    IP=${array[str-1]}
fi

echo "docker swarm init --advertise-addr ${IP}:2375"
docker swarm init --advertise-addr ${IP}:2375