# Comportamiento de la TUI del instalador

## Backup y restauración

- Antes de escribir cualquier archivo, Liun-Dots crea un backup automático.
- Ruta del backup: `%USERPROFILE%\.liun-dots\backups\<timestamp>\`
- `manifest.json` guarda:
  - fecha y hora
  - ruta original
  - ruta del backup
  - si el archivo original ya existía o no
- El rollback se hace desde la opción **Recuperar backup (rollback)**.

### Falla parcial durante instalación

Si la instalación falla a mitad (por permisos, rutas inválidas o archivos bloqueados), el backup ya existe porque se crea antes de escribir.

Recuperación recomendada:

1. Salir o volver al menú principal.
2. Ejecutar **Recuperar backup (rollback)**.
3. Verificar que los archivos se restauraron.
4. Resolver la causa de error y volver a instalar.

Esto evita quedar con una mezcla inconsistente entre archivos nuevos y viejos.

## Actualizaciones

- La búsqueda de actualizaciones es manual desde **Actualizar Liun-Dots**.
- La verificación usa **GitHub Releases API** (`/releases/latest`).
- El resultado se guarda en cache por 24 horas en `%USERPROFILE%\.liun-dots\cache\update-check.json`.
- No se ejecutan chequeos en segundo plano ni al iniciar PowerShell.

## Pantalla de prerequisitos

- Antes de instalar, Liun-Dots muestra si detectó PowerShell 7, Windows Terminal, Neovim y Git.
- Si algo falta, muestra un comando concreto para resolverlo.
- También explica qué pasa si decidís continuar igual: los archivos se pueden copiar con backup, pero el entorno puede quedar incompleto hasta instalar las herramientas faltantes.
