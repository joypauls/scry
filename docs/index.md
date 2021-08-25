## How To Use

### Installation (via Binary)

M1 Mac:
```
wget https://github.com/joypauls/scry/releases/download/v0.0.1/scry-darwin-arm64.tar.gz
tar -xvzf scry-darwin-arm64.tar.gz
sudo mv scry-darwin-arm64 /usr/local/bin/scry
```

Confirm your installation with `scry -v`

### Basic Usage

### Compatibility

The only currently supported OS - that is, confirmed by testing... probably will still work on many others!) is Darwin. However, you will likely be able to build from source on your machine without issue if you are setup for Go development, see [Build From Source](#build-from-source).

Keep in mind this will also depend on which [terminal emulator](https://en.wikipedia.org/wiki/List_of_terminal_emulators) you are using. If you see problems running this on your OS/terminal, [let me know](#support-and-bugs).

### Customization

Use the `.scry.yaml` file

## Developers

This app is a just a regular old Go program so if you're set up to develop Go on your system you can easily compile it yourself.

### Build from Source

1. Clone the repository `git clone https://github.com/joypauls/scry.git`

2. Build the project `make build`
    - By default, it will compile to `<repo>/bin/scry`

3. Check that it worked with `<repo>/bin/scry -v`

## Support and Bugs

A bug? In my code?!? Unlikely!

Just kidding 😊 Thank you for trying **scry** out! Please head over to the project page on Github and [open an issue](https://github.com/joypauls/scry/issues/new) and let me know how I can help or what problems you're seeing!
