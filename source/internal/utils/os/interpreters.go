package os

var interpreters = map[string][]string{
	"bash":    {"%file%"},
	"dash":    {"%file%"},
	"fish":    {"%file%"},
	"go":      {"run", "%file%"},
	"lua":     {"%file%"},
	"node":    {"-e", "%file%"},
	"nodejs":  {"-e", "%file%"},
	"perl":    {"%file%"},
	"php":     {"%file%"},
	"python":  {"%file%"},
	"python2": {"%file%"},
	"python3": {"%file%"},
	"ruby":    {"%file%"},
	"sh":      {"%file%"},
	"zsh":     {"%file%"},
}

var interpreter_extensions = map[string][]string{
	"bash":    {"sh"},
	"dash":    {"sh"},
	"fish":    {"sh"},
	"go":      {"go"},
	"lua":     {"lua"},
	"node":    {"mjs", "js"},
	"nodejs":  {"mjs", "js"},
	"perl":    {"pl"},
	"php":     {"php", "php4", "php5", "php6", "php7", "php8"},
	"python":  {"py"},
	"python2": {"py"},
	"python3": {"py"},
	"ruby":    {"rb"},
	"sh":      {"sh"},
	"zsh":     {"sh", "zsh"},
}
