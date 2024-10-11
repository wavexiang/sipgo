git remote add upstream https://github.com/emiago/sipgo.git
git fetch upstream
git checkout main  # 切换到你的主分支，如果不是master分支，请将其替换为你的主分支名称
git merge upstream/main  # 将上游仓库的更新合并到你的主分支
