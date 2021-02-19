# passgen
Instead of trusting companies, apps just generate your password every time.

## Why
I have always struggle between keeping all my passwords safe, I use password managers but I am not comfortable with trusting
a prpoitary software so I created passgen.

## How
Passgen uses below algorithm to generate passwords for you based on the app/website.<br>
```
Password = EncryptionAlgorithm([Website/App] Unique Name + Master Password)
```

## Usage
```bash
GO111MODULE=off go get -v github.com/amirrezaask/passgen-cli/cmd/passgen
GO111MODULE=off go install github.com/amirrezaask/passgen-cli/cmd/passgen
passgen generate-config "your master password"
passgen generate primary-gmail-or-any-unique-name
```

