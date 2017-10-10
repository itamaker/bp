# Betaphor - bp

Betaphor(bp) is a light weight command line tool trying to make your daily work more efficient. It works like `alias` except that no modifying `~/.bashrc` or  `~/.zshrc` and no `source ~/.?shrc` is ever required.

By now, it's for macOS ONLY.

## Installation

#### 1. manually (from source code)

* Install Golang (>=1.9.1), follow instructions on [golang.org](https://golang.org/dl/)

* Clone project and install

```shell
$ git clone https://github.com/keyOfVv/bp.git
$ mv ./bp/ ~/go/src/
$ cd ~/go/src/bp
$ ./install.sh
```

#### 2. homebrew - Coming soon.

## Update

```shell
$ cd ~/go/src/bp
$ ./update.sh
```

## Usage

#### 1. new bp shortcut

* add shortcut for specific webpage

```shell
$ bp add
> Enter alias name: ggl
> Enter command literal: open https://google.com
```

* add shorcuts for launch/close specific app

```shell
# launch ShadowsocksX from command line
$ bp add
> Enter alias name: sss
> Enter command literal: /Applications/ShadowsocksX.app
```

```shell
# quit ShadowsocksX from command line
$ bp add
> Enter alias name: kss
> Enter command literal: osascript -e 'quit app "ShadowsocksX"'
```

#### 2. use bp shortcut

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

#### 4. remove shortcut(s)

```shell
# remove one specific shortcuts
$ bp rm sss
# remove all shortcuts
$ bp rmAll
```
## Uninstallation

```shell
$ cd ~/go/src/bp
$ ./uninstall.sh
```

