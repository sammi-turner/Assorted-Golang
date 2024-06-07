memo = {};

fib = fn(n) {
  if (memo[n] != nil) {
    return memo[n];
  }

  if (n <= 1) {
    return n;
  }

  val = fib(n - 1) + fib(n - 2);
  memo[n] = val;
  val;
};

N = 100;
puts(fib(N));
