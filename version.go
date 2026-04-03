// Copyright (C) 2026 XLR8discovery PBC
// See LICENSE for copying information.

package modify

import (
	"strings"

	"github.com/compose-spec/compose-go/types"
	"github.com/spf13/cobra"

	"xlr8d.io/bigbang/cmd"
	"xlr8d.io/bigbang/pkg/common"
	"xlr8d.io/bigbang/pkg/recipe"
	"xlr8d.io/bigbang/pkg/runtime/runtime"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version <selector>... <version>",
		Short: "set version (docker image tag) for specified services",
		Args:  cobra.MinimumNArgs(2),
		RunE: cmd.ExecuteBigBang(func(st recipe.Stack, rt runtime.Runtime, args []string) error {
			selector, version := common.SplitArgsSelector1(args)
			return cmd.ChangeCompose(st, rt, selector, func(composeService *types.ServiceConfig) error {
				return updateVersion(composeService, version)
			})
		}),
	}
}

func init() {
	cmd.RootCmd.AddCommand(versionCmd())
}

func updateVersion(composeService *types.ServiceConfig, version string) error {
	newImage := strings.ReplaceAll(composeService.Image, "@sha256", "")
	newImage = strings.Split(newImage, ":")[0] + ":" + version
	composeService.Image = newImage
	return nil
}