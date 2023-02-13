fn main() {
  let total_numbers = 60;
  let numbers_drawn = 6;
  let bet_numbers = 25;

  let numerator = (bet_numbers as f64).choose(numbers_drawn as f64);
  let denominator = (total_numbers as f64).choose(numbers_drawn as f64);
  let probability = numerator / denominator;

  println!("The probability of winning the Brazilian Mega-Sena with a bet of 25 numbers is {:.10}%", probability * 100.0);
}
