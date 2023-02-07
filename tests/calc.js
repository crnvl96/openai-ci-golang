const a = 3;
const b = 4;

const sum = () => a + b;
const sub = () => a - b;
const mult = () => a * b;
const div = () => a / b;

const calculator = (type) => {
  switch (type) {
    case 'sum':
      return sum();
    case 'sub':
      return sub();
    case 'mult':
      return mult();
    case 'div':
      return div();
    default:
      return 'Invalid operation';
  }
}