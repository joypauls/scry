# Documentation

Welcome, happy to see you! While this project solidifies, feel free to head over to the project page and open an issue with any questions or requests. Happy hacking!

- [Get Started](#get-started)
- [Usage](#usage)
- [Additional Stuff](#additional-stuff)


<iframe src="https://giphy.com/embed/l0Iy29zHAcTFJ7jXO" width="480" height="80" frameBorder="0" class="giphy-embed" allowFullScreen></iframe><p><a href="https://giphy.com/gifs/internet-2d-looping-l0Iy29zHAcTFJ7jXO">via GIPHY</a></p>

<br/>
<br/>

# Get Started

## Installation via Binary

Right now, the binaries are not being built automatically. Our advice would be to go straight to the source for now if you're looking to experiment with the project. If you still want to use the pre-built releases, see which are available on the [releases page](https://github.com/joypauls/scry/releases).

M1 Mac:
```
wget https://github.com/joypauls/scry/releases/download/v0.0.1/scry-darwin-arm64.tar.gz
tar -xvzf scry-darwin-arm64.tar.gz
sudo mv scry-darwin-arm64 /usr/local/bin/scry
```

Confirm your installation with `scry -v`

## Build from Source

1. Clone the repository `git clone https://github.com/joypauls/scry.git`

2. Build the project `make build`
    - By default, it will compile to `<repo>/bin/scry`

3. Check that it worked with `<repo>/bin/scry -v`

# Usage

## Command Line Interface

Simply use `scry` to get started, or pass a path argument (`scry /some/place/cool`) like the classic `ls` program if you're feeling spicy.

Take a look at `scry --help` for more options.

## Basic

| Key | Description | Alternative |
| --- | --- | --- |
| Up Arrow | Scroll through directory | W |
| Down Arrow | Scroll through directory | S |
| Left Arrow | Scroll through directory | A |
| Right Arrow | Scroll through directory | D |
| Q | Exit | ESC |

<br/>
<br/>
---------
<br/>
<br/>

# Additional Stuff

## Compatibility

The only currently supported OS (that is, confirmed by user testing... probably will still work on many others!) is Darwin. However, you will likely be able to build from source on your machine without issue if you are setup for Go development, see [Build From Source](#build-from-source).

Keep in mind this could also depend on which [terminal emulator](https://en.wikipedia.org/wiki/List_of_terminal_emulators) you are using. If you see problems running this on your OS/terminal, [let me know](#support-and-bugs).

## Customization

Use the `.config/scry/config.yaml` file. Look at [the parser](https://github.com/joypauls/scry/blob/main/app/config.go) for the currently implemented features (we are actively adding to these).

## Developers

Check out the project on [GitHub](https://github.com/joypauls/scry)!

## Support and Bugs

A bug? In my code?!?

<iframe src="https://giphy.com/embed/3o7aTIGlhSo1bL8QUg" width="480" height="270" frameBorder="0" class="giphy-embed" allowFullScreen></iframe><p><a href="https://giphy.com/gifs/filmeditor-clueless-movie-3o7aTIGlhSo1bL8QUg">via GIPHY</a></p>

Just kidding 😊 Thank you for trying **scry** out! Feedback from actual real world testing is huge and we really appreciate it. Please head over to the project page on Github and [open an issue](https://github.com/joypauls/scry/issues/new) and let me know how I can help or what problems you're seeing!
