package main

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/gizak/termui"
)

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()
	
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	ls := termui.NewList()
	ls.ItemFgColor = termui.ColorYellow
	ls.BorderLabel = "Containers ID"
	ls.Height = 20
	ls.Width = 50
	ls.Y = 0

	termui.Handle("/timer/1s", func(e termui.Event) {
		conts, _ := client.ListContainers(docker.ListContainersOptions{All: false})

		var l []string
		for _, cont := range conts {
			l = append(l ,cont.ID)
		}
		ls.Items = l
		termui.Render(ls)
		termui.Handle("/sys/kbd/q", func(termui.Event) {
			termui.StopLoop()
		})
	})
	termui.Loop()
}

