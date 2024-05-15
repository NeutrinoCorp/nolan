# Nolan

Welcome `nolan`, a comprehensive toolkit designed to enhance the development of Golang applications by providing a range of modules and utilities that streamline common programming tasks.

## Features

This library includes a variety of modules that offer functionalities such as:

- **Data Manipulation**: Tools for handling and transforming data structures.
- **Concurrency Utilities**: Enhancements for Go's native concurrency model, making goroutine management easier.
- **Process Utilities**: Easy-to-use modules for process scheduling and management.

## Getting Started

To use this library, you need to have Go installed on your machine. You can install Go by following the instructions on the [official Go website](https://golang.org/dl/).

### Installation

Install the library with the following Go command:

```bash
go get github.com/neutrinocorp/nolan
```

### Usage

Here's a simple example of how to use a module from this library:

#### Data Structures

```go
package main

import (
	"fmt"

	"github.com/neutrinocorp/nolan"
	"github.com/neutrinocorp/nolan/collection/list"
	"github.com/neutrinocorp/nolan/collection/queue"
)

func main() {
	ls := list.NewDoublyLinkedList[int]()
	ls.Add(1)
	ls.Add(2)
	ls.Add(3)
	fmt.Println(ls.ToSlice())
	
	deque := queue.NewDequeList[int](ls)
	fmt.Println(deque.PollLast()) // 3
	fmt.Println(deque.PollLast()) // 2
	fmt.Println(deque.PollLast()) // 1
}
```

#### Processing

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/neutrinocorp/nolan"
	"github.com/neutrinocorp/nolan/collection/list"
	"github.com/neutrinocorp/nolan/collection/queue"
	"github.com/neutrinocorp/nolan/function"
	"github.com/neutrinocorp/nolan/proc"
)

func main() {
	var delegateFunc function.DelegateSafeFuncWithContext[string] = func(ctx context.Context, args string) error {
		log.Printf("some args %s", args)
		return nil
	}

	sched := proc.NewTaskScheduler[string](delegateFunc)
	go func() {
		if err := sched.Start(); err != nil {
			panic(err)
		}
	}()

	if err := sched.SubmitWork("some job"); err != nil {
		panic(err) // this will return an error if the process is already closed
	}
	if err := sched.SubmitWork("some job"); err != nil {
		panic(err)
	}

	// ... os.Stdout will print jobs as delegate will get executed

	// execute graceful shutdown
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second * 30)
	defer cancelFunc()
	if err := sched.Stop(ctx); err != nil {
		panic(err) // will return error if context reaches deadline before actual process termination (i.e., interrupted)
    }
}
```

## Documentation

For more detailed documentation on each module and its functions, visit the [Documentation](https://github.com/neutrinocorp/nolan/wiki) section of this repository.

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the Apache 2.0 License. See `LICENSE` for more information.

## Contact

Your Name - [@neutrinocorp](https://twitter.com/neuntrinocorp) - oss@neutrinocorp.org

Project Link: [https://github.com/neuntrinocorp/nolan](https://github.com/neuntrinocorp/nolan)
