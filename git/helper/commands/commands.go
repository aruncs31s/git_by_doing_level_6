package commands

var ModifiedFilesCMD = "git diff --name-only HEAD^0"
var UntrackedFilesCMD = "git ls-files --others --exclude-standard"
var GetNumCommitsLocalCMD = "git rev-list --count HEAD"
var GetUserNameCMD = "git config --get user.name"
var GetUserEmailCMD = "git config --get user.email"
