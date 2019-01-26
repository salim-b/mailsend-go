// DO NOT MODIFY - Automatically generated from ./files/examples.txt

package main


import "fmt"

const (
    examples = `



EXAMPLES


Each example mailsend-go command is a single line. In Unix back slash  
can be used to continue in the next line. Also in Unix, use single
quotes instead of double quotes, otherwise if input has any shell
character like $ etc, it will get expanded by the shell.


Show SMTP server information

StartTLS will be used if server supports it

      mailsend-go -info -smtp smtp.gmail.com -port 587

    [S] 220 smtp.gmail.com ESMTP k185-v6sm17739711qkd.27 - gsmtp
    [C] HELO localhost
    [C] EHLO localhost
    [S] 250-smtp.gmail.com at your service, [x.x.x.x]
    [S] 250-SIZE 35882577
    [S] 250-8BITMIME
    [S] 250-STARTTLS
    [S] 250-ENHANCEDSTATUSCODES
    [S] 250-PIPELINING
    [S] 250-CHUNKING
    [S] 250-SMTPUTF8
    [C] STARTTLS
    [S] 220-2.0.0 Ready to start TLS
    [C] EHLO localhost
    [S] 250-smtp.gmail.com at your service, [x.x.x.x]
    [S] 250-SIZE 35882577
    [S] 250-8BITMIME
    [S] 250-AUTH LOGIN PLAIN XOAUTH2 PLAIN-CLIENTTOKEN OAUTHBEARER XOAUTH
    [S] 250-ENHANCEDSTATUSCODES
    [S] 250-PIPELINING
    [S] 250-CHUNKING
    [S] 250-SMTPUTF8
    Certificate of smtp.gmail.com:
     Version: 3 (0x3)
     Serial Number: 149685795415515161014990164765 (0x1e3a9301cfc7206383f9a531d)
     Signature Algorithm: SHA256-RSA
     Subject: CN=Google Internet Authority G3,O=Google Trust Services,C=US
     Issuer: GlobalSign
     Not before: 2017-06-15 00:00:42 +0000 UTC
     Not after: 2021-12-15 00:00:42 +0000 UTC
    [C] QUIT
    [S] 221-2.0.0 closing connection k185-v6sm17739711qkd.27 - gsmtp

Use SSL. Note the port is different

      mailsend-go -info -smtp smtp.gmail.com -port 465 -ssl


Send mail with a text message

Notice “auth” is a command and it takes -user and -pass arguments.
“body” is also a command and here it took -msg as an argument. The
command “body” can not repeat, if specified more than once, the last one
will be used.

        mailsend-go -sub "Test"  -smtp smtp.gmail.com -port 587 \
         auth \
          -user jsnow@gmail.com -pass "secret" \
         -from "jsnow@gmail.com" -to  "mjane@example.com" \
         body \
           -msg "hello, world!"

The environament variable “SMTP_USER_PASS” can be used instead of the
flag -pass.


Send mail with a HTML message

        mailsend-go -sub "Test"  \
        -smtp smtp.gmail.com -port 587 \
        auth \
         -user jsnow@gmail.com -pass "secret" \
        -from "jsnow@gmail.com"  \
        -to  "mjane@example.com" -from "jsnow@gmail.com" \
        body \
         -msg "<b>hello, world!</b>"


Attach a PDF file

MIME type will be detected. Content-Disposition will be set to
“attchment”, Content-Transfer-Encoding will be “Base64”. Notice,
“attach” is a command it took -file as an arg. The command “attach” can
repeat.

        mailsend-go -sub "Test"  \
        -smtp smtp.gmail.com -port 587 \
        auth \
         -user jsnow@gmail.com -pass "secret" \
        -from "jsnow@gmail.com"  \
        -to  "mjane@example.com" -from "jsnow@gmail.com" \
        body \
         -msg "A PDF file is attached" \
        attach \
         -file "/path/file.pdf"


Attach a PDF file and an image

Notice, the “attach” command is repeated here.

        mailsend-go -sub "Test"  \
        -smtp smtp.gmail.com -port 587 \
        auth \
         -user jsnow@gmail.com -pass "secret" \
        -from "jsnow@gmail.com"  \
        -to  "mjane@example.com" -from "jsnow@gmail.com" \
        body \
         -msg "A PDF file and a PNG file is attached" \
        attach \
         -file "/path/file.pdf" \
        attach \
         -file "/path/file.png"


Attach a PDF file and embed an image

Content-Dispositon for the image will be set to “inline”. It’s an hint
to the mail redear to display the image on the page. Note: it is just a
hint, it is up to the mail reader to respect it or ignore it.

        mailsend-go -sub "Test"  \
        -smtp smtp.gmail.com -port 587 \
        auth \
         -user jsnow@gmail.com -pass "secret" \
        -from "jsnow@gmail.com"  \
        -to  "mjane@example.com" -from "jsnow@gmail.com" \
        body \
         -msg "A PDF file is attached, image should be displayed inline" \
        attach \
         -file "/path/file.pdf" \
        attach \
         -file "/path/file.png" \
         -inline


Set Sender and Recipient’s name

        mailsend-go -sub "Testing -fname and -tname"  \
        -smtp smtp.gmail.com -port 587 \
        auth \
         -user example@gmail.com -pass "secret" \
         -to jsoe@example.com \
         -tname "John Soe" \
         -fname "Example Foo" \
         -f "example@gmail.com" \
         body -msg "Testing Recipient and Sender's name"


Set Carbon Copy and Blind Carbon copy

        mailsend-go -sub "Testing -cc and -bcc" \
        -smtp smtp.gmail.com -port 587 \
        auth \
         -user example@gmail.com -pass "secret" \
         -to jsoe@example.com \
         -f "example@gmail.com" \
         -cc "user1@example.com,user2@example.com" \
         -bcc "foo@example.com" \
         body -msg "Testing Carbon Copy and Blind Carbon copy"

Cc addresses will be visible to the recipients but Bcc address will not
be.


Send mail to a list of users

Create a file with list of users. The syntax is Name,email_address in a
line. Name can be empty but comma must be specified. Example of a list
file:

    John Snow,jsnow@example.com
    Mary Jane,mjane@example.com
    ,foobar@example.com

Specify the list file with -list flag.

        mailsend-go -sub "Test sending mail to a list of users" \
        -smtp smtp.gmail.com -port 587 \
        auth \
         -user example@gmail.com -pass "secret" \
            -f "me@example.com" \
            -to "xyz@example.com" \
            body \
            -msg "This is a test of sendmail mail to a list of users" \
            attach \
                -file "cat.jpg" \
             attach \
                -file "flower.jpg" \
                -inline \
             -list "list.txt"


Add Custom Headers

Use the command “header” to add custom headers. The command “header” can
be repeated.

        mailsend-go -sub "Testing custom headers" \
        -smtp smtp.gmail.com -port 587 \
        auth \
         -user example@gmail.com -pass "secret" \
         -to jsoe@example.com \
         -f "example@gmail.com" \
         body -msg "Testing adding Custom headers"
         header \
             -name "X-MyHeader-1" -value "Value of X-MyHeader-1" \
         header \
             -name "X-MyHeader-2" -value "Value of X-MyHeader-2"

------------------------------------------------------------------------

(Generated from docs/examples.md)

------------------------------------------------------------------------
`
)

// Print Examples ...
func PrintExamples() {
    fmt.Println(examples)
}
