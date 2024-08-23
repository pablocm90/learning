list = [
  noun: nil,
  verb: nil,
  adjective: nil,
  adverb: nil
]

defmodule GetWords do
  def get_word(word) do
    IO.write("Enter a #{word}: ")
    IO.gets("") |> String.trim()
  end

  def get_words(list) do
    Enum.reduce(list, %{}, fn {key, _}, acc ->
      Map.put(acc, key, get_word(key))
    end)
  end
end

defmodule StoryTeller do
  def tell_story(words) do
    IO.puts("Do you #{words[:verb]} your #{words[:adjective]} #{words[:noun]} #{words[:adverb]}? That's hilarious!")
  end
end

GetWords.get_words(list) |> StoryTeller.tell_story()