package main

import (
   "log"
   "github.com/comail/colog"
   "os"
)

func main() {

   logfile, err := os.OpenFile(os.Args[1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
   if err != nil {
      panic("cannnot open test.log:" + err.Error())
   }
   defer logfile.Close()
   
   colog.SetOutput(logfile)
   colog.SetDefaultLevel(colog.LDebug)
   colog.SetMinLevel(colog.LError)
   colog.SetFormatter(&colog.StdFormatter{
      Colors: false,
      Flag:   log.Ldate | log.Ltime | log.Lshortfile,
   })
   colog.Register()

   log.Printf("trace: this is a trace log.")
   log.Printf("debug: this is a debug log.")
   log.Printf("info: this is an info log.")
   log.Printf("warn: this is a warning log.")
   log.Printf("error: this is an error log.")
   log.Printf("alert: this is an alert log.")

   log.Printf("this is a default level log.")
}






// package main

// import (
//    "log"
//    "os"
//    "github.com/comail/colog"
// )

// func main() {
//    cl := colog.NewCoLog(os.Stdout, "worker ", log.LstdFlags)
//    cl.SetMinLevel(colog.LInfo)
//    cl.SetDefaultLevel(colog.LWarning)
//    cl.FixedValue("worker_id", 42)

//    logger := cl.NewLogger()
//    logger.Print("this gets warning level")
//    logger.Print("debug: this won't be displayed")
// }




// package main

// import (
//    "log"
//    "os"
//    "time"

//    "github.com/comail/colog"
// )

// func main() {
//    file, err := os.OpenFile("temp_json.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
//    if err != nil {
//       panic(err)
//    }

//    colog.Register()
//    colog.SetOutput(file)
//    colog.ParseFields(true)
//    colog.SetFormatter(&colog.JSONFormatter{
//       TimeFormat: time.RFC3339,
//       Flag:       log.Lshortfile,
//    })

//    log.Print("info: logging this to json")
//    log.Print("warning: with fields foo=bar")
// }





// package main

// import (
// 	"log"
// 	"github.com/comail/colog"
//    "os"
// )

// func main() {

//    logfile, err := os.OpenFile(os.Args[1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
//    if err != nil {
//       panic("cannnot open test.log:" + err.Error())
//    }
//    defer logfile.Close()
   
//    colog.SetOutput(logfile)
// 	colog.SetDefaultLevel(colog.LDebug)
// 	colog.SetMinLevel(colog.LTrace)
// 	colog.SetFormatter(&colog.StdFormatter{
// 		Colors: true,
// 		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
// 	})
// 	colog.Register()

// 	log.Printf("trace: this is a trace log.")
// 	log.Printf("debug: this is a debug log.")
// 	log.Printf("info: this is an info log.")
// 	log.Printf("warn: this is a warning log.")
// 	log.Printf("error: this is an error log.")
// 	log.Printf("alert: this is an alert log.")

// 	log.Printf("this is a default level log.")
// }
