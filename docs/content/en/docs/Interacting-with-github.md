## Interacting with github

The documentation source is stored on github.com
and you use the standard github facilities to modify it.
The following notes are provided for people who need a little help.

### Fork and clone the repository

Perform the following steps to create a copy of this repository on your local machine:

1. Fork the Keptn repository:

     * Log into GitHub (or create a GitHub account and then log into it).
     * Go to the [Keptn lifecycle-toolkit repository](https://github.com/keptn/lifecycle-toolkit).
     * Click the **Fork** button at the top of the screen.
     * Choose the user for the fork from the options you are given,
       usually your GitHub ID.

   A copy of this repository is now available in your GitHub account.

2. Get the string to use when cloning your fork:

     * Click the green "Code" button on the UI page.
     * Select the protocol to use for this clone (either HTTPS or SSH).
     * A box is displayed that gives the URL for the selected protocol.
       Click the icon at the right end of that box to copy that URL.

3. Run the **git clone** command from the shell of a local directory
    to clone the forked repository to a directory on your local machine,
    pasting in the URl you saved in the previous step:

    ```console
    git clone https://github.com/<UserName>/lifecycle-toolkit
    ```

    or

    ```console
    git clone git@github.com:<UserName>/lifecycle-toolkit.git
    ```

    Where <*UserName*> is your GitHub username.
    The lifecycle-toolkit directory is now available in the local directory.

4. Associate your clone with `upstream`.
   To do this, use the same string you used to clone your fork.

   * Be sure that you are in the root folder of the project
     and run *git status* to confirm that you are on the `main` branch.
   * Type the following to associate `upstream` with your clone,
     pasting in the string for the main repo that you copied above.:

     ```console
     git remote add upstream https://github.com/keptn/lifecycle-toolkit.git 
     ```

### Create a new branch and make your changes

Now you have your upstream setup in your local machine.
You need to create a local branch where you will make your changes.
Be sure that your branch is based on and sync'ed with `main`,
unless you intend to create a derivative PR:

The following sequence of commands does that:

```console
git checkout main
git pull upstream main
git push origin main
git checkout -b <my-new-branch>
```

Now you can make your changes, build them locally to check formatting,
then create a PR to add these changes to the documentation set

### Submitting new content through a pull request

If you have forked and cloned the repository,
you can modify the documentation or add new material
by editing the markdown file(s) in the local clone of your fork
and then submitting a *pull request (PR)*.

Note that you can also modify the source using the GitHub editor.
This is very useful when you want to fix a typo or make a small editorial change
but, if you are doing significant writing,
it is generally better to work on files in your local clone.

1. Execute the following and check the output to ensure that your branch is set up correctly:

   ```console
   git status
   ```

1. Do the writing you want to do in your local branch, checking the formatted version at `localhost:1314/docs-dev`.

1. When you have completed the writing you want to do, close all files in your branch and run `git status` to confirm
   that it correctly reflects the files you have modified, added, and deleted.

1. Add and commit your changes.  Here, we commit all modified files but you can specify individual files to the
   `git add` command.
The `git commit -s` command commits the files and signs that you are contributing this intellectual property to the
Keptn project.

   ```console
   git add .
   git commit -s
   ```

   Use vi commands to add a description of the PR
   (should be 80 characters or less) to the commit.
   The title text should be prefixed with `docs:`
   to conform to our semantic commit scheme.
   This title is displayed as the title of the PR in listings.
   You can add multiple lines explaining the PR here but, in general,
   it is better to only supply the PR title here;
   you can add more information and edit the PR title
   when you create the PR on the GitHub UI page.

1. Push your branch to github:
   * If you cloned your fork to use SSH, the command is:

      ```console
      git push --set-upstream origin <branch-name>
      ```

      > **Note**
      You can just issue the `git push` command.
      Git responds with an error message that gives you the full command line to use; you can copy this command and
      paste it into your shell to push the content.
   * If you cloned your fork to use HTTP, the command is:

      ```console
      git push <need options/arguments>
      ```

1. Create the PR by going to the [lifecycle-toolkit](https://github.com/keptn/lifecycle-toolkit) GitHub repository.
   * You should see a yellow shaded area that says something like:

     ```console
     <GitID>:<branch> had recent pushes less than a minute ago
     ```

   * Click on the button in that shaded area marked:

     ```console
     Compare & pull request
     ```

   * Check that the title of the PR is correct; click the "Edit" button to modify it.
Add "WIP" (Work in Progress) or "Draft" to the title if the PR is not yet ready for general review.
   * Add a full description of the work in the PR, including any notes for reviewers, a reference to the relevant GitHub
      issue (if any), and tags for specific people (if any) who may be interested in this PR.
   * Carefully review the changes GitHub displays for this PR to ensure that they are what you want.
   * Click the green "Create pull request" button to create the PR.
      You may want to record the PR number somewhere for future reference although you can always find the PR in the
      GitHub lists of open and closed PRs.
   * GitHub automatically populates the "Reviewers" block.
   * If this PR is not ready for review, click the "Still in progress? Convert to draft" string under the list of
      reviewers.
      People can still review the content but can not merge the PR until you remove the "Draft" status.
   * The block of the PR that reports on checks will include the following item:

     ```console
     This pull request is still a work in progress
     Draft pull requests cannot be merged.
     ```

   * When the PR is ready to be reviewed, approved, and merged, click the "Ready to review" button to remove the "Draft"
      status. Then, if you added "WIP" or "Draft" to the PR title, remove it now.

1. Your PR should be reviewed within a few days.  Watch for any comments that may be added by reviewers and implement or
   respond to the recommended changes as soon as possible.

   * If a reviewer makes a GitHub suggestion and you agree with the change, just click "Accept this change" to create a
      commit for that modification.
      You can also group several suggestions into a single commit using the GitHub tools.
   * You can make other changes using the GitHub editor or you can work in your local branch to make modifications.

      * If changes have been made using the GitHub editor, you will need to do a `git pull` request to pull those
         commits back to your local branch before you push the new changes.
      * After modifying the local source, issue the `git add .`, `git commit`, and `git push` commands to push your
         changes to github.

1. When your PR has the appropriate approvals, it will be merged and the revised content should be published on the
   website within a few minutes.

1. When your PR has been approved and merged,
   you can delete your local branch with the following command:

   ```console
   git branch -d <branch-name>
   ```