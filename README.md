# Wave

Wave is a markup language written in Go that transpiles its syntax to HTML with inline CSS.

### Getting Wave

1. Clone Wave on your system:

```
$ git clone https://www.github.com/KILLinefficiency/Wave.git
```

2. Compile Wave from the source code:

```
$ cd Wave
$ make
```

3. If you are a Windows user, then compile Wave as follows:

```
$ git clone https://www.github.com/KILLinefficiency/Wave.git
$ cd Wave
$ go build wave.go lib.go themes.go defaults.go variables.go contentLib.go htmlTemplates.go
```

### The Documentation

The documentation is available in the [``guide``](guide) directory in this repository.

The documentation for Wave is written in Wave itself. It is available as a [Wave script](guide/guide.txt) as well as a complete transpiled [HTML document](guide/guide.html).

See the online hosted documentation [here](https://killinefficiency.github.io/Wave/guide/guide.html).

### Examples

There are examples available in the [``examples``](examples) directory. Examples are available as both Wave scripts and complete HTML documents.

See the online hosted examples:

* [Example 1](https://killinefficiency.github.io/Wave/examples/example1/example1.html)
* [Example 2](https://killinefficiency.github.io/Wave/examples/example2/example2.html)
* [Example 3](https://killinefficiency.github.io/Wave/examples/example3/example3.html)
