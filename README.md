# re.Do

```
var str string
// 2 tries
err := re.Do(2, func() error {
  var e error
  str, e = YourFunc()
  return e
})
```
