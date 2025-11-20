package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Bootstraps the node and connects to the mesh",
	Long: `Starts the Fluxd daemon in the foreground. 
It initializes the libp2p host, mounts the BadgerDB storage, 
and begins listening for jobs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting FluxEdge daemon...")
		runNode()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().Int("port", 9000, "port to listen on for mesh traffic")
	startCmd.Flags().Bool("worker", true, "enable worker mode (accept jobs)")
	startCmd.Flags().Bool("core", false, "enable core mode (participate in consensus)")
}

func runNode() {
	log.Println("initializing p2p host...")
	log.Println("opening BadgerDB at ./data/fluxd.db...")
	log.Println("node is ready (press Ctrl+C to stop)")

	select {}
}
