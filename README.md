# Discord Online Status Keeper

This program allows you to keep your status online 24/7 in Discord.

## Another Golang Version

You can find another project here: [ProjectDGT by Milou4Dev](https://github.com/Milou4Dev/ProjectDGT)

## Description

This program ensures that your Discord status remains online continuously.

## Build

To build the project, run the following command:

`go build -o discord-online main.go`

## Run

Set up the `DISCORD_USER_TOKEN` environment variable and run the program:

`DISCORD_USER_TOKEN=your_token ./discord-online`

Alternatively, you can set the `DISCORD_USER_TOKEN` variable in your operating system's environment settings.

## How to Get Your Discord Token

1. Log in to your Discord account using a web browser.
2. Press `Ctrl+Shift+I` to open Chrome Developer Tools.
3. Go to the **Network** tab.
4. Keep the Developer Tools open and refresh the page.
5. Type `/api` in the filter search box.
6. Click the entry that has `science` as the Name.
7. In the sub-menu, go to **Headers**.
8. Scroll down until you see an entry named `Authorization`. Copy the line next to it.
9. This is your token. **DO NOT GIVE IT TO ANYONE.**

## Important Note

Your Discord token is sensitive information. Keep it secure and do not share it with anyone. Misuse of your token can
lead to unauthorized access to your Discord account.