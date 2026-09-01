# Git

Cheatsheet. Edit this file.

## Setup

```
git config --global user.name "name"
git config --global user.email "you@example.com"
```

## Daily

```
git status
git add -p
git commit -m "msg"
git pull --rebase
git push
```

## Branch

```
git switch -c feature/x
git switch main
git branch -d feature/x
```

## Undo

```
git restore --staged FILE
git restore FILE
git reset --soft HEAD~1
```

## Log

```
git log --oneline -20
git diff
git show HEAD
```
