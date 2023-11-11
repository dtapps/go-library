// Copyright 2019 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version means the version of reverse
	Version = "0.1+dev"

	reverseFile string
	versionFlag *bool

	rootCmd = &cobra.Command{
		Version: Version,
		Use:     "reverse",
		Short:   "Reverse is a database reverse command line tool",
		Long:    `A flexsible and powerful command line tool to generate codes/docs from databases(SQLITE/Mysql/Postgres/MSSQL)`,
		Run: func(cmd *cobra.Command, args []string) {
			if versionFlag != nil && *versionFlag {
				fmt.Printf("Reverse %s\n", Version)
				return
			}
			if reverseFile == "" {
				fmt.Println("Need reverse file")
				return
			}

			err := reverseFromConfig(reverseFile)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	versionFlag = rootCmd.Flags().BoolP("version", "v", false, "version of the tool")
	rootCmd.Flags().StringVarP(&reverseFile, "file", "f", "", "yml file to apply for reverse")
}

// Execute represnets execute command
func Execute() error {
	return rootCmd.Execute()
}
