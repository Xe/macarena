Deploying Macarena
==================

This is an example folder showing how to deploy macarena to a server. Its 
[Dockerfile](https://github.com/Xe/macarena/blob/master/Dockerfile) will 
automatically figure things out based on the `config.json` you stick in 
a folder with the Dockerfile in it.

Files
-----

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

The two files in this folder should allow you to make a suitable example 
deployment. **PLEASE CHANGE THE CHANNEL, NICK, USER, AND NETWORK SETTINGS**.

Running
-------

Simply run:

```console
$ docker build -t yourbot .
$ docker run -dit --name yourbot yourbot
```
