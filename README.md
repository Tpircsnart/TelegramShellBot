# Telegram Shell Bot
* Description
    * A Golang based Telegram bot which get the result you made as shell files
* How to use
    * Must change values under key:**token** in `settings.json`
    * Edit the shell files(default in commands)
    * Run
        * straight by terminal command
            ```sh
            go run main.go
            ```
        * build and run
            ```sh
            go build && ./telegram-shell-bot 
            ```
* Bot Commands
    * **open**
        * to open the command keyboard
    * **close**
        * to close the command keyboard
    * **reset**
        * to re-read files from `shell_location` in `./env/settings.json`
* Limits
    * the length of result message is limited to below **4096** characters by Telegram Policy