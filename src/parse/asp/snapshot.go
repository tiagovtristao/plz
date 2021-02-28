package asp

import (
	"github.com/tiagovtristao/plz/src/parse/snapshot"
)

func getInitialisedCallSnapshot(f *pyFunc, s *scope, c *Call, args []pyObject) snapshot.Interpreter {
	var buildFileName string
	if s.pkg != nil {
		buildFileName = s.pkg.Filename
	}

	return snapshot.Interpreter{
		BuildFileName: buildFileName,
		InitialisedCall: &snapshot.InitialisedCall{
			Name: f.name,
			Args: snapshotNamedArguments(f, s, c, args),
		},
	}
}

// TODO: This should be expanded to other arguments
func snapshotNamedArguments(f *pyFunc, s *scope, c *Call, args []pyObject) map[string]snapshot.Argument {
	snapshot := make(map[string]snapshot.Argument)

	for _, arg := range c.Arguments {
		if arg.Name != "" {
			if idx, exists := f.argIndices[arg.Name]; exists && args != nil {
				snapshot[arg.Name] = snapshotArgument(args[idx])
			} else if localArg, exists := s.locals[arg.Name]; exists {
				snapshot[arg.Name] = snapshotArgument(localArg)
			}
		}
	}

	return snapshot
}

// Very rudimentary for now...
func snapshotArgument(v pyObject) snapshot.Argument {
	if v == nil {
		return snapshot.Argument{}
	}

	switch v.Type() {
	case "str":
		return snapshot.Argument{
			Str: &snapshot.StringArgument{
				Value: v.(pyString).String(),
			},
		}

	case "list":
		list := make([]string, 0)

		for _, item := range v.(pyList) {
			if item.Type() == "str" {
				list = append(list, item.(pyString).String())
			} else {
				return snapshot.Argument{Other: v}
			}
		}

		return snapshot.Argument{
			StrList: &snapshot.StringListArgument{
				Value: list,
			},
		}

	// TODO
	case "dict":
		fallthrough

	default:
		return snapshot.Argument{Other: v}
	}
}
