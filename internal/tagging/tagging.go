// Package tagging defines the standard tag set every provisioner applies
// and the helpers used to query AWS/Cloudflare resources by tag.
package tagging

const (
	// ManagedByKey marks a resource as created by this CLI.
	ManagedByKey = "managed-by"
	// ManagedByValue is the constant value applied to ManagedByKey.
	ManagedByValue = "deploy"
	// StackKey carries the user-supplied stack name.
	StackKey = "stack"
)

// StackTags returns the standard tag set every provisioner applies to its resources.
func StackTags(stackName string) map[string]string {
	return map[string]string{
		ManagedByKey: ManagedByValue,
		StackKey:     stackName,
	}
}

// ParseStack extracts the stack name from a resource's tag map. It reports
// false if the resource is not managed by this CLI or has no stack tag.
func ParseStack(tags map[string]string) (string, bool) {
	if tags[ManagedByKey] != ManagedByValue {
		return "", false
	}
	name, ok := tags[StackKey]
	if !ok || name == "" {
		return "", false
	}
	return name, true
}
