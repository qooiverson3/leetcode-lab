/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// twosumCmd represents the twosum command
var twosumCmd = &cobra.Command{
	Use:   "twosum",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("twosum called")
	},
}

func init() {
	rootCmd.AddCommand(twosumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// twosumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// twosumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func twoSum(nums []int, target int) []int {
	temp := map[int]int{}

	for k, v := range nums {
		temp[v] = k
	}

	for k, v := range nums {
		gap := target - v
		if key, ok := temp[gap]; ok && key != k {
			return []int{k, key}
		}
	}
	return []int{}
}
