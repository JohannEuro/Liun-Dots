return {
  "nvim-mini/mini.files",
  opts = {
    mappings = {
      go_in = "l",
      go_out = "h",
    },
  },
  config = function(_, opts)
    require("mini.files").setup(opts)
    vim.api.nvim_create_autocmd("User", {
      pattern = "MiniFilesBufferCreate",
      callback = function(args)
        local buf_id = args.data.buf_id
        local map = function(lhs, rhs, desc)
          vim.keymap.set("n", lhs, rhs, { buffer = buf_id, desc = desc })
        end
        map("<Right>", function() require("mini.files").go_in() end, "Go in")
        map("<Left>", function() require("mini.files").go_out() end, "Go out")
        map("L", function() require("mini.files").go_in() end, "Go in")
        map("H", function() require("mini.files").go_out() end, "Go out")
        map("l", function() require("mini.files").go_in() end, "Go in")
        map("h", function() require("mini.files").go_out() end, "Go out")
      end,
    })
  end,
}
