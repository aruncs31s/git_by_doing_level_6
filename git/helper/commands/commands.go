package commands

var (
	ModifiedFilesCMD      = "git diff --name-only HEAD^0"
	UntrackedFilesCMD     = "git ls-files --others --exclude-standard"
	GetNumCommitsLocalCMD = "git rev-list --count HEAD"
	GetUserNameCMD        = "git config --get user.name"
	GetUserEmailCMD       = "git config --get user.email"
	GitPullCMD            = "git pull origin main"
	GitAddCMD             = "git add -A"
	GitCommitCMD          = "git commit -m \"Attendance update\""
	GitPushCMD            = "git push origin main"
)
