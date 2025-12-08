# collections

[![Go Reference](https://pkg.go.dev/badge/github.com/khajamoddin/collections.svg)](https://pkg.go.dev/github.com/khajamoddin/collections) 
[![Go Report Card](https://goreportcard.com/badge/github.com/khajamoddin/collections)](https://goreportcard.com/report/github.com/khajamoddin/collections)
[![Build Status](https://github.com/khajamoddin/collections/actions/workflows/go.yml/badge.svg)](https://github.com/khajamoddin/collections/actions)
[![Coverage Status](https://codecov.io/gh/khajamoddin/collections/branch/main/graph/badge.svg)](https://codecov.io/gh/khajamoddin/collections)

---

## Overview

**collections** is a modern, idiomatic, zero-boilerplate **generic collections library for Go**.  
It provides high-frequency data structures such as:

- `Set[T]`
- `Deque[T]`
- `PriorityQueue[T]`
- `OrderedMap[K,V]`
- `MultiMap[K,V]`

…and utilities for slices, maps, and iterators.  

This library is designed for clarity, performance, and Go-style minimalism. It fills a gap in the Go ecosystem by providing **well-tested, reusable generic data structures**.

---

## Features

- **Generic, type-safe collections** leveraging Go 1.18+ generics  
- Zero-value usability wherever possible  
- Clean, minimal API following Go conventions  
- Fully tested and benchmarked  
- Ready for real-world production use

---

## Installation

```bash
go get github.com/khajamoddin/collections