
IO.puts("Enter the first number: ")
first_number = IO.gets("") |> String.trim() |> String.to_integer()

IO.puts("Enter the second number: ")
last_number = IO.gets("") |> String.trim() |> String.to_integer()

result = """
#{first_number} + #{last_number} = #{first_number + last_number}
#{first_number} - #{last_number} = #{first_number - last_number}
#{first_number} * #{last_number} = #{first_number * last_number}
#{first_number} / #{last_number} = #{first_number / last_number}
"""
IO.puts(result)