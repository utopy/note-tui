package main

import "note-taking/core"

func main() {
	app := core.CreateApplication()

	pages := core.InitPages()

	app.LoadLayout(pages)

	app.Run()

}
