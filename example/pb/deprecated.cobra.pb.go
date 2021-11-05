// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	client "github.com/ZackButcher/protoc-gen-cobra/client"
	flag "github.com/ZackButcher/protoc-gen-cobra/flag"
	iocodec "github.com/ZackButcher/protoc-gen-cobra/iocodec"
	proto "github.com/gogo/protobuf/proto"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
)

func DeprecatedClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:        cfg.CommandNamer("Deprecated"),
		Short:      "Deprecated service client",
		Long:       "",
		Deprecated: "deprecated",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_DeprecatedObsoleteCommand(cfg),
	)
	return cmd
}

func _DeprecatedObsoleteCommand(cfg *client.Config) *cobra.Command {
	req := &ObsoleteRequest{}

	cmd := &cobra.Command{
		Use:        cfg.CommandNamer("Obsolete"),
		Short:      "Obsolete RPC client",
		Long:       "",
		Deprecated: "deprecated",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Deprecated"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Deprecated", "Obsolete"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc *grpc.ClientConn, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewDeprecatedClient(cc)
				v := &ObsoleteRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Obsolete(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Unused, cfg.FlagNamer("Unused"), "", "")
	_ = cmd.PersistentFlags().MarkDeprecated(cfg.FlagNamer("Unused"), "deprecated")

	return cmd
}
