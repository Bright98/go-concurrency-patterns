# Golang Concurrency Patterns

<div  style="display: flex; width: 100%; justify-content: center;">
<img src="https://banner2.cleanpng.com/20180731/tzw/kisspng-gopher-docker-computer-programming-clojure-5b60bcbbaea281.7058312815330664277153.jpg" width="300" style="background-color: transparent;" />
</div>

This repository serves as a collection of Golang concurrency patterns, and some examples.
It aims to provide developers with a reference for understanding and implementing concurrency in Go.


## Table of Contents

- [Introduction](#introduction)
- [Concurrency Patterns](#concurrency-patterns)
    - [1. Wait for result](#1-wait-for-result)
    - [2. Wait for task](#2-wait-for-task)
    - [3. Fan-out, Fan-in](#3-fan-out-fan-in)
    - [4. Pooling](#4-pooling)
    - [5. Drop](#5-drop)
    - [6. Cancellation](#6-cancellation)
- [Examples](#examples)

## Introduction

Concurrency is a key feature of the Go programming language, allowing developers to write efficient and scalable programs. This repository focuses on providing practical examples and explanations for various concurrency patterns in Go.

## Concurrency Patterns

### 1. Wait for result
   In some scenarios, you may need to wait for the result of a specific goroutine. This can be achieved using channels or other synchronization mechanisms.

### 2. Wait for task
Use a combination of goroutines and channels to implement a task queue. Workers can pick up tasks and process them concurrently.

### 3. Fan-out, Fan-in

Fan-out involves starting multiple goroutines to handle incoming data, and fan-in merges their results into a single channel.

### 4. Pooling

Pooling is a common pattern where a fixed number of worker goroutines are created to handle tasks from a shared queue.

### 5. Drop

In some scenarios, you may want to drop tasks if the processing capacity is exceeded. This can be achieved by using a buffered channel.

### 6. Cancellation

Implement cancellation of goroutines using the context package. This allows for the graceful termination of concurrent tasks.
