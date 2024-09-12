# Container

This repo contains slides and code of presentation "Let's Build a Container in Go"

## Code examples

It's good to run all of the examples in Linux environment.

To start local ubuntu we can use lima:

```bash
limactl start --name=dev lima/dev.yaml
```

and `ssh`:

```bash
limactl shell dev
```

## Slides

To run the slides you just need to install go present tool

```bash
go get golang.org/x/tools/cmd/present
```

And run the presentation:

```bash
present
```

## Inspiration

Great article about building primitive Docker in Go:
[https://www.infoq.com/articles/build-a-container-golang](https://www.infoq.com/articles/build-a-container-golang)

Nice presentation with live-coding:
[https://www.youtube.com/watch?v=HPuvDm8IC-4](https://www.youtube.com/watch?v=HPuvDm8IC-4)

Containers From Scratch with Golang:
[https://medium.com/@ssttehrani/containers-from-scratch-with-golang-5276576f9909](https://medium.com/@ssttehrani/containers-from-scratch-with-golang-5276576f9909)