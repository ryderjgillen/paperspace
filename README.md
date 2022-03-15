## Design Notes

- The `NETLINK_SOCK_DIAG` api is used to retreive the bulk of the data. By parsing the data in `/proc` I was able to map the inode to the command associated with the process that opened to port.

- Prometheus metrics are exposed under the `/metrics` endpoint.


## How to Run

The project has been configured to run inside of a Docker container.

### Service

- `docker run -p 59001:59001 -p 59002:59002 port-service` -- runs the service using the default configuration

You can override the defaults by passings the following flags

- `-data-interval 5m` -- the duration the data is cached for

- `-address localhost` -- sets the address the GRPC service will listen for requests on

- `-port 59001` -- sets the port the GRPC service will listen for requests on

- `--prom-address localhost` -- sets the address that prometheus will listen for requests to the `/metrics` endpoint on

- `--prom-port 59002` -- sets the port that prometheus will listen for requests to the `/metrics` endpoint on

### Client

A simple client has been created to aid in testing. The client will poll the service for data at a specificed interval, printing the results to STDOUT.

- `docker run --network=host port-service-client poll` -- polls the service on an interval using the default port (59001).

- `docker run --network=host port-service-client poll --help` -- displays detailed help (there are a few flags that override the defaults, be sure to check them out)