package credentials

import "github.com/joho/godotenv"
import "log"

func Load() {
    if err := godotenv.Load(); err != nil {
        log.Println("error loading env file !")
        log.Fatal(err)
    }
}
