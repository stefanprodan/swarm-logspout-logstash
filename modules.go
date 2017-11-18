package main

import (
	_ "github.com/stefanprodan/swarm-logspout-logstash/logstash"
	_ "github.com/gliderlabs/logspout/transports/udp"
	_ "github.com/gliderlabs/logspout/transports/tcp"
)
