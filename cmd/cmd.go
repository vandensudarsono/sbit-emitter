package cmd

import (
	"fmt"
	"sbit-emitter/config"
	"sbit-emitter/infrastructure/balance"
	"sbit-emitter/infrastructure/emitter"
	logger "sbit-emitter/infrastructure/log"
	"sbit-emitter/infrastructure/server/grpc"

	"github.com/spf13/cobra"
)

type Command struct {
	rootCmd *cobra.Command
}

var text = "sbit-emitter service v1.0"

// NewCommandEngine the command line boot loader
func NewCommand() *Command {
	var rootCmd = &cobra.Command{
		Use:   "sbit-emitter",
		Short: "sbit-emitter service command line",
		Long:  "sbit-emitter service command line",
	}

	return &Command{
		rootCmd: rootCmd,
	}
}

// Run the all command line
func (c *Command) Run() {
	//container := containers.NewEngine()
	var rootCommands = []*cobra.Command{
		{
			Use:   "grpc",
			Short: "Run sbit-emitter GRPC service",
			Long:  "Run sbit-emitter GRPC service",
			PreRun: func(cmd *cobra.Command, args []string) {
				// Show display text
				fmt.Println(text)

				// Load config
				config.LoadConfig()

				logger.WithFields(logger.Fields{"component": "command", "action": "grpc"}).
					Infof("starting GRPC server...")
			},
			Run: func(cmd *cobra.Command, args []string) {
				grpc.RunServer()                                                              // Run GRPC server
				fmt.Printf("error running b processor :%v\n", balance.RunBallanceProcessor()) // run balance processor
			},
			PostRun: func(cmd *cobra.Command, args []string) {
				// close database connection
				emitter.CloseEmitter()
				logger.WithFields(logger.Fields{"component": "command", "action": "grpc"}).
					Infof("stopping GRPC server done")
			},
		},
	}

	for _, command := range rootCommands {
		c.rootCmd.AddCommand(command)
	}

	c.rootCmd.Execute()
}

// GetRoot the command line service
func (c *Command) GetRoot() *cobra.Command {
	return c.rootCmd
}
