defmodule Greeter do
  def greet(name) do
    name_greeting = case String.graphemes(name) do
      [] -> "World"
      [head | _] -> "#{name} - #{head}"
    end

    IO.puts("Hello, #{name_greeting}!")
  end
end

IO.write("What is your name? ")
name = IO.gets("") |> String.trim()
Greeter.greet(name)