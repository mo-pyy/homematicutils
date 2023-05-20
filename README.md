# homematicutils

## Quickstart

```go
import "github.com/mo-pyy/homematicutils"

func main() {
  var homematic = &homematicutils.HomematicInfo{
      Hostname: os.Args[1],
      User:     os.Args[2],
      Password: os.Args[3],
  }
  client := &http.Client{Timeout: time.Second * 10}
  var value int = 0
  homematicutils.SetIntVar(*client, *homematic, "test_wert", value)
}
```
