macarena
========

A relay bot for IRC channels spanning many networks.

Usage
-----

```console
$ go get github.com/Xe/macarena
$ cp $GOPATH/src/github.com/Xe/macarena/config/example.conf.json somewhere.json
$ macarena -conf somewhere.json
```

Please be sure to edit your config. See
[this](http://godoc.org/github.com/Xe/macarena/config) for help.

Via Docker
----------

### Files

The dockerfile:

```Dockerfile
FROM xena/macarena
```

and the configuration:

```json
{
        "networks": [
                {
                        "name": "PonyChat",
                        "host": "irc.ponychat.net",
                        "port": 6697,
                        "ssl":  true,
                        "pass": "foobang"
                },
                {
                        "name": "ShadowNET",
                        "host": "irc.yolo-swag.com",
                        "port": 6667,
                        "ssl":  false,
                        "pass": "foobang"
                }
        ],
        "myinfo": {
                "nick": "Macarena",
                "user": "relay",
                "real": "IRC Relay bot"
        },
        "channels": ["#macarena"]
}
```

The two files in the `example/` folder in the root of this repository should
allow you to make a suitable example deployment. **PLEASE CHANGE THE CHANNEL,
NICK, USER, AND NETWORK SETTINGS**.

### Running

Simply run:

```console
$ docker build -t yourbot .
$ docker run -dit --name yourbot yourbot
```

Notes
-----

Macarena does not detach from the active console. As such it is suggested to 
run macarena inside [dtach](https://github.com/bogner/dtach), screen, tmux, as 
a service with upstart/systemd, or as a container.

Support
-------

For help, please connect to `irc.ponychat.net` and join `#macarena`, or open 
a github issue on [the main repository](https://github.com/Xe/macarena).
