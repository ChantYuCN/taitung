# Develop

```
make rest-api
# coding
make nerdctl-build
make docker-build
kubectl  run   gserver --port=50021  --image chant/rest/coolguy:test-0.1  /home/restserver
kubectl expose po gserver --port 50029 --target-port 50029
```

# Test

```
export HLWORKSPACE=/home/labuser/habanashared/
./bin/grestserver

curl -X POST  127.0.0.1:9911/upload --form file='@filelocation'

```