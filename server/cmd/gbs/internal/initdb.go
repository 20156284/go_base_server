/*
Copyright © 2020 SliverHorn

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
package internal

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"go_base_server/boot"
	"go_base_server/library/data"
	"go_base_server/library/global"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "go_base_server初始化数据",
	Long:  `go_base_server初始化数据适配数据库情况: 1. mysql完美适配, 2. postgresql未适配, 3. sqlite未适配, 4. sqlserver未适配`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		boot.Viper.Initialize(path)
		boot.Mysql.Check()
		boot.Mysql.Initialize()
		if global.Config.System.DbType == "mysql" {
			if err := data.Initialize(); err == nil {
				color.Info.Println("\n[Mysql] --> 初始化数据成功!\n")
			}
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
	initdbCmd.Flags().StringP("path", "p", "./config/viper.yaml", "自定配置文件路径(绝对路径)")
	initdbCmd.Flags().StringP("type", "t", "mysql", "可选参数为mysql,postgresql,sqlite,sqlserver")
}
