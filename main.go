package main

import (
 "fmt"
 "os"
 "log"
 "bytes"
 "io"
 "errors"
 "math/rand"
 "time"
 "bufio"
 "net/http"
 "github.com/spf13/viper"
)

func lineCounter(r io.Reader) (int, error) {
    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
    }
}


func ReadLine(r io.Reader, lineNum int) (line string, lastLine int, err error) {
    sc := bufio.NewScanner(r)
    for sc.Scan() {
        lastLine++
        if lastLine == lineNum {
            return sc.Text(), lastLine, sc.Err()
        }
    }
    return line, lastLine, io.EOF
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func main()  {
  url := "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Fuzzing/User-Agents/user-agents-whatismybrowserdotcom-large.txt"
  //
  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath("$HOME/.rua/")
  //
  err := viper.ReadInConfig()
  if err!=nil {
     if _,err:= os.Stat(os.Getenv("HOME")+"/.rua");
     errors.Is(err,os.ErrNotExist){
       fmt.Println("need to create directory")
       os.Mkdir(os.Getenv("HOME")+"/.rua",0755)
       os.Create(os.Getenv("HOME")+"/.rua/config")
  //
       viper.SetDefault("user_agent_file", os.Getenv("HOME")+"/.rua/"+"user-agents-whatismybrowserdotcom-large.txt")
  //
       viper.SetConfigName("config")
       viper.SetConfigType("yaml")
       viper.AddConfigPath("$HOME/.rua/")
  //
       err :=viper.WriteConfig();
       if err!=nil{
         fmt.Println(err)
       }
     }else{
       fmt.Println("no idea")
     }
  }

  //file,err := os.Open("/home/shoxx/Documents/Codes/Go/UserAgent/user-agents-whatismybrowserdotcom-large.txt")
  if _,err := os.Stat(viper.GetString("user_agent_file"))
  errors.Is(err,os.ErrNotExist){
    DownloadFile(viper.GetString("user_agent_file"),url)
  }


  file,err := os.Open(viper.GetString("user_agent_file"))
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  s1:=rand.NewSource(time.Now().UnixNano())
  r1:=rand.New(s1)

  lineNumber,err := lineCounter(file)
  randomLine := r1.Intn(lineNumber)

  file,err = os.Open(viper.GetString("user_agent_file"))

  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  userAgent,_,err:=ReadLine(file,randomLine)
  fmt.Println(userAgent)
}
