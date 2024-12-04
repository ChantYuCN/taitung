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
export HLWORKSPACE=/mnt/
./bin/grestserver

curl -X POST  127.0.0.1:9911/upload --form file='@filelocation'


curl -X GET 127.0.0.1:9911/filepath  -H "Content-Type: application/json" -d '{"sn":"AO45001234"}'
```