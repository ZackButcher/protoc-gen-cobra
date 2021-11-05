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

func NestedClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Nested"),
		Short: "Nested service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_NestedGetCommand(cfg),
		_NestedGetDeepCommand(cfg),
		_NestedGetOneOfCommand(cfg),
		_NestedGetOneOfDeepCommand(cfg),
	)
	return cmd
}

func _NestedGetCommand(cfg *client.Config) *cobra.Command {
	req := &NestedRequest{
		Top:   &Top{},
		Inner: &NestedRequest_Inner{},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Get"),
		Short: "Get RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested", "Get"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc *grpc.ClientConn, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewNestedClient(cc)
				v := &NestedRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Get(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Top.Value, cfg.FlagNamer("Top Value"), "", "")
	cmd.PersistentFlags().StringVar(&req.Inner.Value, cfg.FlagNamer("Inner Value"), "", "")

	return cmd
}

func _NestedGetDeepCommand(cfg *client.Config) *cobra.Command {
	req := &DeepRequest{
		L0: &DeepRequest_Outer{
			L1: &DeepRequest_Outer_Middle{
				L2: &DeepRequest_Outer_Middle_Inner{},
			},
		},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetDeep"),
		Short: "GetDeep RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested", "GetDeep"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc *grpc.ClientConn, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewNestedClient(cc)
				v := &DeepRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetDeep(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.L0.L1.L2.Value, cfg.FlagNamer("L0 L1 L2 Value"), "", "")

	return cmd
}

func _NestedGetOneOfCommand(cfg *client.Config) *cobra.Command {
	req := &OneOfRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetOneOf"),
		Short: "GetOneOf RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested", "GetOneOf"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc *grpc.ClientConn, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewNestedClient(cc)
				v := &OneOfRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetOneOf(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	Option1 := &OneOfRequest_First{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("Option1"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option1"), func() { req.Choose = &OneOfRequest_Option1{Option1: Option1} })
	cmd.PersistentFlags().StringVar(&Option1.Value, cfg.FlagNamer("Option1 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option1 Value"), func() { req.Choose = &OneOfRequest_Option1{Option1: Option1} })
	Option2 := &OneOfRequest_Second{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("Option2"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option2"), func() { req.Choose = &OneOfRequest_Option2{Option2: Option2} })
	cmd.PersistentFlags().StringVar(&Option2.Value, cfg.FlagNamer("Option2 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option2 Value"), func() { req.Choose = &OneOfRequest_Option2{Option2: Option2} })
	Option3 := &OneOfRequest_Third{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("Option3"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option3"), func() { req.Choose = &OneOfRequest_Option3{Option3: Option3} })
	cmd.PersistentFlags().StringVar(&Option3.Value, cfg.FlagNamer("Option3 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option3 Value"), func() { req.Choose = &OneOfRequest_Option3{Option3: Option3} })

	return cmd
}

func _NestedGetOneOfDeepCommand(cfg *client.Config) *cobra.Command {
	req := &OneOfDeepRequest{
		L0: &OneOfDeepRequest_Outer{},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetOneOfDeep"),
		Short: "GetOneOfDeep RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested", "GetOneOfDeep"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc *grpc.ClientConn, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewNestedClient(cc)
				v := &OneOfDeepRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetOneOfDeep(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	L0Option1 := &OneOfDeepRequest_Outer_First{
		L1: &OneOfDeepRequest_Outer_Middle{
			L2: &OneOfDeepRequest_Outer_Middle_Inner{},
		},
	}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option1"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1"), func() { req.L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: L0Option1} })
	L0Option1L1L2Option1 := &OneOfDeepRequest_Outer_Middle_Inner_First{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option1 L1 L2 Option1"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option1"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: L0Option1}
		L0Option1.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: L0Option1L1L2Option1}
	})
	cmd.PersistentFlags().StringVar(&L0Option1L1L2Option1.Value, cfg.FlagNamer("L0 Option1 L1 L2 Option1 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option1 Value"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: L0Option1}
		L0Option1.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: L0Option1L1L2Option1}
	})
	L0Option1L1L2Option2 := &OneOfDeepRequest_Outer_Middle_Inner_Second{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option1 L1 L2 Option2"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option2"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: L0Option1}
		L0Option1.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: L0Option1L1L2Option2}
	})
	cmd.PersistentFlags().StringVar(&L0Option1L1L2Option2.Value, cfg.FlagNamer("L0 Option1 L1 L2 Option2 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option2 Value"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: L0Option1}
		L0Option1.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: L0Option1L1L2Option2}
	})
	L0Option1L1L2Option3 := &OneOfDeepRequest_Outer_Middle_Inner_Third{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option1 L1 L2 Option3"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option3"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: L0Option1}
		L0Option1.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: L0Option1L1L2Option3}
	})
	cmd.PersistentFlags().StringVar(&L0Option1L1L2Option3.Value, cfg.FlagNamer("L0 Option1 L1 L2 Option3 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option3 Value"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: L0Option1}
		L0Option1.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: L0Option1L1L2Option3}
	})
	L0Option2 := &OneOfDeepRequest_Outer_Second{
		L1: &OneOfDeepRequest_Outer_Middle{
			L2: &OneOfDeepRequest_Outer_Middle_Inner{},
		},
	}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option2"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2"), func() { req.L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: L0Option2} })
	L0Option2L1L2Option1 := &OneOfDeepRequest_Outer_Middle_Inner_First{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option2 L1 L2 Option1"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option1"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: L0Option2}
		L0Option2.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: L0Option2L1L2Option1}
	})
	cmd.PersistentFlags().StringVar(&L0Option2L1L2Option1.Value, cfg.FlagNamer("L0 Option2 L1 L2 Option1 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option1 Value"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: L0Option2}
		L0Option2.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: L0Option2L1L2Option1}
	})
	L0Option2L1L2Option2 := &OneOfDeepRequest_Outer_Middle_Inner_Second{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option2 L1 L2 Option2"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option2"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: L0Option2}
		L0Option2.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: L0Option2L1L2Option2}
	})
	cmd.PersistentFlags().StringVar(&L0Option2L1L2Option2.Value, cfg.FlagNamer("L0 Option2 L1 L2 Option2 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option2 Value"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: L0Option2}
		L0Option2.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: L0Option2L1L2Option2}
	})
	L0Option2L1L2Option3 := &OneOfDeepRequest_Outer_Middle_Inner_Third{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option2 L1 L2 Option3"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option3"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: L0Option2}
		L0Option2.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: L0Option2L1L2Option3}
	})
	cmd.PersistentFlags().StringVar(&L0Option2L1L2Option3.Value, cfg.FlagNamer("L0 Option2 L1 L2 Option3 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option3 Value"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: L0Option2}
		L0Option2.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: L0Option2L1L2Option3}
	})
	L0Option3 := &OneOfDeepRequest_Outer_Third{
		L1: &OneOfDeepRequest_Outer_Middle{
			L2: &OneOfDeepRequest_Outer_Middle_Inner{},
		},
	}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option3"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3"), func() { req.L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: L0Option3} })
	L0Option3L1L2Option1 := &OneOfDeepRequest_Outer_Middle_Inner_First{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option3 L1 L2 Option1"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option1"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: L0Option3}
		L0Option3.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: L0Option3L1L2Option1}
	})
	cmd.PersistentFlags().StringVar(&L0Option3L1L2Option1.Value, cfg.FlagNamer("L0 Option3 L1 L2 Option1 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option1 Value"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: L0Option3}
		L0Option3.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: L0Option3L1L2Option1}
	})
	L0Option3L1L2Option2 := &OneOfDeepRequest_Outer_Middle_Inner_Second{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option3 L1 L2 Option2"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option2"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: L0Option3}
		L0Option3.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: L0Option3L1L2Option2}
	})
	cmd.PersistentFlags().StringVar(&L0Option3L1L2Option2.Value, cfg.FlagNamer("L0 Option3 L1 L2 Option2 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option2 Value"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: L0Option3}
		L0Option3.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: L0Option3L1L2Option2}
	})
	L0Option3L1L2Option3 := &OneOfDeepRequest_Outer_Middle_Inner_Third{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option3 L1 L2 Option3"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option3"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: L0Option3}
		L0Option3.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: L0Option3L1L2Option3}
	})
	cmd.PersistentFlags().StringVar(&L0Option3L1L2Option3.Value, cfg.FlagNamer("L0 Option3 L1 L2 Option3 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option3 Value"), func() {
		req.L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: L0Option3}
		L0Option3.L1.L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: L0Option3L1L2Option3}
	})

	return cmd
}
