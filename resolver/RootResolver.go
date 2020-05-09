package resolver

// RootResolver main resolver
type RootResolver struct {
	Session string
}

// ResolverError resolver for handling error
type ResolverError interface {
	error
	Extensions() map[string]interface{}
}
