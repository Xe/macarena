Deploying Macarena
==================

This is an example folder showing how to deploy macarena to a server. 
Macarena's [Dockerfile](https://github.com/Xe/macarena/blob/master/Dockerfile) 
will automatically figure things out based on the `config.json` you stick in 
a folder with the Dockerfile in it.

The two files in this folder should allow you to make a suitable example 
deployment. **PLEASE CHANGE THE CHANNEL, NICK, USER, AND NETWORK SETTINGS**.

Running
-------

Simply run:

```console
$ cd path/you/downloaded/macarena/in/example
### EDIT THE CONFIG ###
$ docker build -t yourbot .
$ docker run -dit --name yourbot yourbot
```
