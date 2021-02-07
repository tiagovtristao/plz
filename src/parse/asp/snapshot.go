package asp

import (
	"github.com/tiagovtristao/plz/src/parse/snapshot"
)

func getInitialisedCallSnapshot(s *scope, name string, c *Call) snapshot.Interpreter {
	var buildFileName string
	if s.pkg != nil {
		buildFileName = s.pkg.Filename
	}

	return snapshot.Interpreter{
		BuildFileName: buildFileName,
		InitialisedCall: &snapshot.InitialisedCall{
			Name: name,
			Args: snapshotArguments(s, c.Arguments),
		},
	}
}

func snapshotArguments(s *scope, args []CallArgument) map[string]snapshot.Argument {
	snapshot := make(map[string]snapshot.Argument)

	for _, arg := range args {
		if arg.Name != "" {
			snapshot[arg.Name] = snapshotArgument(s.locals[arg.Name])
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

	return snapshot.Argument{Other: v}
}
