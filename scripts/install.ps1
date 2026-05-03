Write-Host 'Liun Dots install helper' -ForegroundColor Cyan
Write-Host ''
Write-Host 'This script is intentionally conservative.'
Write-Host 'Review the files before copying them into your live config paths.'
Write-Host ''
Write-Host 'Suggested copy targets:' -ForegroundColor Yellow
Write-Host 'powershell/Microsoft.PowerShell_profile.ps1 -> $HOME\scoop\persist\pwsh\Microsoft.PowerShell_profile.ps1'
Write-Host 'windows-terminal/settings.json -> $HOME\scoop\persist\windows-terminal\settings\settings.json'
Write-Host 'nvim/init.lua -> $env:LOCALAPPDATA\nvim\init.lua'
Write-Host ''
Write-Host 'Run scripts/doctor.ps1 after copying.'
