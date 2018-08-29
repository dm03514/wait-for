# wait-for
A Level 7 (Application) aware way to detect when services are fully initialized.

## [Overview](https://medium.com/dm03514-tech-blog/ci-testing-remove-docker-initialization-race-conditions-96caa159bd86)

## Usage

```
$ ./wait-for
NAME:
   wait-for - wait for a service to become available

USAGE:
   wait-for [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     http
     mysql
     postgres
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --timeout value, -t value          duration to wait until marking as failure and returning (default: 5m0s)
   --poll-interval value, --pi value  interval (default: 100ms)
   --help, -h                         show help
   --version, -v                      print the version
```

### HTTP

```
$ ./wait-for http --help
NAME:
   wait-for http -

USAGE:
   wait-for http [command options] [arguments...]

OPTIONS:
   --method value, -m value  http request method to use for polling (default: "GET")
   --url value               http uri to poll status of
   --body value              optional body to send
   
```

#### Success

    ```
    $ ./wait-for http --url=http://google.com/
    {"level":"info","msg":"polling","time":"2018-08-26T00:50:39Z"}
    {"level":"debug","module":"poller.HTTP","msg":"http_response","status_code":200,"time":"2018-08-26T00:50:40Z"}
    
    $ echo $?
    0
    ```

#### Timeout Reached

    ```
    $ ./wait-for -t 500ms http --url=http://www.google.com/1/2
    {"level":"info","msg":"polling","time":"2018-08-28T12:43:30Z"}
    {"level":"debug","module":"poller.HTTP","msg":"http_response","status_code":404,"time":"2018-08-28T12:43:30Z"}
    {"level":"info","msg":"polling","time":"2018-08-28T12:43:30Z"}
    {"level":"debug","module":"poller.HTTP","msg":"http_response","status_code":404,"time":"2018-08-28T12:43:30Z"}
    {"level":"info","msg":"polling","time":"2018-08-28T12:43:30Z"}
    {"level":"debug","module":"poller.HTTP","msg":"http_response","status_code":404,"time":"2018-08-28T12:43:31Z"}
    {"level":"info","msg":"polling","time":"2018-08-28T12:43:31Z"}
    {"level":"debug","module":"poller.HTTP","msg":"http_response","status_code":404,"time":"2018-08-28T12:43:31Z"}
    {"level":"info","msg":"polling","time":"2018-08-28T12:43:31Z"}
    {"level":"debug","module":"poller.HTTP","msg":"http_response","status_code":404,"time":"2018-08-28T12:43:31Z"}
    {"level":"info","msg":"polling","time":"2018-08-28T12:43:31Z"}
    {"level":"debug","module":"poller.HTTP","msg":"http_response","status_code":404,"time":"2018-08-28T12:43:31Z"}
    {"level":"info","msg":"timeout_reached","time":"2018-08-28T12:43:31Z"}
    {"level":"fatal","msg":"timeout reached: 500ms","time":"2018-08-28T12:43:31Z"}
    
    $ echo $?
    1
    ```
    
### Postgres

```
$ ./bin/wait-for postgres --help
NAME:
   wait-for postgres -

USAGE:
   wait-for postgres [command options] [arguments...]

OPTIONS:
   --connection-string value, --cs value  psql connection string [$WAIT_FOR_POSTGRES_CONNECTION_STRING]
```
 
#### Success

- Start polling while no postgres available 

```
$ export WAIT_FOR_POSTGRES_CONNECTION_STRING=postgresql://root:root@localhost/test?sslmode=disable
$ ./wait-for --poll-interval 1s postgres
{"level":"info","msg":"polling","time":"2018-08-28T13:06:47Z"}
{"err":"dial tcp 127.0.0.1:5432: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T13:06:47Z"}
```
 
- Start postgres

```
$ docker-compose down && docker-compose up
```
    
- Output from `wait-for` as compose comes up

```
 {"level":"info","msg":"polling","time":"2018-08-28T13:06:48Z"}
 {"err":"dial tcp 127.0.0.1:5432: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T13:06:48Z"}
 {"level":"info","msg":"polling","time":"2018-08-28T13:06:49Z"}
 {"err":"dial tcp 127.0.0.1:5432: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T13:06:49Z"}
 {"level":"info","msg":"polling","time":"2018-08-28T13:06:50Z"}
 {"err":"read tcp 127.0.0.1:54470-\u003e127.0.0.1:5432: read: connection reset by peer","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T13:06:51Z"}
 {"level":"info","msg":"polling","time":"2018-08-28T13:06:51Z"}
 {"err":"EOF","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T13:06:51Z"}
 {"level":"info","msg":"polling","time":"2018-08-28T13:06:52Z"}
 {"err":"EOF","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T13:06:52Z"}
 {"level":"info","msg":"polling","time":"2018-08-28T13:06:53Z"}
 {"err":"EOF","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T13:06:53Z"}
 {"level":"info","msg":"polling","time":"2018-08-28T13:06:54Z"}
 {"err":null,"level":"debug","msg":"poll_result","ready":true,"time":"2018-08-28T13:06:54Z"}
 
 $ echo $?
   0
```

