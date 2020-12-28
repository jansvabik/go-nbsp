package nbsp

// CS is a list of setups for Czech language
var CS = []Setup{
	{Regex: `(\d\.?) `, Replacement: "$1" + Repl},
	{Regex: `^(A|I|O|U|K|S|V|Z) `, Replacement: "$1" + Repl},
	{Regex: ` (A|I|O|U|K|S|V|Z|a|i|o|u|k|s|v|z) `, Replacement: " $1" + Repl},
	{Regex: `([A-Z]\.) `, Replacement: "$1" + Repl},
}
