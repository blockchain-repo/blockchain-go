# How to Contribute to the Project

### Step 0 - Decide on an Issue to Resolve, or Create One

We ask that all pull requests should resolve, address, or fix an existing issue. If there is no existing issue, then you should create one first. That way there can be commentary and discussion first, and you can have a better idea of what to expect when you create a corresponding pull request.

When you submit a pull request, please mention the issue (or issues) that it resolves, e.g. "Resolves #123".

Exception: hotfixes and minor changes don't require a pre-existing issue, but please write a thorough pull request description.

### Step 1 - Fork the Project

In your web browser, go to the repository and click the `Fork` button in the top right corner. This creates a new Git repository named  in _your_ Git account.

### Step 2 - Clone Your Fork

(This only has to be done once.) In your local terminal, use Git to clone _your_ repository to your local computer. Also add the original Git repository as a remote named `upstream` (a convention):
```text
git clone git@github.com:your-github-username/unichain.git
cd unichain
git remote add upstream git@github.com:uni-ledger/unichain.git
```

### Step 3 - Fetch and Merge the Latest from `upstream/master`

Switch to the `master` branch locally, fetch all `upstream` branches, and merge the just-fetched `upstream/master` branch with the local `master` branch:
```text
git checkout master
git fetch upstream
git merge upstream/master
```

### Step 4 - Create a New Branch for Each Bug/Feature

If your new branch is to **fix a bug** identified in a specific GitHub Issue with number `ISSNO`, then name your new branch `bug/ISSNO/short-description-here`. For example, `bug/67/fix-leap-year-crash`.

If your new branch is to **add a feature** requested in a specific GitHub Issue with number `ISSNO`, then name your new branch `feat/ISSNO/short-description-here`. For example, `feat/135/blue-background-on-mondays`.

Otherwise, please give your new branch a short, descriptive, all-lowercase name.
```text
git checkout -b new-branch-name
```

### Step 5 - Make Edits, git add, git commit

With your new branch checked out locally, make changes or additions to the code or documentation. Remember to:

* follow Code Style Guide.
* write and run tests for any new or changed code.
* add or update documentation as necessary.

As you go, git add and git commit your changes or additions, e.g.
```text
git add new-or-changed-file-1
git add new-or-changed-file-2
git commit -m "Short description of new or changed things"
```

You will want to merge changes from upstream (i.e. the original repository) into your new branch from time to time, using something like:
```text
git fetch upstream
git merge upstream/master
```

### Step 6 - Push Your New Branch to origin

Make sure you've commited all the additions or changes you want to include in your pull request. Then push your new branch to origin (i.e. _your_ remote unichain repository).
```text
git push origin new-branch-name
```

### Step 7 - Create a Pull Request

Go to the GitHub website and to _your_ remote repository

See [GitHub's documentation on how to initiate and send a pull request](https://help.github.com/articles/using-pull-requests/). Note that the destination repository should be `uni-ledger/unichain` and the destination branch will be `master` (usually, and if it's not, then we can change that if necessary).

Someone will then merge your branch or suggest changes. If we suggest changes, you won't have to open a new pull request, you can just push new code to the same branch (on `origin`) as you did before creating the pull request.

## Quick Links

* [General GitHub Documentation](https://help.github.com/)

(Note: GitHub automatically links to CONTRIBUTING.md when a contributor creates an Issue or opens a Pull Request.)
