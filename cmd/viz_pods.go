package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type vizPodsCmd struct {
	out io.Writer
}

// NewVizPodsCommand creates the command for rendering the Kubernetes pods.
func NewVizPodsCommand(streams genericclioptions.IOStreams) *cobra.Command {
	helloWorldCmd := &vizPodsCmd{
		out: streams.Out,
	}

	cmd := &cobra.Command{
		Use:          "viz_pods",
		Short:        "Prints Kubernetes server version",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) != 0 {
				return errors.New("this command does not accept arguments")
			}
			return helloWorldCmd.run()
		},
	}

	cmd.AddCommand(newVersionCmd(streams.Out))
	return cmd
}

func (sv *vizPodsCmd) run() error {
	// Get the relevant data
	serverVersion, err := getPods()
	if err != nil {
		return err
	}

	// Print out (visuals)
	_, err = fmt.Fprintf(sv.out, "Hello from Kubernetes server with version %s!\n", serverVersion)
	if err != nil {
		return err
	}
	return nil
}

func getPods() (string, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return "", err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", err
	}

	// Get pods via
	sv, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return "", err
	}

	return sv.String(), nil
}
