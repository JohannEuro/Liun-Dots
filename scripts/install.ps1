Write-Host 'Liun-Dots - instalador guiado por TUI' -ForegroundColor Cyan
Write-Host ''
Write-Host 'Flujo recomendado:' -ForegroundColor Yellow
Write-Host '1) scoop bucket add liun-dots https://github.com/JohannEuro/Liun-Dots'
Write-Host '2) scoop install liun-dots'
Write-Host '3) liun-dots'
Write-Host ''
Write-Host 'La TUI te permite:' -ForegroundColor Yellow
Write-Host '- Instalación completa (sobrescribe + backup obligatorio)'
Write-Host '- Instalación segura (solo copia faltantes y respeta lo existente)'
Write-Host '- Buscar actualizaciones de forma manual (sin checks en background)'
Write-Host '- Restaurar el último backup si querés volver atrás'
Write-Host ''
Write-Host 'Archivos que toca el instalador:' -ForegroundColor Yellow
Write-Host 'powershell/Microsoft.PowerShell_profile.ps1 -> $HOME\scoop\persist\pwsh\Microsoft.PowerShell_profile.ps1'
Write-Host 'windows-terminal/settings.json -> $HOME\scoop\persist\windows-terminal\settings\settings.json'
Write-Host 'nvim/init.lua -> $env:LOCALAPPDATA\nvim\init.lua'
Write-Host ''
Write-Host 'Si querés revisar tu entorno después, corré scripts/doctor.ps1'
