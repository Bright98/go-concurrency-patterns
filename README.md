# Golang Concurrency Patterns

<p align="center">
  <img src="images/go-lang.png" width="250"  alt="go-logo"/>
</p>

<p align="center">
This repository serves as a collection of Golang concurrency patterns, and some examples.
It aims to provide developers with a reference for understanding and implementing concurrency in Go.
</p>


## Table of Contents

- [Introduction](#introduction)
- [Concurrency Patterns](#concurrency-patterns)
    - [1. Wait for result](#1-wait-for-result)
    - [2. Wait for task](#2-wait-for-task)
    - [3. Fan-out, Fan-in](#3-fan-out-fan-in)
    - [4. Pooling](#4-pooling)
    - [5. Drop](#5-drop)
    - [6. Cancellation](#6-cancellation)
- [Examples](https://github.com/Bright98/go-concurrency-patterns/tree/main/real-example)

## Introduction

Concurrency is a key feature of the Go programming language, allowing developers to write efficient and scalable programs. This repository focuses on providing practical examples and explanations for various concurrency patterns in Go.

## Concurrency Patterns

### 1. Wait for result

In some scenarios, you may need to wait for the result of a specific goroutine. This can be achieved using channels or other synchronization mechanisms.

[view code](https://github.com/Bright98/go-concurrency-patterns/blob/main/wait-for-result/main.go)

### 2. Wait for task

Use a combination of goroutines and channels to implement a task queue. Workers can pick up tasks and process them concurrently.

[view code](https://github.com/Bright98/go-concurrency-patterns/blob/main/wait-for-task/main.go)

### 3. Fan-out, Fan-in

Fan-out involves starting multiple goroutines to handle incoming data, and fan-in merges their results into a single channel.

[view code](https://github.com/Bright98/go-concurrency-patterns/blob/main/fan-out-in/main.go)

### 4. Pooling

Pooling is a common pattern where a fixed number of worker goroutines are created to handle tasks from a shared queue.

[view code](https://github.com/Bright98/go-concurrency-patterns/blob/main/pooling/main.go)

### 5. Drop

In some scenarios, you may want to drop tasks if the processing capacity is exceeded. This can be achieved by using a buffered channel.

[view code](https://github.com/Bright98/go-concurrency-patterns/blob/main/drop/main.go)

### 6. Cancellation

Implement cancellation of goroutines using the context package. This allows for the graceful termination of concurrent tasks.

[view code](https://github.com/Bright98/go-concurrency-patterns/blob/main/cancellation/main.go)
