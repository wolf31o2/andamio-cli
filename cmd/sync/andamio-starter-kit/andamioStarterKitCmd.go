package andamioStarterKit

import (
	"fmt"

	// "andamio-indexer/src/indexer"

	"github.com/spf13/cobra"
)

var AndamioStarterKitCmd = &cobra.Command{
	Use:   "andamio-starter-kit",
	Short: "Example for Cardano Go PBL",
	Long: `

Not yet implemented.

This is a placeholder for an upcoming example.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Andamio network sync...")
		// indexer.RunIndexer()
	},
}

func init() {
	// AdderPublisherCmd.AddCommand(walletToWallet.WalletToWalletCmd)
}
