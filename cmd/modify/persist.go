// Copyright (C) 2026 XLR8discovery PBC
// See LICENSE for copying information.

package modify

import (
	"github.com/spf13/cobra"

	"xlr8d.io/bigbang/cmd"
	"xlr8d.io/bigbang/pkg/recipe"
	"xlr8d.io/bigbang/pkg/runtime/runtime"
)

func init() {
	cmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "persist <selector>...",
		Short: "Make internal state (database files, storagenode files) persisted between restarts. ",
		Long:  "This is done usually with mounting the directory to the houst. ." + cmd.SelectorHelp,
		Args:  cobra.MinimumNArgs(1),
		RunE:  cmd.ExecuteBigBang(persist),
	})
}

func persist(st recipe.Stack, rt runtime.Runtime, selectors []string) error {
	return runtime.ModifyService(st, rt, selectors, func(s runtime.Service) error {
		rService, err := st.FindRecipeByName(s.ID().Name)
		if err != nil {
			return err
		}

		if rService.Persistence != nil {
			for _, p := range rService.Persistence {
				err := s.Persist(p)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}
