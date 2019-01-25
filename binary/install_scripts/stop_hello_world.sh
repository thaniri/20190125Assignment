#!/bin/bash
appPID=`pgrep hello_world_app`
if [[ -n  $appPID ]]; then
   pkill hello_world_app
fi
