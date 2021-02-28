package asp

import (
	"github.com/tiagovtristao/plz/src/parse/snapshot"
)

func getInitialisedCallSnapshot(s *scope, c *Call, name string, args []pyObject) snapshot.Interpreter {
	var buildFileName string
	if s.pkg != nil {
		buildFileName = s.pkg.Filename
	}

	return snapshot.Interpreter{
		BuildFileName: buildFileName,
		InitialisedCall: &snapshot.InitialisedCall{
			Name: name,
			Args: snapshotArguments(s, c, args),
		},
	}
}

func snapshotArguments(s *scope, c *Call, args []pyObject) map[string]snapshot.Argument {
	snapshot := make(map[string]snapshot.Argument)

	for i, arg := range c.Arguments {
		// TODO: Only named arguments supported for now
		if arg.Name != "" {
			if args != nil {
				snapshot[arg.Name] = snapshotArgument(args[i])
			} else {
				snapshot[arg.Name] = snapshotArgument(s.locals[arg.Name])
			}
		}
	}

	return snapshot
}

// Very rudimentary for now...
func snapshotArgument(v pyObject) snapshot.Argument {
	if v.Type() == "str" {
		return snapshot.Argument{
			Str: &snapshot.StringArgument{
				Value: v.(pyString).String(),
			},
		}
	}

	if v.Type() == "list" {
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
	}

	// TODO: Dict type missing

	return snapshot.Argument{Other: v}
}
