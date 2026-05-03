vim.g.mapleader = " "
vim.g.maplocalleader = "\\"

local opt = vim.opt
opt.number = true
opt.relativenumber = true
opt.mouse = "a"
opt.signcolumn = "yes"
opt.termguicolors = true
opt.ignorecase = true
opt.smartcase = true
opt.updatetime = 200
opt.timeoutlen = 300
opt.splitright = true
opt.splitbelow = true
opt.scrolloff = 6
opt.sidescrolloff = 8
opt.wrap = false
opt.cursorline = true

local lazypath = vim.fn.stdpath("data") .. "/lazy/lazy.nvim"
if not vim.loop.fs_stat(lazypath) then
  vim.fn.system({
    "git",
    "clone",
    "--filter=blob:none",
    "https://github.com/folke/lazy.nvim.git",
    "--branch=stable",
    lazypath,
  })
end
opt.rtp:prepend(lazypath)

require("lazy").setup({
  { "folke/which-key.nvim", event = "VeryLazy", opts = {} },
  {
    "stevearc/oil.nvim",
    dependencies = { "nvim-tree/nvim-web-devicons" },
    event = "VimEnter",
    cmd = { "Oil" },
    keys = {
      { "<leader>e", "<cmd>Oil<cr>", desc = "Explorador de archivos" },
    },
    opts = {
      default_file_explorer = true,
      skip_confirm_for_simple_edits = true,
      view_options = {
        show_hidden = true,
      },
    },
  },
  {
    "ibhagwan/fzf-lua",
    dependencies = { "nvim-tree/nvim-web-devicons" },
    event = "VimEnter",
    cmd = { "FzfLua" },
    keys = {
      { "<leader><leader>", "<cmd>FzfLua files<cr>", desc = "Buscar archivos" },
      { "<leader>/", "<cmd>FzfLua live_grep<cr>", desc = "Buscar texto" },
      { "<leader>fb", "<cmd>FzfLua buffers<cr>", desc = "Buffers" },
      { "<leader>fr", "<cmd>FzfLua oldfiles<cr>", desc = "Recientes" },
      { "<leader>gs", "<cmd>FzfLua git_status<cr>", desc = "Git status" },
      { "<leader>gc", "<cmd>FzfLua git_commits<cr>", desc = "Git commits" },
    },
    opts = {
      winopts = {
        height = 0.85,
        width = 0.80,
      },
    },
  },
  {
    "nvim-lualine/lualine.nvim",
    dependencies = { "nvim-tree/nvim-web-devicons" },
    event = "VeryLazy",
    opts = {
      options = {
        theme = "auto",
        globalstatus = true,
      },
    },
  },
  {
    "lewis6991/gitsigns.nvim",
    event = { "BufReadPre", "BufNewFile" },
    opts = {},
  },
}, {
  defaults = { lazy = true },
  install = { colorscheme = { "habamax" } },
  checker = { enabled = false },
  change_detection = { notify = false },
})

local keymap = vim.keymap.set

local function current_project_dir()
  if vim.bo.filetype == "oil" then
    local ok, oil = pcall(require, "oil")
    if ok then
      local dir = oil.get_current_dir()
      if dir and dir ~= "" then
        return dir
      end
    end
  end

  local file_dir = vim.fn.expand("%:p:h")
  if file_dir and file_dir ~= "" and vim.fn.isdirectory(file_dir) == 1 then
    return file_dir
  end

  return vim.fn.getcwd()
end

local function open_opencode_here()
  local dir = current_project_dir()
  vim.fn.jobstart({
    "wt",
    "-w",
    "0",
    "split-pane",
    "-p",
    "PowerShell 7",
    "-d",
    dir,
    "opencode",
    dir,
  }, { detach = true })
  vim.notify("OpenCode abierto en: " .. dir)
end

local function copy_current_dir()
  local dir = current_project_dir()
  vim.fn.setreg("+", dir)
  vim.fn.setreg('"', dir)
  vim.notify("Ruta copiada: " .. dir)
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "oil",
  callback = function(args)
    vim.keymap.set("n", "q", "<cmd>q<cr>", { buffer = args.buf, desc = "Cerrar Oil" })
  end,
})

keymap("n", "<C-h>", "<C-w>h", { desc = "Pane izquierda" })
keymap("n", "<C-j>", "<C-w>j", { desc = "Pane abajo" })
keymap("n", "<C-k>", "<C-w>k", { desc = "Pane arriba" })
keymap("n", "<C-l>", "<C-w>l", { desc = "Pane derecha" })

keymap("n", "<leader>w", "<cmd>w<cr>", { desc = "Guardar" })
keymap("n", "<leader>q", "<cmd>q<cr>", { desc = "Cerrar ventana" })
keymap("n", "<leader>Q", "<cmd>qall!<cr>", { desc = "Salir sin guardar" })
keymap("n", "<leader>oc", open_opencode_here, { desc = "Abrir OpenCode aqui" })
keymap("n", "<leader>yp", copy_current_dir, { desc = "Copiar ruta actual" })
