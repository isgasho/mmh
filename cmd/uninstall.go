/*
 * Copyright 2018 mritd <mritd1234@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"github.com/mritd/mmh/utils"
	"github.com/spf13/cobra"
)

var uninstallDir string
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall mmh",
	Long: `
Uninstall mmh.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Uninstall(uninstallDir)
	},
}

func init() {
	uninstallCmd.PersistentFlags().StringVar(&uninstallDir, "dir", "/usr/bin", "uninstall dir")
	RootCmd.AddCommand(uninstallCmd)
}
