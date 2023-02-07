function generatePassword() {
  let password = '';
  const possibleChars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  
  for (let i = 0; i < 12; i++) {
    password += possibleChars.charAt(Math.floor(Math.random() * possibleChars.length));
  }
  
  return password;
}

console.log(generatePassword());
