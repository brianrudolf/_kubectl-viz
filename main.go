package main

import (
	"github.com/brianrudolf/kubectl-viz/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
)

func main() {
	vizPodsCmd := cmd.NewVizPodsCommand(genericclioptions.IOStreams{In: os.Stdin,
		Out: os.Stdout, ErrOut: os.Stderr})

	if err := vizPodsCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
