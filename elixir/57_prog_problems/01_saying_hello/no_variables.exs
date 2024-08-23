IO.write "What is your name? "

IO.gets("") |> String.trim() |> (fn name -> IO.puts("Hello, #{name}!") end).()