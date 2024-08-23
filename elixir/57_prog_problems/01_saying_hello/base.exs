# 57_prog_problems/01_saying_hello.exs


IO.write("What is your name? ")
name = IO.gets("") |> String.trim()
IO.puts("Hello, #{name}!")
