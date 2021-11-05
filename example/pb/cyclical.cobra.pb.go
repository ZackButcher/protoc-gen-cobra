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

func CyclicalClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Cyclical"),
		Short: "Cyclical service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_CyclicalTestCommand(cfg),
	)
	return cmd
}

func _CyclicalTestCommand(cfg *client.Config) *cobra.Command {
	req := &Foo{
		Bar1: &Bar{},
		Bar2: &Bar{},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Test"),
		Short: "Test RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Cyclical"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Cyclical", "Test"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc *grpc.ClientConn, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewCyclicalClient(cc)
				v := &Foo{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Test(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Bar1.Value, cfg.FlagNamer("Bar1 Value"), "", "")
	cmd.PersistentFlags().StringVar(&req.Bar2.Value, cfg.FlagNamer("Bar2 Value"), "", "")
	cmd.PersistentFlags().StringVar(&req.Value, cfg.FlagNamer("Value"), "", "")

	return cmd
}
