# swarm-logspout-logstash

Derived work from looplab/logspout-logstash

### Environment variables

Add Logstash tags:

```bash
LOGSTASH_TAGS="zone-1,production"
```

Add Logstash fields:

```bash
LOGSTASH_FIELDS="myfield=something,anotherfield=something_else"
```

Add docker container labels as fields:

```bash
DOCKER_LABELS="on"
```
