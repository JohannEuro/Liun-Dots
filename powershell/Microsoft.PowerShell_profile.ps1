$scoopShims = 'C:\Users\Admin\scoop\shims'
if (Test-Path $scoopShims) {
  $pathParts = $env:PATH -split ';' | Where-Object { $_ -and ($_ -ne $scoopShims) }
  $env:PATH = @($scoopShims; $pathParts) -join ';'
}

$env:EDITOR = 'nvim'
$env:VISUAL = 'nvim'

if (Get-Module -ListAvailable -Name PSReadLine) {
  $psReadLineCommand = Get-Command Set-PSReadLineOption -ErrorAction SilentlyContinue
  if ($psReadLineCommand -and $psReadLineCommand.Parameters.ContainsKey('PredictionSource')) {
    Set-PSReadLineOption -PredictionSource None
  }
  Set-PSReadLineOption -BellStyle None
  Set-PSReadLineOption -HistoryNoDuplicates
}

Set-Alias vim nvim
Set-Alias lg lazygit

if (Test-Path Alias:nv) {
  Remove-Item Alias:nv -Force
}

function ll {
  Get-ChildItem -Force
}

function reload-profile {
  $profilePath = $PROFILE.CurrentUserCurrentHost
  if (Test-Path $profilePath) {
    . $profilePath
  } else {
    Write-Error "No encontre el profile actual en: $profilePath"
  }
}

function ff {
  fd --hidden --strip-cwd-prefix @args
}

function grep {
  rg @args
}

function nv {
  nvim .
}

function oc {
  param(
    [Parameter(Position = 0)]
    [string]$Path = '.'
  )

  opencode $Path
}

function openf {
  param(
    [Parameter(Mandatory = $true, Position = 0)]
    [string]$Path
  )

  Start-Process $Path
}

$crootConfigPath = Join-Path $HOME '.config\gentleman\project-roots.json'

$defaultCrootRoots = @(
  'D:\Users\Admin\source',
  'D:\Users\Admin\projects',
  'D:\Users\Admin\dev',
  'D:\Users\Admin\AndroidStudioProjects',
  'D:\Storage\L\Universidad\Semestre_8\Desarrollo de aplicaciones web',
  'D:\Storage\L\Programacion',
  'D:\Storage\L\PROYECTOS\Things',
  'C:\Users\Admin\source',
  'C:\Users\Admin\projects',
  'C:\Users\Admin\dev'
)

function Save-CrootRoots {
  param(
    [Parameter(Mandatory = $true)]
    [string[]]$Roots
  )

  $configDir = Split-Path $crootConfigPath -Parent
  if (-not (Test-Path $configDir)) {
    New-Item -ItemType Directory -Path $configDir -Force | Out-Null
  }

  $Roots | ConvertTo-Json | Set-Content -Path $crootConfigPath -Encoding utf8
}

function Get-CrootRoots {
  if (-not (Test-Path $crootConfigPath)) {
    Save-CrootRoots -Roots $defaultCrootRoots
  }

  try {
    $roots = Get-Content -Path $crootConfigPath -Raw | ConvertFrom-Json
    if ($roots -isnot [System.Array]) {
      $roots = @($roots)
    }
  } catch {
    $roots = $defaultCrootRoots
    Save-CrootRoots -Roots $roots
  }

  return @(
    $roots |
      Where-Object { $_ -and (Test-Path $_) } |
      Select-Object -Unique
  )
}

function Add-CrootRoot {
  param(
    [Parameter(Mandatory = $true, Position = 0)]
    [string]$Path
  )

  if (-not (Test-Path $Path)) {
    Write-Error "No existe la ruta: $Path"
    return
  }

  $resolvedPath = (Resolve-Path -LiteralPath $Path).Path

  $roots = @()
  if (Test-Path $crootConfigPath) {
    try {
      $loadedRoots = Get-Content -Path $crootConfigPath -Raw | ConvertFrom-Json
      if ($loadedRoots) {
        if ($loadedRoots -is [System.Array]) {
          $roots = @($loadedRoots)
        } else {
          $roots = @($loadedRoots)
        }
      }
    } catch {
      $roots = @($defaultCrootRoots)
    }
  } else {
    $roots = @($defaultCrootRoots)
  }

  if ($roots -contains $resolvedPath) {
    Write-Host "Ya existe en croot: $resolvedPath" -ForegroundColor Yellow
    return
  }

  $roots += $resolvedPath
  Save-CrootRoots -Roots ($roots | Select-Object -Unique)
  Write-Host "Agregada a croot: $resolvedPath" -ForegroundColor Green
}

Set-Alias addroot Add-CrootRoot

function roots {
  Get-CrootRoots
}

function Select-ProjectPath {
  $roots = Get-CrootRoots

  if ($roots.Count -eq 0) {
    Write-Host 'No encontre carpetas base de proyectos todavia.'
    return $null
  }

  $projects = @(
    foreach ($root in $roots) {
      $root
      fd --hidden --follow --type d --max-depth 2 . $root
    }
  )

  return ($projects | Sort-Object -Unique | fzf --height 50% --layout reverse --prompt 'Projects > ')
}

function croot {
  $projects = Get-CrootRoots

  if ($projects.Count -eq 0) {
    Write-Host 'No encontre carpetas base de proyectos todavia.'
    return
  }

  $target = $projects | fzf --height 40% --layout reverse --prompt 'Roots > '
  if ($target) {
    Set-Location $target
  }
}

function cproj {
  $target = Select-ProjectPath

  if ($target) {
    Set-Location $target
  }
}

function vproj {
  $target = Select-ProjectPath

  if ($target) {
    Set-Location $target
    nvim $target
  }
}

function oproj {
  $target = Select-ProjectPath

  if ($target) {
    Set-Location $target
    opencode $target
  }
}

function devcheck {
  Write-Host '=== Shell ===' -ForegroundColor Cyan
  $PSVersionTable.PSVersion
  Write-Host ''

  Write-Host '=== Core Tools ===' -ForegroundColor Cyan
  @('git', 'nvim', 'rg', 'fd', 'fzf', 'lazygit', 'opencode') | ForEach-Object {
    $cmd = Get-Command $_ -ErrorAction SilentlyContinue
    if ($cmd) {
      Write-Host "[ok] $_ -> $($cmd.Source)"
    } else {
      Write-Host "[missing] $_" -ForegroundColor Red
    }
  }

  Write-Host ''
  Write-Host '=== Project Roots ===' -ForegroundColor Cyan
  Get-CrootRoots | ForEach-Object {
    Write-Host ("$_ -> " + (Test-Path $_))
  }
}
