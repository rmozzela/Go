package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strconv"
    "strings"

)

type Song struct {
    Title    string
    Filename string
    Seconds  int
}

type SongPls struct {
    Ext    string
    Filename string
    Seconds string    //need to change back to int after debugging
}

func main() {
    if len(os.Args) == 1 || (!strings.HasSuffix(os.Args[1], ".pls") && !strings.HasSuffix(os.Args[1], ".m3u")) {
        fmt.Printf("usage: %s <file.[pls|m3u]>\n", filepath.Base(os.Args[0]))
        os.Exit(1)
    }

    if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
        log.Fatal(err)

    } else if strings.HasSuffix(os.Args[1], ".m3u") {   //read m3u file type
        songs := readM3uPlaylist(string(rawBytes))     //parse the m3u file type
        writePlsPlaylist(songs)                        //writes to output with appending File, Type, Length
    } else {
        songs := readPlsPlaylist(string(rawBytes))     // read pls file type
        writeM3uPlaylist(songs)
    }
}

//parsing m3u file type. returns a Slice of Song
func readM3uPlaylist(data string) (songs []Song) {
    var song Song
    for _, line := range strings.Split(data, "\n") {
        line = strings.TrimSpace(line)
        if line == "" || strings.HasPrefix(line, "#EXTM3U") {
            continue
        }
        if strings.HasPrefix(line, "#EXTINF:") {  //parse the EXTINF line
            song.Title, song.Seconds = parseExtinfLine(line) //set the title and seconds to struct
        } else {
            song.Filename = strings.Map(mapPlatformDirSeparator, line) //return the second line Filename
        }
        if song.Filename != "" && song.Title != "" && song.Seconds != 0 { //validate contents of songs three fields.
            songs = append(songs, song) //set contents to array
            song = Song{}  //set current song back to zero
        }
    }
    return songs
}

//parsing pls file type. returns a Slice of Song
func readPlsPlaylist(data string) (songs []SongPls) {
    var song SongPls
    for _, line := range strings.Split(data, "\n") {
        line = strings.TrimSpace(line)
        if line == "" || strings.HasPrefix(line, "[playlist]") {
            continue
        }

        if strings.HasPrefix(line, "Length") {
           song.Seconds = parseFileinfLine(line)
          }

        if strings.HasPrefix(line, "File") {  //parse the EXTINF line
          song.Filename = parseFileinfLine(line)
       }
       if strings.HasPrefix(line, "Title") {
            song.Ext = parseFileinfLine(line)
        }
        if song.Ext != "" && song.Filename != "" && song.Seconds !="" { //validate contents of songs three fields.
            songs = append(songs, song) //set contents to array
            song = SongPls{}  //set current song back to zero
        }
    }
    return songs
}

//Parsing M3U File type
func parseExtinfLine(line string) (title string, seconds int) {
    if i := strings.IndexAny(line, "-0123456789"); i > -1 {
        const separator = ","
        line = line[i:]
        if j := strings.Index(line, separator); j > -1 {
            title = line[j+len(separator):]
            var err error
            if seconds, err = strconv.Atoi(line[:j]); err != nil {
                log.Printf("failed to read the duration for '%s': %v\n",
                    title, err)
                seconds = -1
            }
        }
    }
    return title, seconds
}
//Parsing PLS File type
func parseFileinfLine(line string) (line_parse string) {
    if i := strings.IndexAny(line, "-0123456789"); i > -1 {
        const separator = "="
        line = line[i:]
        if j := strings.Index(line, separator); j > -1 {
            line_parse = line[j+len(separator):]
        }
    }
   return line_parse
}

func mapPlatformDirSeparator(char rune) rune {
    if char == '/' || char == '\\' {  //linux or windows file directory type
        return filepath.Separator
    }
    return char
}

func writeM3uPlaylist(songs []SongPls) {
    fmt.Println("#EXTM3U")

    for i, song := range songs {
        i++
      fmt.Printf("EXTINF:%s,%s\n", song.Seconds, song.Ext)
      fmt.Printf("%s\n",song.Filename)
    }
}
func writePlsPlaylist(songs []Song) {
    fmt.Println("[playlist]")
    for i, song := range songs {
        i++
        fmt.Printf("File%d=%s\n", i, song.Filename)
        fmt.Printf("Title%d=%s\n", i, song.Title)
        fmt.Printf("Length%d=%d\n", i, song.Seconds)
    }
    fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}
