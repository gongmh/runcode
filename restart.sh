#!/bin/bash

go build .

kill `ps aux | grep './runcode' | grep -v grep | awk '{print $2}'`

nohup ./runcode >> nohup.out 2>&1 &
