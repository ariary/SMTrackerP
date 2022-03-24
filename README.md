# Simple Mail T(racker) Protocol 

<div align=center>
<img src= https://github.com/ariary/SMTrackerP/blob/main/img/logo.png width=180>

<br><strong><i>Ensure your email has been read</i></strong>

👁️ <strong>•</strong> 🎣 <strong>•</strong> 📬
</div> 


***Why?***
* In phishing campaign (with the prior agreements of course!), to perform information gathering. (retrieve information about who read, with which browser, at what time etc..)
* Because many SMTP clients don't provide "acknowledgement of receipt" anymore
* For commercial mail, to see who has been attracted by your email, but it is definitely not cool!

*Current state is a very naive implementation, may be enhancements will be made espacially for the phishing use case (deal with multiple targets, directly send the email providing a mail list, etc..)*

## 🚀 Launch instructions

**1️ Launch `smtrackerp`**

```shell
smtrackerp --url [YOUR_EXTERNAL_URL] -t [TARGET_MAIL]
```

**2️ Insert payload in your mail**

`smtrackerp` will generate a html file containing an invisible image. Insert it in your mail.

Then send it!

**3️ Wait..**

![demo](https://github.com/ariary/SMTrackerP/blob/main/img/demo.png)

<sup> I know, i know .. i'm using `ngrok`, for testing purpose only</sup>

### 📨 Send & track in one-step

You can use `smtrackerp` to directly send your mail and then waiting for the "proof of reading".

**1️ Put the mail body in a file**

*You could use `test/test.html` if you want*

**2️ Load your SMTP credential**

Via environment variable, the credentials are used to communicate with your SMTP server.
I encourage you to save it in a file:
```shell
export SMPT_USERNAME=[YOUR_MAIL]
export SMPT_PASSWORD=[YOUR_PASSWORD]
```
and source it

**3️ Send and wait**

```shell
smtrackerp send -u [YOUR_EXTERNAL_URL] -r [LIST_OF_RECIPIENTS] -s [MAIL_SUBJECT] -b [BODY_FILENAME] -a [SMP_HOST] -p [SMTP_PORT]  --track
```

#### Demo
What the target sees:

![target](https://github.com/ariary/SMTrackerP/blob/main/img/mail.png)

What we see:

![demo-send](https://github.com/ariary/SMTrackerP/blob/main/img/send-demo.png)
