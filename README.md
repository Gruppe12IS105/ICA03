# IS-105 ICA03

**Det er dette dokumentet som skal være den endelige besvarelsen**. Først kan alle de som skal være med på besvarelsen skrive i google-docs, også legger vi det inn her når vi føler oss ferdige.

I md-filer (markdown) kan vi legge inn kode med syntax-highlighting, også får vi hele besvarelsen i Github i stedet for en pdf med link til Repository-en.

Link til guide om hvordan man skriver markdown-language:
[<u>Markdown](https://guides.github.com/features/mastering-markdown/)

### Pakker å sjekke ut:
- "hex"
- "syscall"
- "io"
- "ioutils"
- "net"
- "strings" (strings.Trim)
- "fmt" - sjekk ut Sprint/Sprintln/Sprintf
- "os"

OS-pakke for å finne ut hvilken språk koden er skrevet i (?) - sjekk ut

- Char: ' '
- String: " "
- Goroutines
- Channels

Eksempel på select:
```go
 select {
   case v := <-ch1:
      fmt.Println("channel 1 sends", v)
   case v := <-ch2:
      fmt.Println("channel 2 sends", v)
   default: //valgfritt
      fmt.Println("Ingen av channelsene sender noe")

 }
```
