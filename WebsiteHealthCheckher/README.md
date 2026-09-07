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


# Go + PostgreSQL Resources

## 1. PostgreSQL Connection String

```text
postgres://USERNAME:PASSWORD@HOST:PORT/DATABASE_NAME?sslmode=disable
```

**কাজ:** Go application-কে PostgreSQL database-এর সাথে connect করে।

---

## 2. pgx

https://github.com/jackc/pgx

**কাজ:** Go থেকে PostgreSQL database-এর সাথে connect এবং query করার জন্য ব্যবহার করা হয়।

```bash
go get github.com/jackc/pgx/v5
```

---

## 3. godotenv

https://github.com/joho/godotenv

**কাজ:** `.env` file থেকে secret এবং environment variable load করার জন্য ব্যবহার করা হয়।

```bash
go get github.com/joho/godotenv
```

---

## Flow

```text
.env
  ↓
godotenv
  ↓
DATABASE_URL
  ↓
pgx
  ↓
PostgreSQL
```
