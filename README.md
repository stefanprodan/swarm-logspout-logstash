# swarm-logspout-logstash

Logspout adapter to send Docker Swarm logs to Logstash

### Usage

Stack:

```yaml
version: "3.3"

  logspout:
    image: stefanprodan/swarm-logspout-logstash
    environment:
      - "ROUTE_URIS=logstash+tcp://localhost:50005"
      - "LOGSPOUT=ignore"
      - "HTTP_PORT=55444"
      - "DOCKER_LABELS=on"
    volumes:
      - /etc/hostname:/etc/host_hostname:ro
      - /var/run/docker.sock:/var/run/docker.sock
    networks:
      - host-net
    deploy:
      mode: global
      restart_policy:
        condition: on-failure
        delay: 10s
        window: 5s
      resources:
        limits:
          cpus: '0.10'
          memory: 256M
        reservations:
          memory: 64M

networks:
  host-net:
    external:
      name: "host"
```

Logstash pipeline:

```
input {
  udp {
    port  => 50005
    codec => json
  }
  tcp {
    port  => 50005
    codec => json
  }
}

output {
	elasticsearch {
		hosts => ["es1:9200", "es2:9200", "es3:9200"]
	}
}
```
