Write-Host '=== Shell ===' -ForegroundColor Cyan
$PSVersionTable.PSVersion
Write-Host ''

Write-Host '=== Core Tools ===' -ForegroundColor Cyan
@('git', 'pwsh', 'nvim', 'rg', 'fd', 'fzf', 'lazygit', 'opencode') | ForEach-Object {
  $cmd = Get-Command $_ -ErrorAction SilentlyContinue
  if ($cmd) {
    Write-Host "[ok] $_ -> $($cmd.Source)"
  } else {
    Write-Host "[missing] $_" -ForegroundColor Red
  }
}

Write-Host ''
Write-Host '=== Important Paths ===' -ForegroundColor Cyan
@(
  "$HOME\scoop\shims",
  "$HOME\scoop\persist\pwsh\Microsoft.PowerShell_profile.ps1",
  "$HOME\scoop\persist\windows-terminal\settings\settings.json",
  "$env:LOCALAPPDATA\nvim\init.lua",
  "$HOME\.config\gentleman\project-roots.json"
) | ForEach-Object {
  Write-Host ("$_ -> " + (Test-Path $_))
}
