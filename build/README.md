macarena
========

A relay bot for IRC channels spanning many networks.

Usage
-----

```console
$ cp example.conf.json config.json
$ ./run.sh
```

Please be sure to edit your config. See
[this](http://godoc.org/github.com/Xe/macarena/config) for help.

### Running

For a more complicated setup:

```console
$ ./macarena -h
Usage of ./macarena:
  -conf="./config.json": config file to use
$ ./macarena -conf /path/to/my/config
```

Macarena does not detach from the active console. As such it is suggested to 
run macarena inside [dtach](https://github.com/bogner/dtach), screen, tmux, as 
a service with upstart/systemd, or as a container.

For more information about the configuration file, see 
[here](https://godoc.org/github.com/Xe/macarena/config) or the included 
config.md file.

Support
-------

For help, please connect to `irc.ponychat.net` and join `#macarena`, or open 
a github issue on [the main repository](https://github.com/Xe/macarena).
