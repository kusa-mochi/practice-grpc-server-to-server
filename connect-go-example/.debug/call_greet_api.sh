#!/bin/bash

curl --header "Content-Type: application/json" --data '{"name": "Shibushi"}' http://localhost:8080/greet.v1.GreetService/Greet
