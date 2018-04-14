package main

import (
    "testing"
)

func TestMain(t *testing.T) {
    res, err := fetchTopre("http://www.topre.co.jp/products/elec/keyboards/index.html")
    if err != nil {
        t.Fatalf("failed test %#v", err)
    }
    if res == "" {
        t.Fatalf("no response")
    }
}
