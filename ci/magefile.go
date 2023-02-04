//go:build mage
// +build mage

package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"

	"github.com/magefile/mage/mg"
	"github.com/sagikazarmark/goci/lib/golang"
)

// Run tests
func Test(ctx context.Context) error {
	var clientOpts []dagger.ClientOpt

	if os.Getenv("DEBUG") == "true" {
		clientOpts = append(clientOpts, dagger.WithLogOutput(os.Stderr))
	}

	client, err := dagger.Connect(ctx, clientOpts...)
	if err != nil {
		return err
	}
	defer client.Close()

	goVersion := os.Getenv("GO_VERSION")
	if goVersion == "" {
		goVersion = "1.19.5"
	}

	c := golang.Test(
		client,

		golang.Version(goVersion),
		golang.CoverMode(golang.AtomicCoverMode),
		golang.CoverProfile("coverage.txt"),
	)

	err = process(ctx, c)
	if err != nil {
		return err
	}

	_, err = c.File("/src/coverage.txt").Export(ctx, "coverage.txt")
	if err != nil {
		return err
	}

	return nil
}

// Run linter
func Lint(ctx context.Context) error {
	var clientOpts []dagger.ClientOpt

	if os.Getenv("DEBUG") == "true" {
		clientOpts = append(clientOpts, dagger.WithLogOutput(os.Stderr))
	}

	client, err := dagger.Connect(ctx, clientOpts...)
	if err != nil {
		return err
	}
	defer client.Close()

	goVersion := os.Getenv("GO_VERSION")
	if goVersion == "" {
		goVersion = "1.19"
	}

	return process(ctx, golang.Lint(
		client,

		golang.Version(goVersion),
		golang.LinterVersion("v1.51.0"),
	))
}

// Run all checks
func Checks(ctx context.Context) error {
	var clientOpts []dagger.ClientOpt

	if os.Getenv("DEBUG") == "true" {
		clientOpts = append(clientOpts, dagger.WithLogOutput(os.Stderr))
	}

	client, err := dagger.Connect(ctx, clientOpts...)
	if err != nil {
		return err
	}
	defer client.Close()

	goVersions := []string{
		"1.18",
		"1.19",
	}

	var pipelines []*dagger.Container

	for _, goVersion := range goVersions {
		pipelines = append(pipelines, golang.Test(
			client,

			golang.Version(goVersion),
			golang.CoverMode(golang.AtomicCoverMode),
			golang.CoverProfile("coverage.txt"),
		))
	}

	for _, pipeline := range pipelines {
		err = process(ctx, pipeline)
		if err != nil {
			return err
		}
	}

	return nil
}

func process(ctx context.Context, container *dagger.Container) error {
	output, err := container.Stdout(ctx)

	fmt.Print(output)

	// if err != nil {
	// 	return err
	// }

	erroutput, err := container.Stderr(ctx)

	fmt.Print(erroutput)

	if err != nil {
		return err
	}

	exit, err := container.ExitCode(ctx)
	if err != nil {
		return err
	}

	if exit > 0 {
		return mg.Fatal(exit)
	}

	return nil
}
