#!/bin/bash

gnome-terminal -- /bin/bash -c 'go run ../../cmd/server2server/main.go -MyPort 8080 -PartnerPort 8081 -ServerName Taro; exec bash'
gnome-terminal -- /bin/bash -c 'go run ../../cmd/server2server/main.go -MyPort 8081 -PartnerPort 8080 -ServerName Jiro; exec bash'
