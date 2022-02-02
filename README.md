# ScaleoAcriveBot_GetLink

### Description:

>This script generates a link to log in to your personal dashboard using the scaleo service API. The script processes requests from users in the telegram bot, if the user is in the users.json database - it gets the login link in the reply message. The link can be obtained only during working hours and only for existing users. The .env contains a link with the API key and the telegram bot token. users.json contains the email and pass from scale and the personal chat_id from telegram''

>[a similar project without telegram, just script](https://github.com/sarff/ScaleoTempLink)

1) rename .env_example to .env
2) rename users.json_example to users.json
3) Fill these files with the correct information (users.json and .env)

example how to run:
1) nohup /root/ScaleoTempLink/ScaleoActiveBot_GetLink&
