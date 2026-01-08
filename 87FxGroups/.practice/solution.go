// FX Value Groups allow you to collect multiple values of the same type
// and inject them as a slice. This is useful for plugin-style architectures
// where you want to register multiple implementations.

package main

import (
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Plugin is an interface that all plugins must implement
type Plugin interface {
	Name() string
	Execute() string
}

// LoggerPlugin logs messages
type LoggerPlugin struct {
	log *zap.Logger
}

func NewLoggerPlugin(log *zap.Logger) *LoggerPlugin {
	return &LoggerPlugin{log: log}
}

func (p *LoggerPlugin) Name() string {
	return "Logger"
}

func (p *LoggerPlugin) Execute() string {
	return "Logging some information..."
}

// GreeterPlugin greets users
type GreeterPlugin struct{}

func NewGreeterPlugin() *GreeterPlugin {
	return &GreeterPlugin{}
}

func (p *GreeterPlugin) Name() string {
	return "Greeter"
}

func (p *GreeterPlugin) Execute() string {
	return "Hello from Greeter Plugin!"
}

// CalculatorPlugin does calculations
type CalculatorPlugin struct{}

func NewCalculatorPlugin() *CalculatorPlugin {
	return &CalculatorPlugin{}
}

func (p *CalculatorPlugin) Name() string {
	return "Calculator"
}

func (p *CalculatorPlugin) Execute() string {
	return "2 + 2 = 4"
}

// AsPlugin annotates a constructor to add it to the "plugins" group
func AsPlugin(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Plugin)),
		fx.ResultTags(`group:"plugins"`),
	)
}

// PluginManager manages all registered plugins
type PluginManager struct {
	plugins []Plugin
}

// NewPluginManager creates a manager with all registered plugins
// The struct tag tells FX to inject all values from the "plugins" group
func NewPluginManager(plugins []Plugin) *PluginManager {
	return &PluginManager{plugins: plugins}
}

// RunAll executes all registered plugins
func (pm *PluginManager) RunAll() {
	fmt.Printf("Running %d plugins:\n", len(pm.plugins))
	for _, plugin := range pm.plugins {
		fmt.Printf("  [%s]: %s\n", plugin.Name(), plugin.Execute())
	}
}

func main() {
	fx.New(
		fx.Provide(
			zap.NewExample,
			AsPlugin(NewLoggerPlugin),
			AsPlugin(NewGreeterPlugin),
			AsPlugin(NewCalculatorPlugin),
			fx.Annotate(
				NewPluginManager,
				fx.ParamTags(`group:"plugins"`),
			),
		),
		fx.Invoke(func(pm *PluginManager) {
			pm.RunAll()
		}),
	).Run()
}

// Notes:
// - Value Groups collect multiple values with the same group tag
// - fx.ResultTags(`group:"name"`) adds a value to a group
// - fx.ParamTags(`group:"name"`) receives all values from a group as a slice
// - fx.As() casts the result to an interface type
// - This pattern is great for plugin systems, middleware chains, and route handlers





