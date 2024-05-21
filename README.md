## Elasticsearch Snapshot Restore Guide - Es快照恢复指南

This guide provides step-by-step instructions for restoring Elasticsearch snapshots from a local backup directory using Docker.

## Prerequisites

- Docker and Docker Compose installed on your system.
- A local backup directory containing your Elasticsearch snapshots (e.g., `/Users/Documents/es_backup/backup`).

## Steps to Restore Elasticsearch Snapshots

### 1. Start a New Elasticsearch Container

Create a `docker-compose.yml` file and start a new Elasticsearch container without mounting any volumes:

```yaml
services:
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:8.14.0
    container_name: elasticsearch
    environment:
      - discovery.type=single-node
      - xpack.security.enabled=false
      - xpack.security.http.ssl.enabled=false
    #      - bootstrap.memory_lock=true
    #      - ES_JAVA_OPTS=-Xms512m -Xmx512m
    #    ulimits:
    #      memlock:
    #        soft: -1
    #        hard: -1
    #    volumes:
    #      - /root/es_data:/usr/share/elasticsearch/data
    ports:
      - "9200:9200"
      - "9300:9300"
```
### 2. Copy Snapshot Data to the Container
Use the docker cp command to copy the snapshot data from your local backup directory to the container:
```shell
docker cp /Users/Documents/es_backup/backup/ elasticsearch:/usr/share/elasticsearch/data/backup
```
### 3. Modify Elasticsearch Configuration
Enter the container and modify the elasticsearch.yml file to add the path.repo setting:
```shell
docker exec -it --user root elasticsearch /bin/bash
cat /usr/share/elasticsearch/config/elasticsearch.yml
```
Add the following line:
```yaml
path.repo: ["/usr/share/elasticsearch/data/backup"]
```
### 4. Restart Elasticsearch Container
Exit the container and restart it to apply the configuration changes:
```shell
exit
docker restart elasticsearch
```
### 5. Configure Snapshot Repository
Configure the snapshot repository in Elasticsearch:
```shell
curl -X PUT "http://localhost:9200/_snapshot/my_backup" -H "Content-Type: application/json" -d'
{
  "type": "fs",
  "settings": {
    "location": "/usr/share/elasticsearch/data/backup"
  }
}
'
```
### 6. List Snapshots in the Repository
Verify that the snapshots are available in the repository:
```shell
curl -X GET "http://localhost:9200/_snapshot/my_backup/_all"
```
### 7. Restore Snapshot
Use the correct snapshot name to restore the data. For example, if the snapshot name is snapshot_1:
```shell
curl -X POST "http://localhost:9200/_snapshot/my_backup/snapshot_1/_restore"
```
### 8. Verify Restored Data
Verify that the data has been restored correctly:
```shell
curl -X GET "http://localhost:9200/_cat/indices?v"
```
### Troubleshooting
If you encounter any issues during the restore process, ensure that:

The snapshot files are correctly copied to the container.
The path.repo setting in elasticsearch.yml matches the location of the snapshot files.
The Elasticsearch container has the necessary permissions to read the snapshot files.

## Swagger Doc
```shell
swag init -g /cmd/main.go -exclude model,pkg
```

