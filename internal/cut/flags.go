package cut

import "github.com/spf13/pflag"

// Flags holds the command line flags for the cut command.
type Flags struct {
	Fields    *string
	Delimiter *string
	Separated *bool
}

// InitFlags initializes the command line flags for the cut command.
func InitFlags() Flags {
	return Flags{
		Fields:    pflag.StringP("fields", "f", "", "List of fields to output (1-based index)."),
		Delimiter: pflag.StringP("delimiter", "d", "\t", "Field delimiter"),
		Separated: pflag.BoolP("separated", "s", false, "Output fields separated by the delimiter."),
	}
}
