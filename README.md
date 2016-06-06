macarena
========

A relay bot for IRC channels spanning many networks.

Usage
-----

### From Binary

```console
$ tar xf downloaded-tarball
$ cd macarena-$VERSION-linux-amd64
$ cp example.conf.json config.json
$ ./run.sh
```

### From Source

```console
$ git clone github.com/Xe/macarena
$ gb build all
$ cp src/config/example.conf.json somewhere.json
$ ./bin/macarena -conf somewhere.json
```

#### Without [`gb`](https://getgb.io)

```console
$ ./build.sh
```

Please be sure to edit your config. See
[this](http://godoc.org/github.com/Xe/macarena/src/config) for help.

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
                        "host": "172.17.0.10",
                        "port": 6667,
                        "ssl":  false,
                        "pass": "foobang",
                        "bindhost": "172.17.0.1"
                }
        ],
        "myinfo": {
                "nick": "Macarena",
                "user": "relay",
                "real": "IRC Relay bot",
                "notify_connections": true
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

Please use [`gb`](http://getgb.io/) for building macarena. Please ensure all
code passes a build before pull requesting.

Support
-------

For help, please connect to `irc.ponychat.net` and join `#macarena`, or open
a github issue on [the main repository](https://github.com/Xe/macarena).
