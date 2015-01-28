package main

import(
  "github.com/codegangsta/cli"
  "github.com/rakyll/magicmime"
  "os"
  "fmt"
  "mime"
  "strings"
  "path/filepath"
)

type Mismatch struct {
  File string
  ByMime string
  ByExt string
}

func main(){
  app := cli.NewApp()
  app.Name = "mimedoc"
  app.Usage = "Cross reference extension-based and content-based mimetypes."
  app.Email = "jondotan@gmail.com"
  app.Author = "Dotan Nahum"
  app.Version = "0.0.1"

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "path, p",
      Value: ".",
      Usage: "destination path",
    },
    cli.BoolFlag{
      Name: "report, r",
      Usage: "generate a live report",
    },
    cli.StringSliceFlag{
      Name: "ext, e",
      Value: &cli.StringSlice{},
      Usage: "Pick a specific extension",
    },
  }

  app.Action = func(c *cli.Context) {
    mm, err := magicmime.New(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR)
    if err != nil {
        panic(err)
    }

    totalUnknownExt := 0
    totalUnknownMime := 0
    totalFiles := 0
    totalChecked := 0

    whiteListedExts := map[string]bool{}
    for _,e := range c.StringSlice("ext"){
      whiteListedExts[e] = true
    }

    mismatches := []Mismatch{}


    //                              gitignore.Visit(func.... ) -> func // gitignore.ShouldIgnore(path)
    filepath.Walk(c.String("path"), func(s string, fi os.FileInfo, err error) error {
      if fi.IsDir() {
        return nil
      }
      totalFiles++

      ext := filepath.Ext(s)
      if len(whiteListedExts) > 0 && !whiteListedExts[ext] {
        return nil
      }

      extType := strings.Split(mime.TypeByExtension(ext), ";")[0]
      if extType == "" {
        totalUnknownExt++
        return nil
      }

      mimetype, err := mm.TypeByFile(s)
      if err != nil{
        totalUnknownMime++
        return nil
      }

      if extType != mimetype{
        mismatches = append(mismatches, Mismatch{ s, mimetype, extType })
      }

      totalChecked++
      return nil
    })


    if c.Bool("report") {
      if len(mismatches) > 0 {
        fmt.Printf("FILE\tMIME\tEXT\n")
        for _, mismatch := range mismatches {
          fmt.Printf("%s\t%s\t%s\n", mismatch.File, mismatch.ByMime, mismatch.ByExt)
        }
      }
      fmt.Printf("%d/%d checked (%.2f%%), %d mismatch, %d unknown by mime, %d unknown by ext\n", totalChecked, totalFiles, float32(totalChecked)*100/float32(totalFiles), len(mismatches), totalUnknownMime, totalUnknownExt)
    }

    if len(mismatches) > 0 {
      os.Exit(1)
    }
  }
  app.Run(os.Args)

}

