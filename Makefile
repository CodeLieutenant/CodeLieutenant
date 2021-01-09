
.PHONY: git-setup
git-setup:
	@git config user.name GitHub
	@git config user.email noreply@github.com
	@git remote set-url origin https://x-access-token:${GITHUB_TOKEN}@github.com/dusanmalusev.git


.PHONY: commit
commit:
	@git add .
ifneq ($(shell git status --porcelain),)
	@git commit --author "github-actions[bot] <github-actions[bot]@users.noreply.github.com>" --message "${MESSAGE}"
	@git push
endif
