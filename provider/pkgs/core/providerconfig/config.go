package providerconfig

// Provider Level Configs
// These are the config values that will
// 1. Be taken as input for your provider from stack configs
// 2. Be taken as arguments when creating a new provider object
type Config struct {
	// The fields must be public
	// The pulumi tag is mandatory for basic variables
	// but unnecessary if it is a struct (for structured config).
	// If it is a struct, each variable in the struct must follow the same rules as this object.
	// The tag name may or may not match the variable, but
	// tt is a good practice to make it match.
	RequiredVariable string `pulumi:"requiredVariable"`
	// Fields marked `optional` are optional, so they should have a pointer
	// ahead of their type.
	OptionalVariable *string `pulumi:"optionalVariable,optional"`
}
