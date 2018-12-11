## Git commands:
For more info check [official documentation](https://git-scm.com/book/en/v2)

*  create new branch in git
    ```
    $ git branch <branch-name>
    ```
*  create and switch to new git branch
    ```
    $ git checkout -b <branch-name>
    ```
*  check git log with graphical mode
    ```
    $ git log --graph --decorate --pretty=oneline --abbrev-commit
    ```
*  force git push 
    ```
    $ git push --force origin <branch-name>
    ```
* squash commits
    
    1. rebase your commits: 
        ```
        $ git rebase -i HEAD~[NUMBER OF COMMITS]
        ```
        or
        ```
        $ git rebase -i [SHA]
        ```
    2. edit commit file and choose which commits are you goint to `squash`

    3. If you have previously pushed your code to a remote branch, you will need to force push.
        ```
        $ git push origin branchName --force
        ```