## jump-folder

jump-folder is a folder bookmarking tool written in Go. It allows you to quickly change the current working directory using your bookmarked locations.

### Installation

```
go get github.com/asessa/jump-folder/...
```

### Usage

Configuration file is stored on ~/.jump-folder

To add a bookmark, simply go to the directory you want to bookmark and then save it.

```
$ jump-folder -a projects
```

To list your bookmarks:

```
$ jump-folder -l
```

Show the path of a bookmark:

```
$ jump-folder -p projects
```

Combine with other commands:

```
$ ls `jump-folder -p projects`
```

Delete a bookmark:

```
$ jump-folder -d projects
```

### Bash/Zsh Integration

Add the code below to your .zshrc / .bash_profile file:

```
eval "$(jump-folder -bash)"
```

To jump to a previously saved bookmark:

```
jump [bookmark]
```
