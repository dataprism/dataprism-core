### Dataprism core

#### Getting Started
1. Start consul in developer mode
```
consul agent -dev
``` 

1. Start nomad in developer mode
```
nomad agent -dev
```

1. Create the configuration inside consul
Go to http://localhost:8500/ui and press on the Key/Value button. Create a new Key at /config (enter 'config', the / 
is already there) with the following payload:

```
{
    "cluster": {
        "servers": [ "localhost:9200" ],
        "buffer_max_ms": 10000,
        "buffer_max_msg": 1000
    },
    "jobs_dir": "/tmp"
}
```

1. Set the environment variables
```
export NOMAD_ADDR="http://localhost:4646"
export CONSUL_HTTP_ADDR="http://localhost:8500"
```

1. Launch the application
```
./dataprism-core
```