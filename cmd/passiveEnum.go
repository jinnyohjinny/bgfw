/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jinnyohjinny/bgfw/internal"
	"github.com/spf13/cobra"
)

var input string

// passiveEnumCmd represents the passiveEnum command
var passiveEnumCmd = &cobra.Command{
	Use:   "passive-enum",
	Short: "Perform passive enumeration to gather target information",
	Long: `This command is used to perform passive enumeration on a specified target.
Using passive reconnaissance techniques, this command collects publicly available information about the target
without directly interacting with it, making it harder to detect.

Example usage:

passive-enum --input domains.txt

This command will gather information about the domains listed in "domains.txt" using various sources
such as DNS searches, public metadata, and other resources that do not require direct interaction.
This tool is ideal for use during the reconnaissance phase of an application security test.`,

	Run: func(cmd *cobra.Command, args []string) {
		domains, err := internal.ReadDomains(input)
		mode := "subdomains"
		if err != nil {
			panic(err)
		}

		for _, domain := range domains {
			prevResult := internal.Count(domain, "subdomains/subdomains")
			fmt.Printf("Perform subdomain enumeration on %s(%d subdomains)\n", domain, prevResult)
			subfinder, err := internal.ParseJson("tools/subfinder.json", domain, mode)
			if err != nil {
				panic(err)
			}
			internal.Command(subfinder)

			assetfinder, err := internal.ParseJson("tools/assetfinder.json", domain, mode)
			if err != nil {
				panic(err)
			}
			internal.Command(assetfinder)
			newResult := internal.Count(domain, "subdomains/subdomains")
			fmt.Printf("Find %d new subdomains\n", newResult-prevResult)

			// resolve
			prevResult = internal.Count(domain, "subdomains/subdomains-resolved")
			fmt.Printf("Perform resolved subdomain enumeration on %s(%d subdomains)\n", domain, prevResult)
			purednsResolve, err := internal.ParseJson("tools/puredns-resolve.json", domain, mode)
			if err != nil {
				panic(err)
			}
			internal.DownloadLists("resolvers")
			internal.Command(purednsResolve)

			massdnsResolve, err := internal.ParseJson("tools/massdns-resolve.json", domain, mode)
			if err != nil {
				panic(err)
			}
			internal.DownloadLists("resolvers")
			internal.Command(massdnsResolve)

			dnsxResolve, err := internal.ParseJson("tools/dnsx-resolve.json", domain, mode)
			if err != nil {
				panic(err)
			}
			internal.DownloadLists("resolvers")
			internal.Command(dnsxResolve)
			newResult = internal.Count(domain, "subdomains/subdomains-resolved")
			fmt.Printf("Total live subdomains: %d\n", newResult)
			fmt.Printf("Find %d new live subdomains\n", newResult-prevResult)
		}

	},
}

func init() {
	rootCmd.AddCommand(passiveEnumCmd)

	passiveEnumCmd.Flags().StringVar(&input, "input", "", "input")
}
