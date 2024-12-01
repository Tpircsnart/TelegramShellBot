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
* Settings
    ```json
        {
            "token": "", // your telegram bot token
            "shell_location": "./commands/", // the directory where you put .sh files, default: ./commands
            "row_button_count": 3, // the maximum number of row buttons to display
            "isDebug": true // whether to display debug information
        }
    ```
* Rule
    * the shell filename is the same as the label on the button
* Limits
    * the length of result message is limited to below **4096** characters by Telegram Policy