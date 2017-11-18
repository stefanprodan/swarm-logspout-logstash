### Docker Swarm ELK v6

Services:

* Elastic Search (3 nodes)
* Logstash (global)
* Logsprout (global)
* Kibana 

### Prerequisites 

Set vm map:

```bash
sudo sysctl -w vm.max_map_count=262144
sudo echo 'vm.max_map_count=262144' >> /etc/sysctl.conf
```

Disable swap in `/etc/fstab` and run:

```bash
sudo swapoff -a
```

Grand group access to local dir bind mount:

```bash
mkdir esdatadir
chmod g+rwx esdatadir
chgrp 1000 esdatadir
```

### Queries

Cluster health:

```bash
curl -XGET '35.198.137.226:9200/_cat/health?v&pretty'
```

List indices:

```bash
curl -XGET '35.198.137.226:9200/_cat/indices?v&pretty'
```

Indices settings:

```bash
curl -XGET '35.198.137.226:9200/_all/_settings?pretty'
```

Copy index:

```bash
curl -XPOST '35.198.178.32:9200/_reindex?pretty' -H 'Content-Type: application/json' -d' { "source": { "index": "logstash-2017.11.18" }, "dest": { "index": "logstash-2017.11.17" } } '
```

