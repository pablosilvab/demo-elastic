# Elastic-Lib
Este proyecto tiene como objetivo proveer una función para escribir logs en Elasticsearch.

## Requisitos

* Tener instalado Go.
* Tener un server de Elasticsearch en ejecución.

<details>
  <summary>Elasticsearch Local</summary>
```
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.6.2
```
</details>


## Ejemplo 🚀

Se ha incluído un ejemplo que únicamente genera un índice y escribe un registro en Elasticsearch. 

```
make example
```
