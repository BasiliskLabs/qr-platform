# qr-platform
A simple QR platform in GoLang to create "dynamic" QR codes.

This is by no means a ground breaking implementation. Over time I would like to add some analytics support possibly, but for now I am happy with it working for free.

Feel free to submit pull requests or any issues!

# Set up
You will need to create a 'redirects.json' file (or change the name of the read-in file).

A simple example of this json file : 

```JSON
[
  {
    "Name": "Joe Shmo",
    "Code": 1,
    "Redirect": "https://www.google.com"
  },
  {
    "Name": "Joe Shmo",
    "Code": 2,
    "Redirect": "https://www.gmail.com"
  }
]
```

# QR codes
In order to properly redirect with QR codes they need to have a similar URL to www.mydomain.com/?name=...&code=...
