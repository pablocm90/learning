module Enumerable
  def each_decreasing_slice(initial_slice_size, decrement = 1, &block)
    next_slice_size = initial_slice_size
    index = 0
    slices = slice_after do |el|
      if (index += 1) == next_slice_size
        initial_slice_size = initial_slice_size > decrement ? initial_slice_size - decrement : 1
        next_slice_size += initial_slice_size
        true
      end
    end

    return slices.enum_for(:each).each(&block) if block_given?

    slices
  end
end


p (1..10).each_decreasing_slice(4).to_a
p (1..10).each_decreasing_slice(4, 2).to_a
p (1..20).each_decreasing_slice(5).to_a
