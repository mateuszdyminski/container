## Container

This repo contains slides and code of presentation "Let's Build a Container in Go"

### Code examples 

Running examples in directories `v1` and `v2` are straightforward, to run examples from `v3`  you need to download simple linux filesystem (for example from: [https://wiki.ubuntu.com/Base](https://wiki.ubuntu.com/Base)) and put it in `/home/rootfs`

### Slides

To run the slides you just need to install go present tool 

```
go get golang.org/x/tools/cmd/present
```

And run the presentation:

```
present
```

### Inspiration

Great article about building primitive Docker in Go:
[https://www.infoq.com/articles/build-a-container-golang](https://www.infoq.com/articles/build-a-container-golang)


Nice presentation with live-coding:
[https://www.youtube.com/watch?v=HPuvDm8IC-4](https://www.youtube.com/watch?v=HPuvDm8IC-4)
