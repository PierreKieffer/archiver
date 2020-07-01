# archiver 

Basic implementation of zip file archiver 

## Download
```bash 
go get github.com/PierreKieffer/archiver
```
## Usage 

params : 
- input : Path of directory to compress
- output : archive file path (.zip extension)

```go
import(
	"github.com/PierreKieffer/archiver"
)
func main() {
        err := archiver.ZipArchiver("/home/user/dir-to-compress", "/home/user/archive-name.zip")
        if err != nil {
                log.Println(err)
        }
}
```
