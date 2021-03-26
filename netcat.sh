#!/bin/bash
while true; do    echo -e "HTTP/1.1 200 OK\n\n $(date)" | nc -l -p 7070 -q 1; done