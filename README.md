# Simple Mail T(racker) Protocol 🎣 📨

<div align=center>
<img src= https://github.com/ariary/SMTrackerP/blob/main/img/A7C1B863-E3CF-4B30-A60C-6C67B6FBB463.png width=180>

<br><strong><i>Ensure your email has been read</i></strong>

👁️ <strong>•</strong> 🔨 <strong>•</strong> 📬
</div> 


***Why?***
* Because many SMTP clients don't provide "acknowledgement of receipt" anymore
* In phishing campaign (with the prior agreements of course!), to perform information gathering. (retrieve information about who read, with which browser, at what time etc..)
* For commercial mail, to see who has been attracted by your email, but it is definitely not cool!


*Current state is a very naive implementation, may be enhancements will be made espacially for the phishing use case (deal with multiple targets, directly send the email providing a mail list, etc..)*

## 🚀 Launch instructions

**1️ Launch `smtrackerp`**

```shell
cd $(mktemp -d) # preferably in temporary dir
smtrackerp -e [your_external_url] [target_mail]
```

**2️ Insert payload in your mail**

`smtrackerp` will generate a html file containing an invisible image. Insert it in your mail.

Send it!

**3️ Wait..**

![demo](https://github.com/ariary/SMTrackerP/blob/main/img/demo.png)

<sup> I know, i know .. i'm using `ngrok`, for testing purpose only</sup>
