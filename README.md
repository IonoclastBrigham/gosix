GOSIX source distribution
-------------------------

GOSIX is a vaguely standards-compliant source distribution of [POSIX][posix] CLI utility
clones. It is intended mostly for me to learn [go][golangorg] by actually writing
a series on small but non-trivial programs, but I'm publishing it publicly in
the hopes that others might find it educational or useful.

It may be used and redistributed under a permissive [MIT-style license](LICENSE).

[posix]: http://en.wikipedia.org/wiki/POSIX
[golangorg]: http://golang.org/

## Included Utilities

Currently implemented utilities:

* **gecho** - implements the `echo` command; respects the `-n` flag.
* **gcut** - implements `cut`; respects `-f` with comma lists (required) and `-d` (defaults to `'\t'`).
* **genv** - implements `env`; respects the `-i` flag to ignore the exported environment.

All commands will also support `-h` or `--help` to get a usage message. If you find one that
doesn't currently, please send me a pull request :)