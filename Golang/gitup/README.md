## Go script to automate repetitive Git commands

Please feel free to download and use the script if you find it helpful (or get the code to build it for your env).
You can also add it to the PATH for ease of use.

### Folder structure

- code is in /src
- releases:

  - /release/win
  - /release/unix&osx

### Usage

#### Options:

-h -> lists options

-m string
Input commit message (Required)

-b string
Input branch name - master or other (Required)

#### It runs:

1. git status
2. git add .
3. git commit -am "commit-message"
4. git push origin "branch-name"
