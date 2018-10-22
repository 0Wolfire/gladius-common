# Gladius Common Utilities

Common utilities with sub-packages that are shared between products and components

## API

## Blockchain

## Database (db)

## Handlers

## Manager

## Requests

## Routing

## Utilities

General collection of useful functions that are shared between Gladius Packages and Projects.

### Config

Gladius Projects use a common configuration pattern. Below are instructions to setup a project using this pattern.

* Dependencies

The only current dependency we use is [viper](https://github.com/spf13/viper). We use the default singleton class to manage state of our projects.

* Setup

```golang
// Dependencies
import (
	"github.com/spf13/viper"
)

// Setup config environment
func initializeConfiguration() {
	// Grab the Gladius Base Directory
	base, err := utils.GetGladiusBase()
	if err != nil {
		log.Warn().Err(err).Msg("Error retrieving base directory")
	}

	// Add config file name and searching
	viper.SetConfigName("gladius-config-example") // looks for gladius-config-example.toml in utils.GetGladiusBase()
	viper.AddConfigPath(base)

	// Setup env variable handling
	viper.SetEnvPrefix("ENV-PREFIX")
	r := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(r)
	viper.AutomaticEnv()

	// Load config
	err = viper.ReadInConfig()
	if err != nil {
		log.Warn().Err(err).Msg("Error reading config file, it may not exist or is corrupted. Using defaults.")
	}

	// Build our config options
	buildOptions(base)
}

// Set up defaults using helper functions
func buildOptions(base string) {
	// Example option
	ConfigOption("Example.SubOption", "Default Value")
	
	// More can go here

	// Add the GladiusBase to access Gladius Folder
	ConfigOption("GladiusBase", base)
}

// Helper function for above
func configOption(key string, defaultValue interface{}) string {
	viper.SetDefault(key, defaultValue)
	return key
}
```