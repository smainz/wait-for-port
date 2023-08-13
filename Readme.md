# wait-for-port

Woodpecker Plugin to wait for some networked service to become available within a given timeout.

## Build

Build the executable:

```bash
GOOS=linux
GOARCH=amd64

go build -a -v -o release/${GOOS}/${GOARCH}/wait-for-port
```

Build Docker image:

```bash
docker build -t wait-for-port:$(<version.txt) .
```

## Usage

Woodpecker pipeline:

```yaml
  steps:
    - name: Wait for Database service to come up
      image: smainz/wait-for-port
      settings:
        host: postgres
        port: 5432
        timeout: 20s      
```

Execute from the working directory:

```bash
docker run --rm \
  -e PLUGIN_HOST=www.dynaware.de \
  -e PLUGIN_PORT=80 \
  -e PLUGIN_TIMEOUT=120s \
  smainz/wait-for-port
```

Use local binary:

```bash
wait-for-port --host=localhost --port=80 --timeout=120s
```
