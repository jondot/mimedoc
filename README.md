# Mimedoc

A small utility that makes sure a file name is really representing its
content, by doing extension and content based mimetype identification.

## Quick Start


```bash
$ git clone https://github.com/jondot/mimedoc
$ make dev
$ ./mimedoc --help
```

`mimedoc` will recursively walk a path you've given it. You'll get this help screen with `--help`:

```
NAME:
   mimedoc - Cross reference extension-based and content-based mimetypes.

USAGE:
   mimedoc [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
  jondotan@gmail.com - <unknown@email>

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --path, -p "."				destination path
   --report, -r					generate a live report
   --ext, -e [--ext option --ext option]	Pick a specific extension
   --help, -h					show help
   --version, -v				print the version

```

Example runs:

```
➜  mimedoc  ./mimedoc --report --path=../connect-correlationid --ext .png
FILE	MIME	EXT
../connect-correlationid/sword.png	inode/x-empty	image/png
1/18 checked (5.56%), 1 mismatch, 0 unknown by mime, 0 unknown by ext
```

And a simple hand inspection will reveal that `sword.png` is really a hoax:

```
➜  mimedoc  ls -la ../connect-correlationid/sword.png
-rw-r--r--  1 dotan  staff  0 Dec 25 18:45 ../connect-correlationid/sword.png
```

# Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).


# Copyright

Copyright (c) 2014 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.



