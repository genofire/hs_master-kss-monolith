#!/bin/bash
host="HOST_FOR_$1"
port="PORT_FOR_$1"
remote="circleci@${!host}"
echo "deploying on: $remote"

scp -p $port bin/stock $remote:~/bin/stock;
rsync -e "ssh -p $port $host" -a webroot/ $remote:~/lib/stock/www;
rsync -e "ssh -p $port $host" -a contrib/ $remote:~/lib/stock/contrib;
ssh -p $port $remote sudo systemctl restart stock;
