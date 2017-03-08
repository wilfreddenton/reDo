# re.Do

```
var str string
// 2 tries
err := re.Do(2, func(i int) error {
  var err error
  str, err = YourFunc()
  return err
})
```
