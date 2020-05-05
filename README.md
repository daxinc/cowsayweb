# Cowsay

A simple function prints out quote as if it is said by a cow.

# Compile

- Mac: `make osx`
- Linux: `make linux`
- Windows: `make windows`

Compile files are located in `dist` folder.

# URLs

- Main: http://localhost:8080/
- Heath check: http://localhost:8080/health
- Terminate application: http://localhost:8080/kill

E.g.:

```
 ________________________________
< Rome was not built in one day. >
 --------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

# Build Image

## Run

- `CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix cgo -o cowsayweb .`
- `docker build -t cowsay .`
- `docker run -d -p 8080:8080 cowsay`
- Open http://localhost:8080

## Stop

- `docker ps`
- `docker stop CONTAINTER_ID`
