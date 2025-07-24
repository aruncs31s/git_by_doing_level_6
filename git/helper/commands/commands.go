package commands

var ModifiedFilesCMD = "git diff --name-only HEAD^0"
var UntrackedFilesCMD = "git ls-files --others --exclude-standard"
var GetNumCommitsLocalCMD = "git rev-list --count HEAD"
