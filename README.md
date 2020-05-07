# GomegaX

This repository contains custom matchers for [Gomega](https://github.com/onsi/gomega).

- `EqualX()` - like `gomega.Equal()` but it uses [`go-cmp`](github.com/google/go-cmp/cmp) instead of `reflect.DeepEqual`
