# 🔄 Go `strconv`

`strconv` ব্যবহার করা হয় এক ধরনের **data type অন্য data type-এ convert করার জন্য**।

## 📚 Resource

* Go `strconv`: https://pkg.go.dev/strconv

---

## 🔢 String → Integer

`Atoi()`

```text
"10" → 10
```

---

## ✅ String → Boolean

`ParseBool()`

```text
"true" → true
```

---

## 🔢 String → Float

`ParseFloat()`

```text
"3.14" → 3.14
```

---

## 📝 Integer → String

`Itoa()`

```text
10 → "10"
```

---

## 📝 Boolean → String

`FormatBool()`

```text
true → "true"
```

---

# 🧠 Easy Way to Remember

```text
String → Other Type

Atoi()       → int
ParseBool()  → bool
ParseFloat() → float
```

```text
Other Type → String

Itoa()       → int → string
FormatBool() → bool → string
```


1 : postgres://USERNAME:PASSWORD@HOST:PORT/DATABASE_NAME?sslmode=disable
2 : https://github.com/jackc/pgx