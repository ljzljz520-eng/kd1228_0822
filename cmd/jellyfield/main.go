package main

import (
	"flag"
	"fmt"
	"jellyfield/app"
	"jellyfield/model"
	"jellyfield/persistence"
	"os"
)

func main() {
	dbPath := flag.String("db", "jellyfield.db", "embedded scene database path")
	steps := flag.Int("steps", 3, "deterministic frames to advance")
	jsonOutput := flag.Bool("json", false, "print the render frame as JSON")
	flag.Parse()
	store, err := persistence.Open(*dbPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer store.Close()
	service, err := app.New(app.DemoScene(), store)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if _, err := store.LoadScene(service.Scene().ID); err != nil {
		if err := service.Save(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		if err := service.Restore(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if err := service.SetSpeedAt(model.Vector{X: 220, Y: 210}, 0.35); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for i := 0; i < *steps; i++ {
		service.Step()
	}
	if *jsonOutput {
		output, err := service.Render()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(output)
	} else {
		fmt.Println(service.Summary())
		fmt.Printf("metrics frame=%d particles=%d average_speed=%.2f brightness=%.2f\n", service.Metrics().Frame, service.Metrics().ParticleCount, service.Metrics().AverageSpeed, service.Metrics().Brightness)
	}
}
