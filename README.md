# Elastic-Lib
This project allows writing logs to Elasticsearch.

## Requirements

* Go
* Run Elasticsearch

For Elastic I use Docker
```
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.6.2
```

## Installation

Run 
```
go get github.com/pablosilvab/elastic
```

## Example 🚀

The example create an index and write a record in Elasticsearch.
```
make example
```
