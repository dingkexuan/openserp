#!/bin/bash

nohup ./openserp serve -a 127.0.0.1 -p 7000 >> nohup_logs.txt 2>&1 &
ps -ef | grep "openserp" | grep -v "grep" | awk '{print $2}'
