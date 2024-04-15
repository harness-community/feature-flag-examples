# connectivity check

python script to test connectivity to harness, or your proxy

## setup

the following environment variables are needed:

- FF_SDK_KEY: some FF sdk key from your harness account

### additional configuration

- FF_IDENTIFIER: flag id to test with, defaults to `test`
- RELAY_PROXY_ADDRESS: relay proxy address to connect to, defaults to harness saas

## running

you can run the script via python or docker

### python

```
pip install -r requirements.txt
FF_SDK_KEY=XXXX python main.py
```

### docker

```
docker run -e FF_SDK_KEY=XXXX --rm -it harnesscommunity/feature-flag-connection-test
```
