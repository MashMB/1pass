# Changelog

History of **1Pass** releases with changes description.

## Release 1.1.0

- [CLI] Pretty print for `overview` and `details` commands
- [CLI] Output of `list` command as table
- [CLI] Output of `categories` command as table
- [CLI] Commands `list`, `overview` and `details` works with trashed items (`-t` flag)
- [CLI] Command used to display all available item categories
- [CLI] `list` command with item category filtering (`-c` flag)
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
