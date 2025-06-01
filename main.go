// source https://github.com/ndanilo8/razer-mic-keepalive

package main

import (
	"fmt"
	"os/exec"
	"syscall"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

var ffmpegCmd *exec.Cmd

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Mic Keep-Alive")
	systray.SetTooltip("Keeping your mic alive with ffmpeg")

	mQuit := systray.AddMenuItem("Quit", "Stop ffmpeg and exit")

	// Start ffmpeg to keep mic alive
	micName := `Microphone (Razer BlackShark V2 Pro 2.4)`
	ffmpegCmd = exec.Command("ffmpeg",
		"-f", "dshow",
		"-i", "audio="+micName,
		"-t", "86400",
		"-f", "null", "-")

	// hide ffmpeg terminal window
	ffmpegCmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	err := ffmpegCmd.Start()
	if err != nil {
		fmt.Println("Failed to start ffmpeg:", err)
		systray.Quit()
	}

	go func() {
		<-mQuit.ClickedCh
		if ffmpegCmd != nil && ffmpegCmd.Process != nil {
			ffmpegCmd.Process.Kill()
		}
		systray.Quit()
	}()
}

func onExit() {
	// Cleanup
	if ffmpegCmd != nil && ffmpegCmd.Process != nil {
		ffmpegCmd.Process.Kill()
	}
}

func main() {
	systray.Run(onReady, onExit)
}
