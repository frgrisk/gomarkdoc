<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# unexported

```go
import "github.com/frgrisk/gomarkdoc/testData/unexported"
```

Package unexported contains some simple code to exercise basic scenarios for documentation purposes.

## Index

- [type Num](<#Num>)
  - [func AddNums\(num1, num2 Num\) Num](<#AddNums>)
  - [func addInternal\(num1, num2 Num\) Num](<#addInternal>)
  - [func \(n Num\) Add\(num Num\) Num](<#Num.Add>)


<a name="Num"></a>
## type [Num](<https://github.com/frgrisk/gomarkdoc/blob/master/testData/unexported/main.go#L8>)

Num is a number.

It is just a test type so that we can make sure this works.

```go
type Num int
```

<a name="AddNums"></a>
### func [AddNums](<https://github.com/frgrisk/gomarkdoc/blob/master/testData/unexported/main.go#L16>)

```go
func AddNums(num1, num2 Num) Num
```

AddNums adds two Nums together.

<a name="addInternal"></a>
### func [addInternal](<https://github.com/frgrisk/gomarkdoc/blob/master/testData/unexported/main.go#L21>)

```go
func addInternal(num1, num2 Num) Num
```

addInternal is a private version of AddNums.

<a name="Num.Add"></a>
### func \(Num\) [Add](<https://github.com/frgrisk/gomarkdoc/blob/master/testData/unexported/main.go#L11>)

```go
func (n Num) Add(num Num) Num
```

Add adds the other num to this one.

Generated by [gomarkdoc](<https://github.com/frgrisk/gomarkdoc>)
