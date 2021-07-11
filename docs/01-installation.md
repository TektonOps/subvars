# Installation

### MacOS & Linux Homebrew

```bash
brew install kha7iq/tap/subvars
```

### Linux

* AUR
```bash
yay -S subvars
```

* Binary
```bash
export SUBVARS_VERSION="0.1.3"
wget -q https://github.com/kha7iq/subvars/releases/download/v${SUBVARS_VERSION}/subvars_Darwin_x86_64.tar.gz && \
tar -xf subvars_Darwin_x86_64.tar.gz && \
chmod +x subvars && \
sudo mv subvars /usr/local/bin/subvars
```

### Go Get

```bash
go get -u github.com/kha7iq/subvars
```

### Windows

```powershell
scoop bucket add subvars https://github.com/kha7iq/scoop-bucket.git
scoop install subvars
```

Alternatively you can head over to [release pages](https://github.com/kha7iq/subvars/releases)
and download the binary for windows & all other supported platforms.

### Docker

Docker container is available you can pull the `latest` version or provide specific `tag`
Checkout [release](https://github.com/kha7iq/subvars/releases) page for available versions.

* Running Container

```bash
docker pull khaliq/subvars:latest

docker run -it --rm khaliq/subvars:latest --help
```
