# Advanced Algorithms ![Polska](https://img.shields.io/badge/Language-Polish-red?style=flat&logo=flag-icon&logoColor=white)![English](https://img.shields.io/badge/Language-English-blue?style=flat&logo=flag-icon&logoColor=white)

# O Aplikacji ![Polska](https://img.shields.io/badge/Language-Polish-red?style=flat&logo=flag-icon&logoColor=white)

## Aplikacja składa się z **pięciu** projektów

> Liczby czworacze

> Wyznacznik macierzy do dziesiątego stopnia

> Metoda bisekcji dla wielomianu piątego stopnia

> Algorytm Monte Carlo

## Liczby czworacze

Projekt wyszukujący w podanym przedziale liczb czworaczych, czyli takich, które są liczbami pierwszymi i są postaci `p, p+2, p+6, p+8`.

- `GET /quadruplet` - Zwraca informacje o występujących liczbach czworaczych z danego przedziału

### Request body

```
{
    "start" - int,
    "end" - int
}
```
### Response body
```
{
    "PrimeCount" - int,
    "PairsCount" - int,
    "QuadrupletCount" - int,
    "Quadrupltes" - List<int[4]>
}
```

## Wyznacznik Macierzy

Projekt wyliczający wyznacznik macierzy od pierwszego do dziesiątego stopnia.

### Request body
```
{
    "size" - int
    "matrix" - decimal[size][size]
}
```

### Response body
```
{
   "determinant" - decimal
}
```

## Metoda Bisekcji

Metoda wyszukiwania x`ów, dla których funkcja piątego stopnia przyjmuje wartość równą 0

### Request body
```
{
    "numbers" - decimal[] // Liczby ustawione od x^0 do x^5,
    "delta" - decimal // Mniejsza niż zakres
    "min" - decimal
    "max" - decimal
}
```

### Respone body
```
{
    "number" decimal[]
}
```