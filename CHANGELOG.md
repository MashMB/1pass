# Changelog

History of **1Pass** releases with changes description.

## Release 1.3.1

- (FIX) [API] Do not parse fields with value but without name

## Release 1.3.0 [14.05.2021]

- [GUI] Display notification about update
- [GUI] Check for update asynchronously
- [GUI] Display actual application version
- [GUI] Contextual help for item details widget
- [GUI] Contextual help for items widget
- [GUI] Contextual help for 1pass widget
- [GUI] Display decoded item details
- [GUI] Scroll item details left
- [GUI] Scroll item details right
- [GUI] Scroll item details up
- [GUI] Scroll item details down
- [GUI] Display item overview (no sensitive data)
- [GUI] Scroll items up
- [GUI] Scroll items down
- [GUI] Display list of items assigned to category
- [GUI] Scroll categories menu up
- [GUI] Scroll categories menu down
- [GUI] Display categories that has at least one item assigned (include dynamic all and trashed category)
- [GUI] Lock vault with shortcut
- [GUI] Unlock vault with master password
- [GUI] Dialog used to display errors
- [GUI] Password prompt (with password mask)
- [GUI] Validate OPVault path before GUI launch
- [CLI] Run 1Pass in GUI mode by default (command `1pass`)
- [API] Handle trashed flag in simplified item data structure
- [API] Count OPVault items over category and trashed flag

## Release 1.2.0 [09.05.2021]

- [CLI] Display application update stages
- [CLI] Commands that use default OPVault path (`list`, `overview` and `details`) will prompt for configuration on first run
- [CLI] `update` command wants user confirmation when new version is available
- [CLI] Display changelog of new version on `update` command
- [CLI] Force update check on `update` command
- [CLI] OPVault path as `list`, `overview` and `details` commands flag (`-v [path]`)
- [CLI] No results message for filtering in `list` command
- [CLI] `list` command with items filtering over title (`-n` flag)
- [API] Handle application update stages
- [API] Check if configuration is available (already exists)
- [API] Force update checking
- [API] Get new release changelog from GitHub during update check
- [API] Check for update only once per time period
- [API] Store time stamp of last successful update check
- [API] Configurable update check period
- [API] Configurable HTTP update timeout
- [API] Items filtering over title
- (FIX) [API] Vault lock clears decoded items memory
- (FIX) [API] Validate OPVault path before password prompt

## Release 1.1.0 [06.05.2021]

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

## Release 1.0.0 [25.04.2021]

- [CLI] Prompt for user master password without input displaying
- [CLI] Command used to display single login details
- [CLI] Command used to display single login overview
- [CLI] Command used to list saved in OPVault format logins
- [CLI] Command used to display actual application version
- [API] Get and decode single login details (sensitive data)
- [API] Get and decode single login overview (no sensitive data)
- [API] Get (decode) list of logins stored in OPVault
- [API] Unlock OPVault data format with usage of master password
