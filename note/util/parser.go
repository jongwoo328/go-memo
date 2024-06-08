package util

func trimQuote(option string) string {
	if option[0] == '"' && option[len(option)-1] == '"' {
		return option[1 : len(option)-1]
	} else if option[0] == '\'' && option[len(option)-1] == '\'' {
		return option[1 : len(option)-1]
	}
	return option
}
func ParseCommand(command string) (string, []string) {
	if len(command) == 0 {
		return "", []string{}
	}

	var commandName string = ""
	var commandOptions []string

	var lastCommitedIndex int = 0
	var inQuote bool = false
	var quoteChar string = ""
	var commandParsed bool = false
	for i := 0; i < len(command); i++ {
		var char = command[i]
		if char == '"' || char == '\'' {
			if inQuote {
				if quoteChar == string(char) {
					inQuote = false
				} else {
					continue
				}
			} else {
				inQuote = true
				quoteChar = string(char)
				continue
			}
		}
		if char == ' ' {

			if !inQuote {
				if !commandParsed {
					commandName = command[lastCommitedIndex:i]
					commandParsed = true
				} else {
					option := trimQuote(command[lastCommitedIndex:i])
					commandOptions = append(commandOptions, option)
				}
				lastCommitedIndex = i + 1
			} else {
				continue
			}
		} else {
			continue

		}

	}
	if lastCommitedIndex < len(command) {
		option := trimQuote(command[lastCommitedIndex:])
		commandOptions = append(commandOptions, option)
	}
	return commandName, commandOptions
}
