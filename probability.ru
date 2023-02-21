def dice_probability
  total_outcomes = 36
  successful_outcomes = 6 # There are 6 possible combinations that add up to 7 (1+6, 2+5, 3+4, 4+3, 5+2, and 6+1)
  probability = successful_outcomes.to_f / total_outcomes
  probability
end

puts "The probability of throwing two dice and getting exactly 7 is #{dice_probability}."
