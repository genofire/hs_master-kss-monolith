#!/bin/bash
host=$1
port=$2
remote="circleci@${host}"
echo "deploying..."
ssh -p $port $remote sudo systemctl stop stock;
RETVAL=$?
[ $RETVAL -ne 0 ] && exit 1
scp -q -P $port ~/.go_workspace/bin/stock $remote:~/bin/stock;
RETVAL=$?
# sed -ie 's/http:\/\/localhost:8080/https:\/\/stock.pub.warehost.de/g' webroot/static/js/main.js
# [ $RETVAL -eq 0 ] && RETVAL=$?
rsync -e "ssh -p $port" -a webroot/ $remote:~/lib/stock/www;
[ $RETVAL -eq 0 ] && RETVAL=$?
rsync -e "ssh -p $port" -a contrib/ $remote:~/lib/stock/contrib;
[ $RETVAL -eq 0 ] && RETVAL=$?
ssh -p $port $remote sudo systemctl start stock;
[ $RETVAL -eq 0 ] && RETVAL=$?
[ $RETVAL -ne 0 ] && exit 1
echo "deployed"
