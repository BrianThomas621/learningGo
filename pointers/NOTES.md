# Pointers in Go

Everything in Go is passed by value! Thus a function always gets a copy of the thing being passed.

In this first function, we simply pass an *integer value* to this function, which makes a copy of the `int`. By simply passing in the value (`x` in our code), `x` will remain at whatever value was passed in since we are not passing a memory address but just a value.

```
func pass_intvalue(z int) {
	z = 0
}

...
x := 5
pass_intvalue(x)

// Prints 5 not 0! 
```

In the next function below, here again we pass in a value, but this time we are passing in `*int` (read "pointer to int"), which is a *pointer value* of type `int`. This contains the actual address of the variable passed in. This means that we cannot pass in `x` but the *address of x*: `&x`, which is the address where the value `x` is stored. Inside the function, `*z` "dereferences" the pointer: because the pointer contains the address of the variable, passing a pointer value to this function allows it to update the value, i.e. the address location, that was passed in. Dereferencing essentially says: "give me the value at that location, not the address."

```
func pass_pointervalue(z *int) {
	*z = 0
}

... 
x := 5
pass_pointervalue(&x)

// Prints 0! 
```

## Why use Pointers? 

Rather than passing around potentially large amounts of data, we can simply "point" to the address where the data lives and update that instead, allowing for potentially far more performant code. 

Remember: in Go, everything is *pass by value*: when we pass a memory address, we are passing a value, not a copy of the value.