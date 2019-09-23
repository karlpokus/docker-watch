# docker-watch
docker cli plugin to watch and reload services per [this](https://github.com/docker/cli/issues/1534) design proposal. A work in progress.

# install
```bash
# Run build and copy binary to the docker cli-plugins dir
$ ./build.sh $OS && ./copy.sh $OS
```

# test
```bash
$ go test
```

# usage
```bash
# start watching dirs and have container restart on file change
# runs in foreground, stop by sending SIGINT
$ docker watch start $DIR:$CONTAINER [...more pairs]
# print version and exit
$ docker watch version
```

# todos
- [ ] ignore-syntax
- [x] parse cli args
- [ ] watcher
- [x] add version arg to metadata func

# license
MIT
