# vCard-Generator
Generate mockup vCard file for mobile phone



## How to use
Run command

`go run MockContact.go`

Then, Enter the number of contacts that you want.
It will generate mockup vCard into json file.



## Data format
```
BEGIN:VCARD
VERSION:3.0
FN:TestFirstName
EMAIL;TYPE=HOME,INTERNET,pref:TestFirstName@testmail.com
END:VCARD
```