TASK 1: Setting Up a GIT repository:: 
-----------------------------------------------------------------------
>> logical folder structure so the repo doesn't become a storage.

	/src: Application source code.
	README.txt: The manual for the project.

1. Local Initialization:: 

	# Initialize the local Git repository
	git init -b main

	# Create the standard DevOps directory structure
	mkdir src tests infra scripts

	# Create essential files
	touch README.txt .gitignore

2. Configure the .gitignore

	vim .gitignore
	# Secrets and Environment variables
	.env
	*.pem
	*.key

	# Build artifacts and dependencies
	node_modules/
	dist/
	target/
	*.pyc

	# OS files
	.DS_Store
	Thumbs.db

	:wq!

3. The First Commit

	# Add all files to the staging area
	git add .

	# Create the initial commit
	git commit -m "FIRST COMMIT FOR USER"

4. Connecting to a Remote (GitHub)

	# Using the URL to connect to REMOTE repo with tag "ORIGIN"
	git remote add origin https://github.com/Manjunath-	Kapanaiah/DEVOPS.git

	# Push your code to the remote 'main' branch
	git push -u origin main

-----------------------------------------------------------------------

TASK 2: Forking and CLoning a Project
----------------------------------------------------------------------

Step 1: Cloning

	git clone https://github.com/Manjunath-Kapanaiah/REMOTE.git

	cd remote
	ls


Step 2: Goto source directory and Use fork option (button)

-----------------------------------------------------------------------

TASK 3 & 4: Branching strategies and Pull requests :: 	
----------------------------------------------------------------------

git remote -v
git remote set-url origin https://github.com/YourUsername/FORKED_REPO.git
git pull origin main
git branch new_feature 
git checkout new_feature

echo "this file is to show pull request " >> DOWNSTREAM

git add .
git commit -m "COMMIT ALL"
git push origin main OR
git puch https://github.com/Manjunath-Kapanaiah/DEVOPS.git


----------------------------------------------------------------------

TASK 6: Creating a release and tagging
----------------------------------------------------------------------
git pull origin main

vim app
	This is a demo file...

	This is an added line for another commit and tagging...


git add .
git commit -m "COMMIT FOR VERSION"
git tag Ver-10
git log --oneline
git push https://github.com/Manjunath-Kapanaiah/DEVOPS.git

----------------------------------------------------------------------
TASK 5: HANDLING MERGE CONFLICTS 
TASK 7: HOTFIX
TASK 8: STASHING CHANGES
----------------------------------------------------------------------
vim app

git stash app

git checkout main
vim stash 
	" this is for stash demo"
	:wq!

git add stash
git commit -m "URGENT ISSUE"

git push https://github.com/Manjunath-Kapanaiah/DEVOPS.git

	"" as the other files had the same data we had the merge comflicts also hence merge conflict occured.""

git checkout hotfix

git stash pop

	" adding a new line after stash"

git add app
git commit -m " COMMIT AFTER STASH "

git push https://github.com/Manjunath-Kapanaiah/DEVOPS.git

----------------------------------------------------------------------
TASK 9: CHERRY PICK A COMMIT

----------------------------------------------------------------------

# 1. Switch to the branch that has the fix
git checkout other-branch

# 2. Look at the log to find the commit hash (e.g., a1b2c3d)
git log --oneline


# 1. Switch back to your working branch
git checkout current-branch

# 2. Cherry-pick the specific commit
git cherry-pick a1b2c3d

Open the file and resolve the conflict.

git add <file-name>

git cherry-pick --continue (Instead of git commit).

----------------------------------------------------------------------
TASK 10: INTEGRATING CI

----------------------------------------------------------------------
touch .github/workflows/ci.yml
vim ci.yml

	name: CI Pipeline

# 1. Trigger: Run whenever code is pushed to the 'main' branch
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test-and-lint:
    runs-on: ubuntu-latest  # The "Virtual Machine" GitHub provides

    steps:
      - name: Checkout Code
        uses: actions/checkout@v4  # Pulls your code into the VM

      - name: Setup Environment
        uses: actions/setup-node@v4
        with:
          node-version: '20'

      - name: Install Dependencies
        run: npm install

      - name: Run Code Quality Check (Lint)
        run: npm run lint

      - name: Run Automated Tests
        run: npm test

git add .github/workflows/ci.yml
git commit -m "CI FILE COMMIT"
git push https://github.com/Manjunath-Kapanaiah/DEVOPS.git


