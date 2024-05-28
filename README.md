### Description

This program allows you to keep the status online 24/7 in Discord

### Build
`go build -o discrod-online main.go`

### Run
Set up variable DISCORD_USER_TOKEN in environment and run program:
`DISCORD_USER_TOKEN=your_token ./discrod-online`
Or set variable in your OS

### How get discord token

1. Logging in to your discord account in web browser
2. Pressing Ctrl+Shift+I to open Chrome Developer Tools
3. Go to the Network Tab
4. Keep it open and refresh the page
5. Type /api in the filter search box
6. Click the entry that has science as the Name
7. On the sub-menu, go to Headers
8. Scroll down till you see an entry named Authorization. Copy the line next to it.
9. This is your token, DO NOT GIVE IT TO ANYONE.
