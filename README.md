# Mafia formats
This repository contains known reverse-engineered binary format specifications with sample parsers and documentation.

**NOTE THAT THIS IS STILL WORK IN PROGRESS!**

## How to use
To compile the documentation, you need to make sure [go-specgen](https://github.com/zaklaus/go-specgen) is installed and within your `$PATH` variable.

### Usage: Documentation
Simply run `compile.sh` to generate the documentation.

### Usage: Parser utility
Some formats might have sample read/write tools written, so you can easily read and dump the binary data in various formats. Check out the `parsers` directory.

To run parsers written in Go language, make sure you have Go installed, then simply:
```sh
go run parsers/<format>/go/*.go --file=<path-to-file>
```

such as
```sh
go run parsers/textdb/go/*.go --file=<gamedir>/tables/textdb_cz.def
```

[![asciicast](https://asciinema.org/a/224466.svg)](https://asciinema.org/a/224466)


Reference utils output all data into standard output and make use of standard error for error reporting. You can make use of the *nix power to manipulate outputted JSON data easily.

## License
This repository is licensed under Apache 2.0 license. See `COPYING.MD` for more information.

