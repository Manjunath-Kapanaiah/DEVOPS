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

TASK 5: HANDLING MERGE CONFLICTS 
----------------------------------------------------------------------


