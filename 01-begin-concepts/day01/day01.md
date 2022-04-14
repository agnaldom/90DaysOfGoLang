# Concept

```import "fmt"```

import: another keyword meaning "import" or "collect"the Doe library or Doe package.

"fmt" is an abbreviation for format or formatting a standard library that comes with the installation of Go, for building and printing text.

```func main() {} ```

Here we created a function called main:

To create the function, we used the func key, a space, and the function name.

Like the main packege, the main function is also treated specially by the compiler, where it is considered the entry point of the program (Entry point), that is, the instructions in this particular function, specifically within the main package, are run first, unless there is another function Called init ().

```fmt.Print("Welcome to the world\n")```

Here we used the "fmt" package we imported above, to use any package in Go enough to write its name, then a dot, and then the name of the function you want to use from that package, in this case, we simply wanted to print the text "Hello world"and go back to the line, So we used the Print funcion and set it the value "n \hello".

```}```

We closed the main function object using the {sing,}

# The concept of Packages in Go

Packages are a simple way to split your program into smaller pieces divided by purpose or function.
Packages can be referred to as "libraries"or modules.

Go follows the following rules in how to split a program into packages:

A package is often a folder withim your project that contains files named after the package.

You must have at least one package in any program or library.

If the program is an executable program (not a library), then a package called main (ie package main) should be the program's entry as seen in pur first example.

Each Go package must belong to a folder.

A package can exist inside another package but each in a subfolder separately.

The main folder of your project can contain one package, only one, the rest of its packages can be located in subfolder.

# Welcome to the library

What if we want to say "Hello world"in more then one project? Perhps it would be better to make "Hello world"printing in a separate library that we import into other projects. This is just an example of how to build a library in Go. Of course, you don't need a library. All you do is "Hello world"

Let's try to write this in a simple function, call it SayHello and leach it in a special package, call it sayhello.

Let's build this library in a separate project, to be a library that can be broudht in more than one program or other project.

So create a new project folder under Path:

``` $ GOPATH/src/github.com/YOUR_USER/ ```

Create a file name sayhello.go under the sayhello folder.

```$ GOPATH/src/github.com/YOUR_USER/sayhello/sayhello.go```

Now, open the sayhello.go file with yout favorite text editor and the start writing our amazing function. You are not supposed to understand everything right now. The purpose is to explain the concerpt of packages and libraries, yet we will try to explain the instructions of this function:

Note that our function is not much different from the "Hello world"program we first wrote:

We used package sayhello instead package main.

We called the func () function SayHello instead of (funct main).

We now have "Welcome to the World" in the form of a library! But how do we make it feasible? Ie how do we put it in the pkg folder so that the rest of the projects can bring and benefit from it?

All we have to do is typo go install inside the sayhello project folder.

Of course you will not be able to run this libray via the command go run sayhello.go. Let's make sure by writing an automated check of this function!

### In the Go language:

Unit testing can be written easily. It is enough to create a file with the same name as the file you want to test and add test_.

In this case, in the same path as the sayhello project, create a file named sayhello project, create a file named sayhello_test.go next to the same sayhello.go file to test the sayhello package.
