package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes/scheme"
)

func main() {
	// Already familiar stuff...
	configFlags := genericclioptions.NewConfigFlags(true)

	cmd := &cobra.Command{
		Use:  "kubectl (even closer to it this time)",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			// Our hero - The Resource Builder.
			builder := resource.NewBuilder(configFlags)

			namespace := ""
			if configFlags.Namespace != nil {
				namespace = *configFlags.Namespace
			}

			// Let the Builder do all the heavy-lifting.
			obj, _ := builder.
				// Scheme teaches the Builder how to instantiate resources by names.
				WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).
				// Where to look up.
				NamespaceParam(namespace).
				// What to look for.
				ResourceTypeOrNameArgs(true, args...).
				// Do look up, please.
				Do().
				// Convert the result to a runtime.Object
				Object()

			fmt.Println(obj)
		},
	}

	configFlags.AddFlags(cmd.PersistentFlags())

	_ = cmd.Execute()
}
