## EMAIL-VERIFIER
<hr>

## About

This is simple email domain verifier program. Which checks the MX, SPF and DMARC records.
> <a href="https://www.cloudflare.com/learning/dns/dns-records/dns-mx-record/">learn</a> about MX 
<br><a href="https://kinsta.com/knowledgebase/spf-record/">learn</a> about SPF
<br><a href="https://www.cloudflare.com/learning/dns/dns-records/dns-dmarc-record/">learn</a> about DMARC

> Only <b>VerifyEmailDomain()</b> function is exposed, rest are kept private to show different types of access level in GO.

> All packages that are used, are inbuild of GO.

> Common naming conventions are followed for creating the packages.

## Build and Run
<ol>
    <li>type <i>go mod init github.com/anupam/email-verifier</i>. It will create the .mod file.</li>
    <li>type <i>go build</i>. It will create .exe file.</li>
    <li>type <i>go run main.go</i> It will run the code. For this case <i></i> will do the work.
    (OR) type <i>./email-verifier.exe</i>.
    (OR) double click on the .exe file.
</ol>

## Note
> You can change the name of the .mod file, by entering your desired name after `...init` in <i>Step 1</i>

> Enter into the correct directory before running the .exe file from terminal.