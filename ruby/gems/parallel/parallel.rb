require 'parallel'

# each
Parallel.each(1..10) do |i|
  p i * 10 # Parallel execution. unordered output
end

# map
map_result = Parallel.map(1..10) do |item|
  p item * 10 # Parallel execution. unordered output
end
# but expected processing result!
p map_result # [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]