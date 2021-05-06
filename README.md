# 1Pass - 1Password Linux CLI explorer

1Pass is a command line application that allows to explore 1Password [OPVault](https://support.1password.com/opvault-design/) 
format. Application was created because there is no official 1Password desktop client for Linux users. Only official 
desktop application allows for local passwords store in [OPVault](https://support.1password.com/opvault-design/) format. 
As a long term 1Password user, I don't want to change my passwords manager for anything else only because I am working 
on Linux. For me it is really important to have choice where to store my passwords. I don't feel comfortable with my 
passwords in the cloud. So here is the solution. Before I made the application, every time when I forgot password, I needed 
to use my phone to check how it goes in passwords manager. Now I can do it on a Linux PC. What is more, I can do it in 
Linux way - using CLI only.

<p align="center">
  <img src="assets/1pass-categories.gif">
</p>

<p align="center">
  <img src="assets/1pass-list.gif">
</p>

<p align="center">
  <img src="assets/1pass-details.gif">
</p>

## Installation

Application is available only for Linux x86_64. Right now it is distributed as binary only. Installation process:

1. Go to [GitHub releases](https://github.com/mashmb/1pass/releases) section and download the newest release. 
2. Extract downloaded archive in desired location.
3. Run extracted binary.

If application do not run, it is probably permissions problem. Try *chmod 755 1pass*, this should resolve permissions 
problem. 

For more comfy usage, binary can be added to **$PATH** (in .bashrc file):

```
EXPORT PATH=[path_to_binary_directory]:$PATH
```

Recommended way is to unpack downloaded archive in */usr/bin* location. It will automatically make binary executable 
from terminal with typing just **1pass**.

**IMPORTANT**: release 1.1.0 introduced update service for application (details below)

## Application updates

From release 1.1.0, application has implemented GitHub updates mechanism. Application automatically checks for new updates 
and notifies user about pending one. Application is not updating without user permission. To start update, run:

```
1pass update
```

It is recommended to give application root permissions during update because it is working on computer file system.

Whole update process:

1. Check if there is new release on GitHub.
2. Download newer release to temporary directory (with checksums).
3. Extract downloaded archive.
4. Compare checksums.
5. Replace running binary.
6. Clean cache (temporary files and directories).

## Usage

1Pass is a command line tool, so usage is limited to command variations. First of all type:

```
1pass
```

Command should print overall informations about application.

The most important commands are:
Application provides commands:

```
1pass configure
1pass categories
1pass list [-c <category>] [-t] <path>
1pass overiview [-t] <uid> <path>
1pass details [-t] <uid> <path>
1pass update
1pass version
```

1. **configure** - interactive application configuration (answer the questions), use help command to see what can be configured
2. **categories** - display list of OPVault item categories (for filtering purposes)
3. **list** - display list of items stored in OPVault
4. **overview** - display overview of item without sensitive data
5. **details** - display details of item with sensitive data
6. **update** - check for update and upgrade **1pass**
7. **version** - check actual **1pass** version

Legend:

- **uid** - unique UID of item (obtained with **list** command)
- **path** - path to 1Password OPVault
- **-c** - filter items over category
- **-t** - work on trashed items (archived)

## What is new?

- [CLI] Command used to update application
- [CLI] Notify about new update on every command
- [CLI] Command used to configure application in interactive way (answer the questions)
- [CLI] OPVault path is optional for `list`, `overview` and `details` commands (if not defined, use default one from configuration file)
- [CLI] Pretty print for `overview` and `details` commands
- [CLI] Output of `list` command as table
- [CLI] Output of `categories` command as table
- [CLI] Commands `list`, `overview` and `details` works with trashed items (`-t` flag)
- [CLI] Command used to display all available item categories
- [CLI] `list` command with item category filtering (`-c` flag)
- [API] Configurable updates notification
- [API] Application self update
- [API] Download, extract and validate checksum of new update
- [API] Check for updates on GitHub releases section
- [API] Configurable default OPVault path
- [API] Save application configuration (YAML file in `$HOME/.config/1pass/1pass.yml`)
- [API] Read application configuration (YAML file in `$HOME/.config/1pass/1pass.yml`)
- [API] Merge item overview and details (one structure, full items decoding at once, sensitive data masked in control layer)
- [API] Work with items from trash
- [API] Handle all item categories according to [OPVault design](https://support.1password.com/opvault-design/)

## Releases

Versions of last five releases:

- 1.1.0
- 1.0.0

## What next?

All actual work on project can be tracked in [GitHub issues](https://github.com/MashMB/1pass/issues) or 
[GitHub projects](https://github.com/MashMB/1pass/projects/1).

## Contribution guide

Below you will find instructions how to contribute without changing my software developing workflow. All other type of 
contribution or contribution that do not follow rules, will be "banned" out of the box.

## Bugs

Who likes bugged software? Probably no one. If you will find any bug in application I will be really thankful. Bugs can 
be reported with [GitHub issues](https://github.com/MashMB/1pass/issues). For the proper development cycle, I want to 
investigate reported bug. I also want to reproduce it and prepare technical description of issue to resolve it in next 
release. According to that, bugs should be reported with **Bug** issue template and **bug** label. If I will reproduce 
the bug and find fix for it, issue will be linked to issue with **bugfix** label - ready to work in next releases.

## New feature or change request

I am always opened for new ideas. New feature or change request can be reported with [GitHub issues](https://github.com/MashMB/1pass/issues). 
There is special template named **Request** and **request** label. I will discuss this type of issues. If feature request/change 
is accepted, it will be linked with issue that has **feature** label - ready to implement in next releases.

## Pull requests

It is really great if you want make some changes by yourself in my code base. First of all, some bureaucracy. Before you 
open issue, try to understand existing code. As you can see, this is multi module Go lang project in single GIT repository. 
Do you know why? I am really big fan of hexagonal software architecture because it is easier to control changes in external 
dependencies, core code base is a clear language (Go lang without dependencies), code is hermetic and easy to maintain. 
Even if it is overkill for small projects like this, it is my weapon of choice.

How architecture looks like right now?

- **1pass-core** - core of the application (no external dependencies), business logic
- **1pass-parse** - parsing component used to read data from [OPVault](https://support.1password.com/opvault-design/) format
- **1pass-up** - application update component
- **1pass-term** - component used to handle CLI interaction with application
- **1pass-app** - real application (combines all of above)

Independent Go modules makes it easier to track changes than packages.

If everything is clear to this point and you still want to modify code, open an issue with template **Bug** 
(labels **bug** and **pr**) or **Request** (labels **request** and **pr**). The next step is to describe amount of work 
you want to do. The more detailed description, the better. Git branch name should follow the pattern:

```
<latest_release>/pr/<short_issue_title_with_underscores>/<issue_number>

1.0.0/pr/pretty_item_overview/#99
```

I am trying to use [Conventional commits](https://conventionalcommits.org/en/v1.0.0/), so it is really important that your commits 
also should. Example:

```
feat(#11): get item overview
tests(#11): unit tests of get item overview
```

Look at the repository commits, you will handle it.

Every pull request will be discussed with me and merged to **develop** after acceptation (unit tests are welcome).
