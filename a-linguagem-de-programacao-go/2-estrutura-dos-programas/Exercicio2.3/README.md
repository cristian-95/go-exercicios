# Exercicio 2.3

Comparação do desempenho das duas versões de  ```PopCount```, sendo a segunda (```PopCount2```) feita com um loop for no lugar da expressão única.

```
goos: linux
goarch: amd64
pkg: popcount/popcount
cpu: AMD Ryzen 5 2600 Six-Core Processor            
BenchmarkPopCount-12            1000000000               0.2913 ns/op
BenchmarkPopCount2-12           193485337                5.985 ns/op
PASS
ok      popcount/popcount       2.118s
```