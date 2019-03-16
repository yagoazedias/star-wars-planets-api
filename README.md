# Star Wars Planets API
A simple REST API to consult the start wars planets 

## Getting started

### Install docker
To turn our lives easier we going to use `docker` and `docker-compose` to generate a mongodb database container. Click [here](https://docs.docker.com/install/) to install docker and run `make setup` to start our containers.

### Install Golang
We also have [a bash script](https://github.com/canha/golang-tools-install-script/blob/master/goinstall.sh) to install Golang

```bash
wget https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh # Download it
./goinstall.sh # Install it
```

### Request samples

- To create a planet

```bash
curl -X POST \                                                         
  http://localhost:4000/planet \
    -d '{
    "name": "Alderaan",
    "weather": "Cold",
    "terrain": "Low"
}'
```
- To update a planet by id

```bash
curl -X PUT \                                                          
  http://localhost:4000/planet/id/{id} \
    -d '{
    "name": "Alderaan",
    "weather": "New Value",
    "terrain": "Low"
}'
```
- To get all planets

```bash
curl -X GET http://localhost:4000/planet
```
- To get a planet by it id

```bash
curl -X GET http://localhost:4000/planet/id/{id}
```

- To get a planet by it name

```bash
curl -X GET http://localhost:4000/planet\?name\={name}
```

- To delete a planet by it id

```bash
curl -X DELETE http://localhost:4000/planet/id/{id}
```

## Troubleshooting

#### Tests failing in make tests
- The application test is end-to-end so it going to use your local database to persist data. Be sure that you have your container up to date with `make setup`