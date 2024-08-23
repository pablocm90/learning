
defmodule Counter do
  def count_characters(text) do
    String.length(text)
  end

  def get_text() do
    IO.write("Which characters do you want to count? ")
    result = IO.gets("") |> String.trim()
    if result == "" do
      get_text()
    else
      result
    end
  end
end


text = Counter.get_text()

IO.puts("There are #{Counter.count_characters(text)} characters in \"#{text}\".")