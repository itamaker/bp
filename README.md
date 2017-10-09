# betaphor

Betaphor(bp) is a light weight command line tool trying to make your daily work more efficient. It works like `alias` except that no `~/.bashrc`,  `~/.zshrc` modification and `source ~/.zshrc` operation is ever required.

By now, it's for macOS ONLY.

## usage

#### 1. New bp shortcut

* add shortcut for specific webpage

```shell
$ bp add
$ Enter alias name: ggl
$ Enter command literal: open https://google.com
```

* add shorcuts for launch/close specific app

```shell
# launch ShadowsocksX from command line
$ bp add
$ Enter alias name: sss
$ Enter command literal: /Applications/ShadowsocksX.app
```

```shell
# quit ShadowsocksX from command line
$ bp add
$ Enter alias name: kss
$ Enter command literal: osascript -e 'quit app "ShadowsocksX"'
```

#### 2. Use bp shortcut

* open webpage

```shell
# open google.com (added before) in default explorer
$ bp ggl
```
* launch/quit ShadowsocksX

```shell
# launch ShadowsocksX
$ bp sss

# quit ShadowsocksX
$ bp kss
```

#### 3. list all shortcuts

```shell
$ bp ls
```

#### 4. Remove shortcuts

```shell
# remove one specific shortcuts
$ bp rm sss
# remove all shortcuts
$ bp reset
```

