package dialog

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"github.com/dbatbold/beep"
	"github.com/johnswanson/ttc"
	"log"
	"os/exec"
	"strings"
	"time"
)

func promptUser(prompt string) (e error, resp string) {
	music := beep.NewMusic("")
	if err := beep.OpenSoundDevice("default"); err != nil {
		log.Fatal(err)
	}
	if err := beep.InitSoundDevice(); err != nil {
		log.Fatal(err)
	}
	defer beep.CloseSoundDevice()

	musicScore := `
	SR9SS0SD0SA9VPSS9SR9SD9DSHLC5qicbmRS C5qicbmC5qicbmDQC6qicbmHLi
	`
	reader := bufio.NewReader(strings.NewReader(musicScore))
	go music.Play(reader, 100)
	music.Wait()
	beep.FlushSoundBuffer()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "/usr/bin/Xdialog", "--stdout", "--inputbox", prompt, "0", "0")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	done := make(chan error)
	go func() { done <- cmd.Wait() }()
	select {
	case err := <-done:
		if err != nil {
			return errors.New("no response"), ""
		}
		return nil, strings.TrimSpace(out.String())
	case <-time.After(time.Minute):
		return errors.New("no response"), ""
	}

}

func Request(ts int64) (e error, p ttc.Ping) {
	err, tags := promptUser("what are you doing RIGHT NOW?")
	if err != nil {
		return errors.New("No response from user"), p
	}
	p = ttc.Ping{
		Tags:      ttc.PingTags(tags),
		Timestamp: ttc.PingTime(ts * 1000),
	}
	return nil, p
}
