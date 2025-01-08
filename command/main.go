package main

import "github.com/spf13/cobra"

func newVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Short:   "show version",
		Example: "cmd-demo version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("cmd-demo version: ", shortVersion())
		},
	}

	return cmd
}

func newRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "run",
		Short:   "run cmd-demo",
		Example: "cmd-demo run",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("run cmd-demo", args[0])
		},
	}

	return cmd
}

func shortVersion() string {
	return "V1.0"
}

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cmd-demo [command]",
		Long:    "test cmd package usage",
		Version: shortVersion(),
	}

	cmd.AddCommand(newRunCommand())
	cmd.AddCommand(newVersionCommand())

	return cmd
}

func main() {
	cmd := newCommand()
	if err := cmd.Execute(); err != nil {
		panic(any(err))
	}
}
