# Liun Dots

Windows-first development workflow for PowerShell 7, Windows Terminal, Neovim, and OpenCode.

This repository is not trying to be a giant dotfiles framework. It is a small, fast, beginner-friendly baseline: open projects quickly, edit with Neovim, split panes in Windows Terminal, and launch OpenCode in the right folder.

## What is included

- PowerShell profile with project navigation helpers.
- Windows Terminal settings focused on panes and PowerShell 7.
- Minimal Neovim setup with Oil, fzf-lua, lualine, gitsigns, and which-key.
- Workflow documentation and QA checklist.
- Install and doctor scripts for setup/review.

## Core idea

Know where you are first:

- In PowerShell, use shell commands like `cproj`, `vproj`, `oc`, `lg`.
- In Neovim/Oil, use Neovim keymaps like `Space e`, `Space Space`, `Space oc`.
- If a project is missing from the selector, add its base folder with `addroot`.

Read `docs/workflow-cheatsheet.md` first. That is the source of truth for daily usage.

## Safety

These configs are Windows-oriented and may contain machine-specific paths. Review before copying them blindly.
