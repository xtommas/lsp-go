# `lsp-go`

A Language Server built for learning purposes.

It doesn't do anything special for any particular language, it's attached to markdown files and handles hover (shift + j on Neovim), go to definition (gd on Neovim), has a few code actions (replace VS Code with Neovim and replace VS Code with VS C*de), completion (for the word 'Neovim'), and diagnostics (show an error in the editor when the word 'VS Code' is present). It is focused on helping me understand what tools **do** and **how** they work.

Built and tested with Neovim.

## Usage

In order to actually *use* the LSP inside Neovim, you should change the path of `logger` on main.go, run `go build main.go` and add this to your config:

```lua
local client = vim.lsp.start_client {
  name = 'lsp_test',
  cmd = { '<path_to_your_main_executable>' },
}

if not client then
  vim.notify "you didn't load lsp_test correctly"
  return
end

-- attach our lsp to markdown files
vim.api.nvim_create_autocmd('FileType', {
  pattern = 'markdown',
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})
```
