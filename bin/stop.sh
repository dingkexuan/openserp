#!/bin/bash
pids=$(ps -ef | grep "openserp" | grep -v "grep" | awk '{print $2}')
for pid in ${pids}
do
    echo "kill openserp pid" $pid
    kill $pid
done
