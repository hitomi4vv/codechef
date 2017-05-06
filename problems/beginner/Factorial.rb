gets.to_i.times do
  p, c, n = 1, 0, gets.to_i
  while p <= n; c += n/(p*=5) end
  puts c
end
