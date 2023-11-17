package main

func main() {
	app := CreateApplication()

	pages := InitPages()

	app.LoadLayout(pages)

	app.Run()

}
