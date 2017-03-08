# re.Do

```
var str string
err := re.Do(func(i int) (bool, error) {
  var err error
  str, err = YourFunc()
  return i < 2, err // 2 attempts
})
```
