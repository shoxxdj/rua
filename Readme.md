#Rua

A small binary to generate a random user agent string.
No more excuses for using a default tool useragent (Hello to you sqlmap useragent)
Based on SecList User agent file :

```
https://raw.githubusercontent.com/danielmiessler/SecLists/master/Fuzzing/User-Agents/user-agents-whatismybrowserdotcom-large.txt
```

## Setup

```
go install github.com/shoxxdj/rua@latest
```

## Usage

```
rua
```

## Integration in any tools :

```
whateverbinary -whateveroption $(rua)
```
