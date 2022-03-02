#Rua

A small binary to generate a random user agent string.
Based on SecList User agent file :

```
https://raw.githubusercontent.com/danielmiessler/SecLists/master/Fuzzing/User-Agents/user-agents-whatismybrowserdotcom-large.txt
```

## Setup

```
go install github.com/shoxxdj/rua
```

## Usage

```
rua
```

## Integration in any tools :

```
whateverbinary -whateveroption $(rua)
```
