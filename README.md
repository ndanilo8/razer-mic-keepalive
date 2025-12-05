
# Razer Mic Keep-Alive Tray App

This tiny Go-based tray app prevents the **Razer BlackShark V2 Pro** wireless headset from disconnecting due to microphone inactivity.

*This is a workaround as Razer hasn't been able to fix their Sh\*t* 

It does this by launching `ffmpeg` in the background to continuously open the mic input — keeping it "alive" and preventing idle timeouts or power-saving disconnections.

Another workaround is to Disable `Allow the computer to turn off this device to save power` in device manager. (it didn't work for me) [Link here](https://insider.razer.com/razer-support-45/razer-blackshark-v2-pro-random-disconnect-issue-fixed-63698)

### tl;dr 
1. Install ffmpeg
2. [Download](https://github.com/ndanilo8/razer-mic-keepalive/releases/download/v0.1/RazerMicKeepAlive.exe) the executable and run it.
3. (Optional Autostart) Shortcut `RazerMicKeepAlive.exe`  to `%appdata%\Microsoft\Windows\Start Menu\Programs\Startup` folder


## Features

- Tray icon with clean **Quit** button
- Launches `ffmpeg` to keep mic active silently
- Fully self-contained Go app (no GUI window)
- Lighter than running OBS, Discord, etc..

## Getting Started (compile your own binary)

### 1. Install Requirements

- [Go](https://golang.org/dl/) (1.20+)
- [ffmpeg](https://ffmpeg.org/download.html) (add to PATH)


### 2. Find Your Microphone Name

Run in terminal:

```bash
ffmpeg -list_devices true -f dshow -i dummy
```

Look for your mic, e.g.:

```
"Microphone (Razer BlackShark V2 Pro 2.4)"
```

Copy that exact string.

### 3. Set Up the Project

```bash
git clone https://github.com/ndanilo8/razer-mic-keepalive.git
cd razer-mic-keepalive
go mod init razer-mic-keepalive
go get github.com/getlantern/systray
```

Then edit `main.go` with the string above:

```go
micName := `Your Device Name Here`
```


### 4. Build the App

```bash
go mod tidy
go build -ldflags="-H=windowsgui" -o RazerMicKeepAlive.exe main.go
```

> `-H=windowsgui` ensures no terminal window pops up when you launch it.


## Run

Just double-click `RazerMicKeepAlive.exe`. It will:
- Start `ffmpeg` in background
- Show a system tray icon
- Keep your mic open and stable
- Prevent Random Razor Disconnects

To stop it: Right-click tray icon → Quit


## Auto-Launch at Startup (Optional)

- Press `Win + R`, type `shell:startup`
- Place a shortcut to `RazerMicKeepAlive.exe` there

## License

MIT — do whatever you want.

## Credits

- `systray` by [Lantern](https://github.com/getlantern/systray)
- `ffmpeg` for the device-level capture
