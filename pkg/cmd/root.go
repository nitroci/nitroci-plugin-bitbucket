/*
Copyright 2021 The NitroCI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/nitroci/nitroci-core/pkg/core/contexts"

	"github.com/spf13/cobra"
)

var runtimeContext *contexts.RuntimeContext

var rootCmd = &cobra.Command{
	Use:   "nitroci-bitbucket",
	Short: "NitroCI - Bitbucket Plugin",
	Long: "NitroCI - Bitbucket Plugin",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	profile, _ := rootCmd.PersistentFlags().GetString("profile")
	verbose, _ := rootCmd.PersistentFlags().GetBool("verbose")
	workspaceDepth, _ := rootCmd.PersistentFlags().GetInt("workspace")
	runtimeContext = contexts.LoadRuntimeContext(profile, verbose, workspaceDepth)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("profile", "p", "default", "set a specific profile")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "displays verbose output")
	rootCmd.PersistentFlags().IntP("workspace", "w", 0, "set current workspace")
}