#### Failure

```
$ ./bin/wait-for --poll-interval 500ms -t 1s postgres
{"level":"info","msg":"polling","time":"2018-08-28T13:10:53Z"}
{"err":"dial tcp 127.0.0.1:5432: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T13:10:53Z"}
{"level":"info","msg":"polling","time":"2018-08-28T13:10:54Z"}
{"err":"dial tcp 127.0.0.1:5432: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T13:10:54Z"}
{"level":"info","msg":"polling","time":"2018-08-28T13:10:54Z"}
{"err":"dial tcp 127.0.0.1:5432: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T13:10:54Z"}
{"level":"info","msg":"timeout_reached","time":"2018-08-28T13:10:54Z"}
{"level":"fatal","msg":"timeout reached: 1s","time":"2018-08-28T13:10:54Z"}

$ echo $?
1
```

### MySQL

```
$ ./bin/wait-for mysql --help
  NAME:
     wait-for mysql -
  
  USAGE:
     wait-for mysql [command options] [arguments...]
  
  OPTIONS:
     --connection-string value, --cs value  mysql connection string [$WAIT_FOR_MYSQL_CONNECTION_STRING] 
```

#### Success
- Start waiting
```
$ export WAIT_FOR_MYSQL_CONNECTION_STRING=root:root@/performance_schema
$ ./bin/wait-for --poll-interval=1s  mysql
{"level":"info","msg":"polling","time":"2018-08-28T20:17:32Z"}
{"err":"dial tcp 127.0.0.1:3306: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T20:17:32Z"}
{"level":"info","msg":"polling","time":"2018-08-28T20:17:33Z"}
{"err":"dial tcp 127.0.0.1:3306: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T20:17:33Z"}
{"level":"info","msg":"polling","time":"2018-08-28T20:17:34Z"}
{"err":"dial tcp 127.0.0.1:3306: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T20:17:34Z"}
{"level":"info","msg":"polling","time":"2018-08-28T20:17:35Z"}
[mysql] 2018/08/28 20:17:35 packets.go:36: unexpected EOF
[mysql] 2018/08/28 20:17:35 packets.go:36: unexpected EOF
[mysql] 2018/08/28 20:17:35 packets.go:36: unexpected EOF
{"err":"driver: bad connection","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T20:17:35Z"}
....
```

- Start mysql 

```
vagrant@ubuntu-xenial:/vagrant_data/go/src/github.com/dm03514/wait-for/tests/mysql⟫ docker-compose down && docker-compose up
```

- When mysql is initialized logs of wait-for

```
{"err":"driver: bad connection","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T20:17:44Z"}
{"level":"info","msg":"polling","time":"2018-08-28T20:17:45Z"}
[mysql] 2018/08/28 20:17:45 packets.go:36: unexpected EOF
[mysql] 2018/08/28 20:17:45 packets.go:36: unexpected EOF
[mysql] 2018/08/28 20:17:45 packets.go:36: unexpected EOF
{"err":"driver: bad connection","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T20:17:45Z"}
{"level":"info","msg":"polling","time":"2018-08-28T20:17:46Z"}
{"err":null,"level":"debug","msg":"poll_result","ready":true,"time":"2018-08-28T20:17:46Z"}
```

#### Failure

- No MySQL up timeout reached

```
$ export WAIT_FOR_MYSQL_CONNECTION_STRING=root:root@/performance_schema

$ ./bin/wait-for --timeout=1s --poll-interval=250ms mysql
  {"level":"info","msg":"polling","time":"2018-08-28T20:16:12Z"}
  {"err":"dial tcp 127.0.0.1:3306: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T20:16:12Z"}
  {"level":"info","msg":"polling","time":"2018-08-28T20:16:12Z"}
  {"err":"dial tcp 127.0.0.1:3306: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T20:16:12Z"}
  {"level":"info","msg":"polling","time":"2018-08-28T20:16:13Z"}
  {"err":"dial tcp 127.0.0.1:3306: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T20:16:13Z"}
  {"level":"info","msg":"polling","time":"2018-08-28T20:16:13Z"}
  {"err":"dial tcp 127.0.0.1:3306: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T20:16:13Z"}
  {"level":"info","msg":"polling","time":"2018-08-28T20:16:13Z"}
  {"err":"dial tcp 127.0.0.1:3306: connect: connection refused","level":"debug","msg":"poll_result","ready":false,"time":"2018-08-28T20:16:13Z"}
  {"level":"info","msg":"timeout_reached","time":"2018-08-28T20:16:13Z"}
  {"level":"fatal","msg":"timeout reached: 1s","time":"2018-08-28T20:16:13Z"}
  
$ echo $?
  1
```
