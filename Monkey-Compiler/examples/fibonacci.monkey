let fib = fn(x) {
   if (x <= 1) {
     return x;
   }
   fib(x - 1) + fib(x - 2);
};

let N = 15;
puts(fib(N));
