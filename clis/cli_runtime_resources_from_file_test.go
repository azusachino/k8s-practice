package clis

import (
	"os"

	"testing"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes/scheme"
)

func TestFromFile(t *testing.T) {
	runFromFile()
}

// go run main.go resources.yaml
func runFromFile() {
	configFlags := genericclioptions.NewConfigFlags(true)

	cmd := &cobra.Command{
		Use:  "kubectl (almost)",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			builder := resource.NewBuilder(configFlags)
			namespace := ""
			if configFlags.Namespace != nil {
				namespace = *configFlags.Namespace
			}
			enforceNamespace := namespace != ""
			printer := printers.NewTypeSetter(scheme.Scheme).ToPrinter(&printers.YAMLPrinter{})
			err := builder.
				WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).
				NamespaceParam(namespace).
				DefaultNamespace().
				FilenameParam(enforceNamespace, &resource.FilenameOptions{Filenames: args}).
				Do().
				Visit(func(i *resource.Info, err error) error {
					if err != nil {
						return err
					}
					return printer.PrintObj(i.Object, os.Stdout)
				})
			if err != nil {
				panic(err.Error())
			}

		},
	}
	configFlags.AddFlags(cmd.PersistentFlags())
	if err := cmd.Execute(); err != nil {
		panic(err)
	}

}
