<p align="center" width="100%">
<img src="https://user-images.githubusercontent.com/84747244/231623055-e036ac34-bc89-48ed-b273-d9ac0d7623eb.png">
</p>

# Contents
- <a href="#user-content-overview">Overview</a>
- <a href="#user-content-how-to-use">How to Use</a>

# Overview
Impurgesonator takes a Discord User's ID for the `-u` flag. Then Impurgesonator will check for impersonators in every server where you have invited the Bot.

Impurgesonator will consider a user as an impersonator when:
 - a user's Nickname or Username (case-insensitive) matches the Username of the User specified by the `-u` Discord ID,
 - the user with the case-insensitive match does not have the same Discord ID as the ID provided by `-u`
 
 <br>
Impurgesonator will ban all existing impersonators when initialized. 

It will then listen for live events to detect new impersonators.

<br>

Live events are triggered when:
- a new member joins the Discord server
- an existing user changes their Username or Nickname

<br>
When a live event is triggered, Impurgesonator will check whether the user who triggered the event is an impersonator.

<br>
<br>
<br>

# How to Use

### 1. Clone the repo
```
git clone git@github.com:BoilingSoup/impurgesonator.git
```

### 2. Login to Discord Developer Portal and click New Application

https://discord.com/developers

![image](https://user-images.githubusercontent.com/84747244/231594416-7a9b148a-b9bf-4fdf-9f18-64af52a227f2.png)

Make a name for the application and click Create.

### 3. Create a Bot and get the Bot Token

Click Bot in the sidebar and then Add Bot. Confirm the message and proceed.

![image](https://user-images.githubusercontent.com/84747244/231595004-d4d82472-6aea-4d6e-bc86-5538b390b3c3.png)
![image](https://user-images.githubusercontent.com/84747244/231595598-50e44443-3a73-43c4-83bb-26923697b29f.png)

**Store the Bot Token somewhere safe as it will be used later.**

### 4. Update Bot settings
While in the bot tab, I suggest disabling the "PUBLIC BOT" slider.

And also enable the "PRESENCE INTENT" and "SERVER MEMBERS INTENT" so the bot can get info about members in your server and live updates.
![image](https://user-images.githubusercontent.com/84747244/231602754-d685dbca-8af5-4c7a-8132-8ba51a30efcf.png)

Make sure to save your changes.

### 5. Generate the Bot Invite Link
In the sidebar, click on OAuth2, and then URL Generator.
![image](https://user-images.githubusercontent.com/84747244/231599304-a887c38b-0862-4d88-ac4c-c91dd93d2776.png)

Enable the bot scope
![image](https://user-images.githubusercontent.com/84747244/231599603-9df0df6c-b781-4d7d-9a8c-880df0be0177.png)

Enable the Ban Members bot permission
![image](https://user-images.githubusercontent.com/84747244/231599835-85fb7c25-6615-4c41-bdf1-a6b080d80181.png)

Copy the Generated URL
![image](https://user-images.githubusercontent.com/84747244/231600131-68f66a0d-0a1a-4124-8067-513ca2986f8a.png)


### 6. Invite the bot to your server
Visit the Generated URL you copied in the previous step, and add the bot to a server you own.
![image](https://user-images.githubusercontent.com/84747244/231600360-2350a00b-bcc2-492a-be26-817de0e1d652.png)

When added, your bot should be visible in the sidebar of your server.
![image](https://user-images.githubusercontent.com/84747244/231600727-568b9b13-7c87-4539-9d2d-8fb5ff91bd23.png)


### 7. Get the Discord ID of the User whom you want to prevent impersonators

You can do this by right-clicking on the user in the list of server members.

![image](https://user-images.githubusercontent.com/84747244/231601392-b82962ff-e567-448c-836e-73656ade9193.png)

Store this Discord ID somewhere as it will be used in the next step.

### 8. Start the bot

Now you have the 2 necessary credentials.
- Bot Token from step 3
- Discord ID of who should not be impersonated

Go to the directory where you cloned this repo in step 1.

Run the bot with the following command
```
go run . -t <bot_token> -u <discord_id>
```

The bot will do an initial scan of existing members for impersonators and ban those that meet the criteria.

Then it will keep listening for updates to detect new impersonators.
![image](https://user-images.githubusercontent.com/84747244/231603033-6ba72ff8-51f2-4d95-8c6f-cfd437e8077c.png)

You should also now see that the bot is Online in your server.

![image](https://user-images.githubusercontent.com/84747244/231603342-f0ba3274-0d23-4b7f-9f57-38fefd96ffb1.png)


Done! The bot will ban users that change their Username or Nickname to fit the impersonator criteria. 

And it will also ban new members that join the server and fit the impersonator criteria.
