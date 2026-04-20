return {
  "jonroosevelt/gemini-cli.nvim",
  lazy = true,
  cmd = { "Gemini" },
  keys = {
    { "<leader>ag", "<cmd>Gemini<cr>", desc = "Open Gemini CLI" },
  },
  config = function()
    require("gemini").setup()
  end,
}
