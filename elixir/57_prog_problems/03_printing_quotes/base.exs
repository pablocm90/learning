defmodule Quoter do
  def get_quote() do
    IO.puts("Pray tell me your quote: ")
    IO.gets("") |> String.trim()
  end

  def get_author() do
    IO.puts("Who said that?")
    IO.gets("") |> String.trim()
  end
end

quote = Quoter.get_quote()
author = Quoter.get_author()

IO.puts("\"" <> quote <> "\" by: " <> author )

