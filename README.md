# alif-resource

 Do not communicate by sharing memory
 
 Don’t use floating point numbers for money.
 
http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/

Capturing with filters (alway use ssl https) man in middle
https://letsencrypt.org/

About User Enumeration
Login

    Make sure to return a generic “No such username or password” message when a login failure occurs.
    Make sure the HTTP response, and the time taken to respond are no different when a username does not exist, and an incorrect password is entered.

Password Reset

    Make sure your “forgotten password” page does not reveal usernames.
    If your password reset process involves sending an email, have the user enter their email address. Then send an email with a password reset link if the account exists.

Registration

    Avoid having your site tell people that a supplied username is already taken.
    If your usernames are email addresses, send a password reset email if a user tries to sign-up with an existing address.
    If usernames are not email addresses, protect your sign-up page with a CAPTCHA.

Profile Pages

    If your users have profile pages, make sure they are only visible to other users who are already logged in.
    If you hide a profile page, ensure a hidden profile is indistinguishable from a non-existent profile.
    
Similar to the login page, if you have user profile pages, be careful about allowing username enumeration. For example, if someone visits /users/JohnDoe and then /users/JaneDoe, and one returns a 404 Not Found error, while the other returns an 401 Access Denied error, the attacker can infer that one account actually exists and the other does not.

Preventing LFI and RFI abuse

http://localhost/displayFile?filename=../../../etc/passwd   

http://localhost/displayFile?filename=/etc/passwd  

* Validate all inputs

* Finding unlisted files on a web server

* Fuzzing a network service

* Base64 encoding data

https://golanglibs.com/

https://www.owasp.org/index.php/Top_10-2017_Top_10

