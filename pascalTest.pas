program CheckArrayOrder;

const
  N = 5;

var
  A: array[1..N] of integer;
  i: integer;
  isGrowingOrder: boolean;

begin
  
  for i := 1 to N do
    readln(A[i]);

  
  isGrowingOrder := true;
  for i := 2 to N do
    if A[i - 1] >= A[i] then
      isGrowingOrder := false;


  if isGrowingOrder then
    writeln('It is')
  else
    writeln('Is not');
end.
