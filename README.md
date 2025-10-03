# Soundex-Based Documentation Finder

## Overview

This project is a command-line utility designed to help users quickly find and display documentation files using phonetic (Soundex) matching. The tool is robust against typos and similar-sounding words, making it easy to search for documentation even if the exact spelling is unknown.

## Features

- **Phonetic Matching:** Uses Soundex encoding to match keywords to documentation files, allowing for fuzzy searching.
- **Keyword-Location Mapping:** Stores mappings between keywords (e.g., "man", "contribute") and their documentation file locations.
- **Add/Update Entries:** Supports adding new documentation entries and updating existing ones with user confirmation.
- **Display Documentation:** Shows documentation files in the terminal using an external command (currently `bat` for syntax highlighting).
- **Extensible:** Modular design allows for easy addition of new commands and integration with other documentation sources.

## Predicted Usage

- `docfinder man` → Displays the "man" documentation.
- `docfinder contribute` → Displays the "contribute" documentation.
- `docfinder add install` → Adds a new documentation entry for "install".
- Handles fuzzy matching for user input, making it tolerant to typos.

## Example Workflow

1. User types a keyword (even with a typo).
2. The tool uses Soundex to find the closest matching documentation entry.
3. The documentation file is displayed in the terminal.

## Future Extensions

- Support for more commands and configuration options.
- Integration with external documentation sources.
- Customizable display commands (e.g., `cat`, `less`, etc.).
