# BEW 2.5 Drill 3: Pointers

## Q1 How do you get the memory address of a variable in Golang?

Using the special character "&" in front of a variable will cause it to return the address in memory where that variable is stored.

## Q2 How do you create a new pointer?

`var p *int` <- This is an example of instantiating a new pointer to a memory address that holds an integer value

## Q3 How do you assign a value to a pointer?

`p = &x` <- if `x` is an integer variable, then this is an example of how to assign the memory address of `x` to the pointer `p`, assuming that `p` is declared as described above in Q2.

## Q4 What is the value of x after running the program below?

```
func square(x *float64) {
  *x = *x * *x
}

func main() {
  x := 1.5
  square(&x)
}
```

After running the program, x's value becomes 2.25

## Q5 Interview Practice: Using Pointers

Using what we learned in class today, write a program that contains a function named swap. The swap function has two arguments, x and y, both integers.

When swap is called, the function should swap the values of x and y and return them.

EXAMPLE: If x := 1; and y := 2;, a call to the swap(&x, &y) function should return x=2 and y=1.

Paste your code below:

(Answered in main/main.go)
