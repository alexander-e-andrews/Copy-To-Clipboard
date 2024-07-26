# Copy To Clipboard

A Go project that adds an option to the Windows context menu to copy a file to the clipboard. Can copy any utf-8 encoded file.

## Command Line Usage
No Arguments: Add program to the context menu for its current exe location \
-u: Removes program from context menu \
File Path: Attempts to copy the file to the windows clipboard \

## Building
### Inside go bin
1. Use go to install program: `go install github.com/alexander-e-andrews/Copy-To-Clipboard`
2. Run it manually first to add it to context menu: `copytoclipboard`

### Outside go bin
1. `git clone https://github.com/alexander-e-andrews/Copy-To-Clipboard`
2. `go build -ldflags -H=windowsgui`
3. Place the exe wherever you would like it to live, and double click on it to run it. It will add itself to the context menu

## Uninstall
1. Navigate to where the .exe is stored in command line
2. `copytoclipboard.exe -u` <- Removes the program from the context menu
3. Delete the exe
