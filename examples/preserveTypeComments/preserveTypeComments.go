// preservetypecomments is a package-level comment. By default, it is preserved.
// with preserveComments configured to "types", it won't be preserved.
package preservetypecomments

import "github.com/google/uuid"

// This is a block comment in the package body. By default, it is preserved.
// With preserveComments configured to "types", it won't be preserved.

// Type comments are preserved, unless configured to "none"
type UserRole = string

const (
	// Const comments are preserved, unless configured to "none"
	UserRoleDefault UserRole = "viewer"
	UserRoleEditor  UserRole = "editor" // Line comments are preserved by default. With preserveComments configured to "types", it won't be preserved.
)

type User struct {
	// Struct field comments are preserved unless configured to "none"
	ID uuid.NullUUID `json:"id" tstype:"string | null"`
}
