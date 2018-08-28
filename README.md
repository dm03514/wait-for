# wait-for
Detect when services are available.

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
     help, h  Shows a list of commands or help for one command

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

- Success

    ```
    $ ./wait-for http --url=http://google.com/
    {"level":"info","msg":"polling","time":"2018-08-26T00:50:39Z"}
    {"level":"debug","module":"poller.HTTP","msg":"http_response","status_code":200,"time":"2018-08-26T00:50:40Z"}
    
    $ echo $?
    0
    ```

- Timeout Reached

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



