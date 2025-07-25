package commands

var (
	ModifiedFilesCMD      = "git diff --name-only HEAD^0"
	UntrackedFilesCMD     = "git ls-files --others --exclude-standard"
	GetNumCommitsLocalCMD = "git rev-list --count HEAD"
	GetUserNameCMD        = "git config --get user.name"
	GetUserEmailCMD       = "git config --get user.email"
)
