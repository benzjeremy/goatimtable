# goatimtable v2.5

Pre-release: this project is still in development.

## A personal note from Jeremy Benz

> I decided we needed to rename untis-go to goatimtable to avoid potential intellectual-property and trademark issues. The previous name directly referenced Untis GmbH's brand. I wanted to make this change early, before it could lead to legal trouble. Same project, new name — thank you for sticking with it.
>
> — Jeremy Benz, project creator

## Changes

- Renamed the application, module, desktop identity, repository, website and documentation.
- Replaced the application icons with an original project-specific design.
- Kept existing local storage paths and legacy decryption compatible. New encrypted data uses PBKDF2-HMAC-SHA256 with 1,000,000 iterations.
- Updated the built-in updater for the new repository and executable names. **Upgrade from v2.4.1 or earlier by downloading manually**; the old updater rejects the new repository path.
- Calendar event identifiers remain stable.

- Fixed a background-sync shutdown data race found by the release race tests.
- Session tokens use 32 random bytes and fail closed on entropy errors.

## Installation

- [Linux AMD64](https://github.com/benzjeremy/goatimtable/releases/download/v2.5/goatimtable-v2.5-linux-amd64.tar.gz)
- [Windows AMD64](https://github.com/benzjeremy/goatimtable/releases/download/v2.5/goatimtable-v2.5-windows-amd64.zip)
- [Website](https://goatimtable.benzjeremy.pp.ua/) · [Manual](https://wiki.benzjeremy.pp.ua/goatimtable/)
- Source: `go install github.com/benzjeremy/goatimtable@v2.5`

## Compatibility

Existing data remains in its legacy directory. After saving new encrypted data, older application versions cannot read that new ciphertext. Back up the local data before downgrading.

## License

GNU GPL-3.0.
