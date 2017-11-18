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


