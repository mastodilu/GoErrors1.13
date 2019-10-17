# Errors in Go 1.13

[link](https://golang.org/pkg/errors/)

I nuovi metodi aggiunti con questo aggiornamento sono:

package `errors`

- `func New(text string) error`
- `func Is(err, target error) bool`
- `func As(err error, target interface{}) bool`
- `func Unwrap(err error) error`

## errors.New

`func New(text string) error`

Crea un nuovo errore con il testo passato.

Per definire un errore che fa da **wrapper** per un altro errore, ad esempio per fornire più dettagli, si usa `%w` con `fmt.Errorf("...%w...", err)`.

```Go
// new error
err := errors.New("errore 1")
// new error with wrap
err2 := fmt.Errorf("errore 2 %w", err)
err3 := fmt.Errorf("errore 3 %w", err2)
```

```Go
fmt.Println("err")
fmt.Printf("%v\n\n", err)

fmt.Println("err2")
fmt.Printf("%v\n\n", err2)

fmt.Println("err3")
fmt.Printf("%v\n\n", err3)
```

```cmd
err
errore 1

err2
errore 2 errore 1

err3
errore 3 errore 2 errore 1
```

## errors.Is

`func Is(err, target error) bool`

Restituisce true se `target` appare nella catena di errori di `err` (la catena di errori è ottenuta con il verbo `%w`).

```Go
fmt.Println("errors.Is(err3, err)")
bo := errors.Is(err3, err) // true
fmt.Printf("%t\nerr3: %v\nerr: %v\n\n", bo, err3, err)

fmt.Println("errors.Is(err3, err2)")
bo = errors.Is(err3, err2) // true
fmt.Printf("%t\nerr3: %v\nerr2: %v\n\n", bo, err3, err2)

fmt.Println("errors.Is(err, err3)")
bo = errors.Is(err, err3) // false
fmt.Printf("%t\nerr: %v\nerr3: %v\n\n", bo, err, err3)
```

```cmd
errors.Is(err3, err)
true
err3: errore 3 errore 2 errore 1
err: errore 1

errors.Is(err3, err2)
true
err3: errore 3 errore 2 errore 1
err2: errore 2 errore 1

errors.Is(err, err3)
false
err: errore 1
err3: errore 3 errore 2 errore 1
```

## errors.As

`func As(err error, target interface{}) bool`

Restituisce true al primo errore nella catena di errori di `err` che matcha l'**errore puntato** da `target` (target è un indirizzo!).

---

```Go
package main

import (
    "fmt"
    "errors"
)

func main() {
    err := errors.New("errore 1")

    // wrap con %w
    err2 := fmt.Errorf("errore 2 %w", err)
    err3 := fmt.Errorf("errore 3 %w", err2)

    fmt.Println("err")
    fmt.Printf("%v\n\n", err)

    fmt.Println("err2")
    fmt.Printf("%v\n\n", err2)

    fmt.Println("err3")
    fmt.Printf("%v\n\n", err3)

    fmt.Println("errors.Unwrap(err3)")
    fmt.Printf("%v\n\n", errors.Unwrap(err3))

    fmt.Println("errors.Is(err3, err)")
    bo := errors.Is(err3, err)
    fmt.Printf("%t\nerr3: %v\nerr: %v\n\n", bo, err3, err)

    fmt.Println("errors.Is(err3, err2)")
    bo = errors.Is(err3, err2)
    fmt.Printf("%t\nerr3: %v\nerr2: %v\n\n", bo, err3, err2)

    fmt.Println("errors.Is(err, err3)")
    bo = errors.Is(err, err3)
    fmt.Printf("%t\nerr: %v\nerr3: %v\n\n", bo, err, err3)

    fmt.Println("errors.As(err2, &err3)")
    bo = errors.As(err2, &err3)
    fmt.Printf("%t\nerr3: %v\nerr2: %v\n\n", bo, err3, err2)

    fmt.Println("errors.As(err3, &err2)")
    bo = errors.As(err3, &err2)
    fmt.Printf("%t\nerr3: %v\nerr2: %v\n\n", bo, err3, err2)
}

```

```cmd
err
errore 1

err2
errore 2 errore 1

err3
errore 3 errore 2 errore 1

errors.Unwrap(err3)
errore 2 errore 1

errors.Is(err3, err)
true
err3: errore 3 errore 2 errore 1
err: errore 1

errors.Is(err3, err2)
true
err3: errore 3 errore 2 errore 1
err2: errore 2 errore 1

errors.Is(err, err3)
false
err: errore 1
err3: errore 3 errore 2 errore 1

errors.As(err2, &err3)
true
err3: errore 2 errore 1
err2: errore 2 errore 1

errors.As(err3, &err2)
true
err3: errore 2 errore 1
err2: errore 2 errore 1
```
