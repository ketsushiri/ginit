# ginit
[This utility](https://github.com/ketsushiri/contest-init) rewritten in Go for practice purpose. Initialize some files by template. Will work also on windows (maybe). Default ext for files -- `.cpp`.

## Build 
```bash
$ git clone https://github.com/ketsushiri/ginit.git
$ cd ginit
$ go build -o ginit main.go
```

## Usage
Simply launch with `--help` flag or:
```bash
$ ginit <dir-name> [A..Z]
...
$ ginit test A B C
```
Output will be directory with name 'test' and three files with the same content in it.
```bash
$ ginit -d ../other/template.hs -e hs test A B C
```
Output will be directory with name 'test' and three files (A.hs, B.hs, C.hs) in it.
