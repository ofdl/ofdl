package ofdl

import (
	"context"

	"github.com/defval/di"
	"github.com/spf13/cobra"
)

var App *di.Container

func init() {
	App, _ = di.New()
}

func Apply(options ...di.Option) error {
	return App.Apply(options...)
}

func Provide(constructor di.Constructor, options ...di.ProvideOption) {
	App.Provide(constructor, options...)
}

func ProvideValue(value di.Value, options ...di.ProvideOption) {
	App.ProvideValue(value, options...)
}

func Resolve(ptr di.Pointer, options ...di.ResolveOption) error {
	return App.Resolve(ptr, options...)
}

func Invoke(invocation di.Invocation, options ...di.InvokeOption) error {
	return App.Invoke(invocation, options...)
}

func RunE(runE interface{}) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ProvideValue(cmd)
		ProvideValue(cmd.Context(), di.As(new(context.Context)))
		return Invoke(runE)
	}
}
