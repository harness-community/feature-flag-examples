# connectivity check

go script to test connectivity to harness, or your proxy

## setup

the following environment variables are needed:

- FF_SDK_KEY: some FF sdk key from your harness account

### additional configuration

- FF_IDENTIFIER: flag id to test with, defaults to `test`
- RELAY_PROXY_ADDRESS: relay proxy address to connect to, defaults to harness saas

## running

you can run the script via go or docker, or k8s

### go

```
FF_SDK_KEY=XXXX go run main.go
```

### docker

```
docker run -e FF_SDK_KEY=XXXX --rm -it harnesscommunity/feature-flag-connectivity-check:latest
```

### k8s

fill in the deployment yaml and then:
```
kubectl apply -f deployment.yaml
```
