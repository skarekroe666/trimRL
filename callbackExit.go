package main

import "os"

func CallbackExit() error {
	os.Exit(0)

	return nil
}
