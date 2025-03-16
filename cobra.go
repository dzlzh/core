package core

import "github.com/spf13/cobra"

var G_Cobra *cobra.Command

type CobraOption func(*cobra.Command)

func init() {
	G_Cobra = &cobra.Command{}
}

func WithCobra(cmd *cobra.Command, opts ...CobraOption) {
	for _, opt := range opts {
		opt(cmd)
	}
}

func WithCobraRunE(cmd *cobra.Command, runE func(cmd *cobra.Command, args []string) error) {
	cmd.RunE = runE
}

func WithCobraRun(cmd *cobra.Command, run func(cmd *cobra.Command, args []string)) {
	cmd.Run = run
}

func WithCobraShort(cmd *cobra.Command, short string) {
	cmd.Short = short
}

func WithCobraLong(cmd *cobra.Command, long string) {
	cmd.Long = long
}

func WithCobraExample(cmd *cobra.Command, example string) {
	cmd.Example = example
}

func WithCobraAnnotations(cmd *cobra.Command, annotations map[string]string) {
	cmd.Annotations = annotations
}

func WithCobraValidArgs(cmd *cobra.Command, validArgs []string) {
	cmd.ValidArgs = validArgs
}

func WithCobraValidArgNames(cmd *cobra.Command, validArgNames []string) {
	cmd.ValidArgNames = validArgNames
}

func WithCobraArgs(cmd *cobra.Command, args []string) {
	cmd.Args = args
}

func WithCobraArgAliases(cmd *cobra.Command, argAliases []string) {
	cmd.ArgAliases = argAliases
}
