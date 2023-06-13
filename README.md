# chatgpt 
cmdline tool for chatgpt api, get chatgpt response from your cmdline

# pre-requisite
```
1. create API Key from your chatGPT account using https://platform.openai.com/account/api-keys 
2. create .chatgpt file under home directory of user ( example: /home/username/.chatgpt ) and specify api_key in key-value format

sample:
cat /home/username/.chatgpt 
API_KEY=<your_api_key>
```

# install chatgpt cmdline
```
go install github.com/pavansh/chatgpt
```

# usage
```
# chatgpt "your question"
```