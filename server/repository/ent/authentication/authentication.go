// Code generated by ent, DO NOT EDIT.

package authentication

const (
	// Label holds the string label denoting the authentication type in the database.
	Label = "authentication"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// Table holds the table name of the authentication in the database.
	Table = "authentications"
)

// Columns holds all SQL columns for authentication fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldName,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}